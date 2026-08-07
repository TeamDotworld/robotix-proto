package vda5050

import (
	"sync"
	"sync/atomic"
	"time"
)

// timestampLayout is the format required by §7.2: ISO 8601 in UTC with
// millisecond precision. Version 3.0.0 moved from 1/100 s to 1/1000 s, so a
// 2.x implementation emitting two fractional digits is no longer compliant.
const timestampLayout = "2006-01-02T15:04:05.000Z"

// Timestamp formats t as a VDA5050 header timestamp. The value is converted
// to UTC first; sending local time is a common and hard-to-spot interop bug.
func Timestamp(t time.Time) string {
	return t.UTC().Format(timestampLayout)
}

// Now returns the current time as a VDA5050 header timestamp.
func Now() string { return Timestamp(time.Now()) }

// ParseTimestamp parses a VDA5050 header timestamp. It accepts any RFC 3339
// value so that vehicles emitting more or fewer fractional digits than the
// specification requires are still understood; being strict here would mean
// discarding otherwise usable state messages.
func ParseTimestamp(s string) (time.Time, error) {
	if t, err := time.Parse(timestampLayout, s); err == nil {
		return t, nil
	}
	return time.Parse(time.RFC3339Nano, s)
}

// Header is the common preamble of every VDA5050 message (§7.2). It is not a
// nested JSON object: the fields sit at the top level of each message, which
// is why the generated message types repeat them rather than embedding this.
type Header struct {
	HeaderID     uint32 `json:"headerId"`
	Timestamp    string `json:"timestamp"`
	Version      string `json:"version"`
	Manufacturer string `json:"manufacturer"`
	SerialNumber string `json:"serialNumber"`
}

// HeaderCounter issues headerIds. The specification requires a separate,
// monotonically increasing counter per vehicle *and* per topic, so that a
// receiver can detect messages lost or reordered on one topic without being
// confused by traffic on another.
//
// A HeaderCounter is safe for concurrent use.
type HeaderCounter struct {
	mu  sync.Mutex
	seq map[string]*uint32
}

// NewHeaderCounter returns an empty counter set.
func NewHeaderCounter() *HeaderCounter {
	return &HeaderCounter{seq: make(map[string]*uint32)}
}

// Next returns the next headerId for a vehicle/topic pair, starting at 0.
func (h *HeaderCounter) Next(id Identity, topic Topic) uint32 {
	key := id.String() + "/" + string(topic)
	h.mu.Lock()
	p, ok := h.seq[key]
	if !ok {
		p = new(uint32)
		h.seq[key] = p
	}
	h.mu.Unlock()
	// Subtract one so the first issued value is 0 rather than 1.
	return atomic.AddUint32(p, 1) - 1
}

// Reset drops the counters for a vehicle. Call this when a vehicle is removed
// from the fleet, not on reconnect: headerIds should keep climbing across
// reconnects so that a peer can tell a stale retained message from a new one.
func (h *HeaderCounter) Reset(id Identity) {
	prefix := id.String() + "/"
	h.mu.Lock()
	defer h.mu.Unlock()
	for k := range h.seq {
		if len(k) > len(prefix) && k[:len(prefix)] == prefix {
			delete(h.seq, k)
		}
	}
}

// NewHeader builds a header for an outgoing message.
func (h *HeaderCounter) NewHeader(id Identity, topic Topic) Header {
	return Header{
		HeaderID:     h.Next(id, topic),
		Timestamp:    Now(),
		Version:      ProtocolVersion,
		Manufacturer: id.Manufacturer,
		SerialNumber: id.SerialNumber,
	}
}

// Apply stamps a header onto an outgoing message. Because the generated types
// carry the header fields inline rather than as an embedded struct, each
// message type gets a small setter here.
func (h Header) Apply(o *Order) {
	o.HeaderID, o.Timestamp, o.Version = h.HeaderID, h.Timestamp, h.Version
	o.Manufacturer, o.SerialNumber = h.Manufacturer, h.SerialNumber
}

// ApplyInstantActions stamps a header onto an instantActions message.
func (h Header) ApplyInstantActions(m *InstantActions) {
	m.HeaderID, m.Timestamp, m.Version = h.HeaderID, h.Timestamp, h.Version
	m.Manufacturer, m.SerialNumber = h.Manufacturer, h.SerialNumber
}

// ApplyZoneSet stamps a header onto a zoneSet message.
func (h Header) ApplyZoneSet(m *ZoneSetMessage) {
	m.HeaderID, m.Timestamp, m.Version = h.HeaderID, h.Timestamp, h.Version
	m.Manufacturer, m.SerialNumber = h.Manufacturer, h.SerialNumber
}

// ApplyResponses stamps a header onto a responses message.
func (h Header) ApplyResponses(m *Responses) {
	m.HeaderID, m.Timestamp, m.Version = h.HeaderID, h.Timestamp, h.Version
	m.Manufacturer, m.SerialNumber = h.Manufacturer, h.SerialNumber
}
