package vda5050

import (
	"fmt"
	"strings"
)

// Identity names one vehicle on the bus. VDA5050 addresses vehicles by the
// manufacturer/serialNumber pair rather than by a fleet-assigned ID, and both
// values are repeated inside every message header.
type Identity struct {
	Manufacturer string
	SerialNumber string
}

// String renders the identity as "manufacturer/serialNumber".
func (i Identity) String() string { return i.Manufacturer + "/" + i.SerialNumber }

// Valid reports whether both parts are present and free of characters that
// would break the MQTT topic hierarchy (§4.2).
func (i Identity) Valid() error {
	if err := validTopicLevel("manufacturer", i.Manufacturer); err != nil {
		return err
	}
	return validTopicLevel("serialNumber", i.SerialNumber)
}

// forbiddenInTopicLevel are characters that must not appear in a topic level:
// "/" delimits the hierarchy, "+" and "#" are wildcards, and "$" is reserved
// for broker-internal topics.
const forbiddenInTopicLevel = "/+#$"

func validTopicLevel(field, v string) error {
	if v == "" {
		return fmt.Errorf("vda5050: %s must not be empty", field)
	}
	if i := strings.IndexAny(v, forbiddenInTopicLevel); i >= 0 {
		return fmt.Errorf("vda5050: %s %q contains reserved character %q", field, v, v[i])
	}
	return nil
}

// TopicScheme builds and parses MQTT topics. The specification fixes only the
// final level; the preceding levels are configurable because cloud brokers
// (AWS IoT, Azure IoT Hub) impose their own mandatory prefixes.
//
// The zero value is not usable; call NewTopicScheme.
type TopicScheme struct {
	// Prefix is prepended verbatim, without a trailing slash. Empty for a
	// local broker following the suggested layout.
	Prefix string
	// InterfaceName is the first level after Prefix (default "vda5050").
	InterfaceName string
	// Version is the second level (default "v3").
	Version string
}

// NewTopicScheme returns the layout suggested by §4.2:
//
//	vda5050/v3/{manufacturer}/{serialNumber}/{topic}
//
// Pass a non-empty prefix to nest that beneath a broker-mandated root.
func NewTopicScheme(prefix string) TopicScheme {
	return TopicScheme{
		Prefix:        strings.Trim(prefix, "/"),
		InterfaceName: DefaultInterfaceName,
		Version:       MajorVersion,
	}
}

// levels returns the fixed leading levels of every topic in this scheme.
//
// A prefix may itself span several levels ("tenants/acme/plant-a"), which
// cloud brokers frequently require, so it is split rather than treated as one
// level. Parse counts levels to find the manufacturer and serial number, and
// would otherwise miscount by exactly the number of slashes in the prefix.
func (s TopicScheme) levels() []string {
	iface := s.InterfaceName
	if iface == "" {
		iface = DefaultInterfaceName
	}
	ver := s.Version
	if ver == "" {
		ver = MajorVersion
	}
	if s.Prefix == "" {
		return []string{iface, ver}
	}
	prefix := strings.Split(strings.Trim(s.Prefix, "/"), "/")
	out := make([]string, 0, len(prefix)+2)
	out = append(out, prefix...)
	return append(out, iface, ver)
}

// Build returns the full topic for one vehicle and one topic name.
func (s TopicScheme) Build(id Identity, topic Topic) string {
	parts := append(s.levels(), id.Manufacturer, id.SerialNumber, string(topic))
	return strings.Join(parts, "/")
}

// SubscribeAll returns a wildcard filter matching the given topic for every
// manufacturer and serial number, e.g. "vda5050/v3/+/+/state". A fleet
// control uses this to discover vehicles it has not been told about.
func (s TopicScheme) SubscribeAll(topic Topic) string {
	parts := append(s.levels(), "+", "+", string(topic))
	return strings.Join(parts, "/")
}

// SubscribeManufacturer returns a filter matching one topic across every
// serial number of a single manufacturer.
func (s TopicScheme) SubscribeManufacturer(manufacturer string, topic Topic) string {
	parts := append(s.levels(), manufacturer, "+", string(topic))
	return strings.Join(parts, "/")
}

// Parse extracts the vehicle identity and topic from a received topic string.
// It returns an error if the topic does not belong to this scheme, which lets
// a handler safely ignore unrelated traffic on a shared broker.
func (s TopicScheme) Parse(topic string) (Identity, Topic, error) {
	parts := strings.Split(strings.Trim(topic, "/"), "/")
	lead := s.levels()
	// lead + manufacturer + serialNumber + topic
	if len(parts) != len(lead)+3 {
		return Identity{}, "", fmt.Errorf("vda5050: topic %q has %d levels, want %d", topic, len(parts), len(lead)+3)
	}
	for i, want := range lead {
		if parts[i] != want {
			return Identity{}, "", fmt.Errorf("vda5050: topic %q level %d is %q, want %q", topic, i, parts[i], want)
		}
	}
	rest := parts[len(lead):]
	id := Identity{Manufacturer: rest[0], SerialNumber: rest[1]}
	t := Topic(rest[2])
	if !t.known() {
		return Identity{}, "", fmt.Errorf("vda5050: topic %q ends in unknown topic %q", topic, rest[2])
	}
	return id, t, nil
}

func (t Topic) known() bool {
	for _, k := range AllTopics {
		if t == k {
			return true
		}
	}
	return false
}
