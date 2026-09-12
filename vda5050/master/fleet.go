package master

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"sync"
	"time"

	"github.com/GOAT-Robotics/gtstudio-proto/vda5050"
	"github.com/GOAT-Robotics/gtstudio-proto/vda5050/transport"
)

// Logger is the logging surface the fleet control needs.
type Logger interface {
	Infof(format string, args ...any)
	Warnf(format string, args ...any)
	Errorf(format string, args ...any)
	Debugf(format string, args ...any)
}

type nopLogger struct{}

func (nopLogger) Infof(string, ...any)  {}
func (nopLogger) Warnf(string, ...any)  {}
func (nopLogger) Errorf(string, ...any) {}
func (nopLogger) Debugf(string, ...any) {}

// Options configures a Fleet.
type Options struct {
	// Scheme controls how topics are built. The zero value uses the layout
	// suggested by §4.2 with no broker prefix.
	Scheme vda5050.TopicScheme

	// Discover subscribes with manufacturer/serialNumber wildcards so that any
	// vehicle publishing on the bus is picked up automatically. With it off,
	// only vehicles registered through Register are subscribed to — which is
	// what you want on a broker shared with other systems, or where the fleet
	// roster is authoritative.
	Discover bool

	// ValidateOutgoing checks every published message against the official
	// JSON schema before it goes out. It costs a validation pass per publish
	// and is worth keeping on: a malformed order is otherwise diagnosed only
	// by the vehicle rejecting it, several seconds later and with less detail.
	ValidateOutgoing bool

	// ValidateIncoming checks received messages against their schema. Failures
	// are logged but the message is still processed: a vehicle slightly out of
	// spec is more useful under supervision than silently dropped.
	ValidateIncoming bool

	// StaleAfter is how long without a state message before a vehicle is
	// considered stale. Defaults to three times the specification's 30-second
	// maximum state interval.
	StaleAfter time.Duration

	// Log receives diagnostics.
	Log Logger

	// Hooks. All are optional and run on the transport's delivery goroutines,
	// so they must not block: hand slow work to a queue.
	OnState         func(*Vehicle, *vda5050.State)
	OnConnection    func(*Vehicle, vda5050.ConnectionState)
	OnFactsheet     func(*Vehicle, *vda5050.Factsheet)
	OnVisualization func(*Vehicle, *vda5050.Visualization)

	// OnZoneRequest and OnEdgeRequest decide requests raised by a vehicle in
	// its state (§6.9): zone access or replanning, and corridor use. Returning
	// a zero Decision leaves the request unanswered, which the vehicle treats
	// as "not granted". Wire these into herdIQ's traffic layer so that VDA5050
	// vehicles contend for space on the same terms as native robots.
	OnZoneRequest func(*Vehicle, vda5050.ZoneRequest) Decision
	OnEdgeRequest func(*Vehicle, vda5050.EdgeRequest) Decision

	// ReviewActiveRequests also passes requests the fleet control has already
	// answered -- those sitting in QUEUED or GRANTED -- back to the hooks on
	// every state message, so that a decision can be changed after the fact.
	//
	// Without it there is no path to REVOKED at all: §6.9 has the fleet
	// control withdraw a permission it has granted, and Figure 16 shows
	// exactly that transition, but a hook that only ever sees REQUESTED can
	// never issue one. The cost is that the hooks are called once per active
	// request per state message, so a hook that logs unconditionally will be
	// noisy; returning a zero Decision means "no change" and is free.
	ReviewActiveRequests bool
}

// Decision is the fleet control's answer to a vehicle request.
type Decision struct {
	// Grant is the verdict. The zero value ("") means "do not answer yet".
	Grant vda5050.GrantType
	// LeaseExpiry bounds a GRANTED decision. Zero means no expiry. Prefer a
	// lease: an unbounded grant cannot be reclaimed if the fleet control
	// restarts and forgets it was given.
	LeaseExpiry time.Time
}

// Grant permits a request until expiry. A zero expiry means no time limit.
func Grant(expiry time.Time) Decision {
	return Decision{Grant: vda5050.GrantTypeGranted, LeaseExpiry: expiry}
}

// Queue acknowledges a request without granting it.
func Queue() Decision { return Decision{Grant: vda5050.GrantTypeQueued} }

// Reject refuses a request.
func Reject() Decision { return Decision{Grant: vda5050.GrantTypeRejected} }

// Revoke withdraws a permission granted earlier.
func Revoke() Decision { return Decision{Grant: vda5050.GrantTypeRevoked} }

// Fleet is the VDA5050 fleet control. It owns the broker connection, the
// vehicle registry and the publish/subscribe plumbing.
type Fleet struct {
	broker  transport.Broker
	scheme  vda5050.TopicScheme
	headers *vda5050.HeaderCounter
	opts    Options
	log     Logger

	mu       sync.RWMutex
	vehicles map[string]*Vehicle

	started bool
}

// NewFleet builds a fleet control over the given broker. Passing a disabled
// broker (transport.Disabled, which transport.New returns when no broker URL
// is configured) yields a Fleet whose Start is a no-op and which reports no
// vehicles, so callers need no conditional wiring.
func NewFleet(broker transport.Broker, opts Options) *Fleet {
	if opts.Log == nil {
		opts.Log = nopLogger{}
	}
	if opts.StaleAfter <= 0 {
		opts.StaleAfter = 3 * vda5050.StateIntervalSeconds * time.Second
	}
	scheme := opts.Scheme
	if scheme.InterfaceName == "" && scheme.Version == "" && scheme.Prefix == "" {
		scheme = vda5050.NewTopicScheme("")
	}
	return &Fleet{
		broker:   broker,
		scheme:   scheme,
		headers:  vda5050.NewHeaderCounter(),
		opts:     opts,
		log:      opts.Log,
		vehicles: make(map[string]*Vehicle),
	}
}

// Enabled reports whether VDA5050 is switched on for this deployment.
func (f *Fleet) Enabled() bool {
	_, disabled := f.broker.(transport.Disabled)
	return !disabled
}

// Start connects to the broker and installs subscriptions. It is safe to call
// when VDA5050 is disabled: it returns nil and does nothing, so herdIQ can
// call it unconditionally.
func (f *Fleet) Start(ctx context.Context) error {
	if !f.Enabled() {
		f.log.Infof("[vda5050] no broker configured; VDA5050 fleet control is disabled")
		return nil
	}

	// Subscribe before connecting: the transport records the handlers and
	// installs them in its OnConnect callback, so no retained message
	// published between connect and subscribe can slip past.
	if f.opts.Discover {
		for _, t := range vda5050.RobotTopics {
			topic := t
			filter := f.scheme.SubscribeAll(topic)
			if err := f.broker.Subscribe(ctx, filter, topic.QoS(), f.handler(topic)); err != nil {
				return fmt.Errorf("vda5050: subscribing to %s: %w", filter, err)
			}
		}
	} else {
		f.mu.RLock()
		known := make([]vda5050.Identity, 0, len(f.vehicles))
		for _, v := range f.vehicles {
			known = append(known, v.ID)
		}
		f.mu.RUnlock()
		for _, id := range known {
			if err := f.subscribeVehicle(ctx, id); err != nil {
				return err
			}
		}
	}

	if err := f.broker.Connect(ctx); err != nil {
		return err
	}

	f.mu.Lock()
	f.started = true
	f.mu.Unlock()
	return nil
}

// Stop disconnects from the broker.
func (f *Fleet) Stop() error {
	f.mu.Lock()
	f.started = false
	f.mu.Unlock()
	return f.broker.Close()
}

// Register adds a vehicle to the roster and, once started, subscribes to its
// topics. It is idempotent, so it is safe to call on every config refresh.
func (f *Fleet) Register(ctx context.Context, id vda5050.Identity) (*Vehicle, error) {
	if err := id.Valid(); err != nil {
		return nil, err
	}
	f.mu.Lock()
	v, ok := f.vehicles[id.String()]
	if !ok {
		v = newVehicle(id, f.headers)
		f.vehicles[id.String()] = v
	}
	started := f.started
	f.mu.Unlock()

	if !ok {
		f.log.Infof("[vda5050] registered vehicle %s", id)
	}
	// With discovery on, the wildcard subscription already covers this
	// vehicle; adding a specific one would only duplicate deliveries.
	if started && f.Enabled() && !f.opts.Discover && !ok {
		if err := f.subscribeVehicle(ctx, id); err != nil {
			return v, err
		}
	}
	return v, nil
}

// Deregister removes a vehicle and its subscriptions.
func (f *Fleet) Deregister(ctx context.Context, id vda5050.Identity) {
	f.mu.Lock()
	delete(f.vehicles, id.String())
	f.mu.Unlock()
	f.headers.Reset(id)

	if f.Enabled() && !f.opts.Discover {
		for _, t := range vda5050.RobotTopics {
			_ = f.broker.Unsubscribe(ctx, f.scheme.Build(id, t))
		}
	}
	f.log.Infof("[vda5050] deregistered vehicle %s", id)
}

func (f *Fleet) subscribeVehicle(ctx context.Context, id vda5050.Identity) error {
	for _, t := range vda5050.RobotTopics {
		topic := t
		filter := f.scheme.Build(id, topic)
		if err := f.broker.Subscribe(ctx, filter, topic.QoS(), f.handler(topic)); err != nil {
			return fmt.Errorf("vda5050: subscribing to %s: %w", filter, err)
		}
	}
	return nil
}

// Vehicle returns a registered vehicle, or nil.
func (f *Fleet) Vehicle(id vda5050.Identity) *Vehicle {
	f.mu.RLock()
	defer f.mu.RUnlock()
	return f.vehicles[id.String()]
}

// Vehicles returns every known vehicle, ordered by identity so that callers
// rendering a fleet view get a stable list rather than Go's randomised map
// iteration order.
func (f *Fleet) Vehicles() []*Vehicle {
	f.mu.RLock()
	out := make([]*Vehicle, 0, len(f.vehicles))
	for _, v := range f.vehicles {
		out = append(out, v)
	}
	f.mu.RUnlock()
	sort.Slice(out, func(i, j int) bool { return out[i].ID.String() < out[j].ID.String() })
	return out
}

// ---------------------------------------------------------------------------
// Publishing
// ---------------------------------------------------------------------------

// ErrNotRegistered is returned when publishing to an unknown vehicle.
var ErrNotRegistered = errors.New("vda5050: vehicle is not registered with this fleet control")

// SendOrder publishes an order. Build it with the vehicle's OrderTracker so
// that orderId, orderUpdateId, sequence IDs and the base/horizon split are
// consistent with what has already been sent.
func (f *Fleet) SendOrder(ctx context.Context, id vda5050.Identity, o *vda5050.Order) error {
	return f.publish(ctx, id, vda5050.TopicOrder, o)
}

// SendInstantActions publishes one or more instant actions.
func (f *Fleet) SendInstantActions(ctx context.Context, id vda5050.Identity, actions ...vda5050.InstantAction) error {
	if len(actions) == 0 {
		return nil
	}
	msg := &vda5050.InstantActions{Actions: actions}
	f.headers.NewHeader(id, vda5050.TopicInstantActions).ApplyInstantActions(msg)
	return f.publish(ctx, id, vda5050.TopicInstantActions, msg)
}

// SendZoneSet publishes a zone set (§6.4, new in 3.0). Zones let fleet
// control shape vehicle behaviour by area — blocked regions, speed limits,
// one-way corridors — without encoding it into every edge of every order.
func (f *Fleet) SendZoneSet(ctx context.Context, id vda5050.Identity, zs vda5050.ZoneSet) error {
	msg := &vda5050.ZoneSetMessage{ZoneSet: zs}
	f.headers.NewHeader(id, vda5050.TopicZoneSet).ApplyZoneSet(msg)
	return f.publish(ctx, id, vda5050.TopicZoneSet, msg)
}

// SendResponses answers requests a vehicle raised in its state (§6.9).
func (f *Fleet) SendResponses(ctx context.Context, id vda5050.Identity, responses ...vda5050.Response) error {
	if len(responses) == 0 {
		return nil
	}
	msg := &vda5050.Responses{Responses: responses}
	f.headers.NewHeader(id, vda5050.TopicResponses).ApplyResponses(msg)
	return f.publish(ctx, id, vda5050.TopicResponses, msg)
}

// CancelOrder cancels a vehicle's active order and marks its tracker so that
// no further updates are sent for that orderId (§6.1.3.1). The order is not
// finished until the vehicle reports the cancelOrder action FINISHED.
func (f *Fleet) CancelOrder(ctx context.Context, id vda5050.Identity) error {
	v := f.Vehicle(id)
	if v == nil {
		return ErrNotRegistered
	}
	action, err := v.Orders.Cancel()
	if err != nil {
		return err
	}
	return f.SendInstantActions(ctx, id, action)
}

// Pause and Resume drive the mandatory startPause/stopPause instant actions.
func (f *Fleet) Pause(ctx context.Context, id vda5050.Identity) error {
	return f.SendInstantActions(ctx, id, vda5050.StartPause())
}

// Resume lifts a pause applied with Pause.
func (f *Fleet) Resume(ctx context.Context, id vda5050.Identity) error {
	return f.SendInstantActions(ctx, id, vda5050.StopPause())
}

// RequestFactsheet asks a vehicle to republish its factsheet.
func (f *Fleet) RequestFactsheet(ctx context.Context, id vda5050.Identity) error {
	return f.SendInstantActions(ctx, id, vda5050.FactsheetRequest())
}

// RequestState asks a vehicle to publish a state message immediately rather
// than waiting for its next scheduled one.
func (f *Fleet) RequestState(ctx context.Context, id vda5050.Identity) error {
	return f.SendInstantActions(ctx, id, vda5050.StateRequest())
}

func (f *Fleet) publish(ctx context.Context, id vda5050.Identity, topic vda5050.Topic, msg any) error {
	if !f.Enabled() {
		return transport.ErrDisabled
	}
	if f.Vehicle(id) == nil {
		return fmt.Errorf("%w: %s", ErrNotRegistered, id)
	}
	if f.opts.ValidateOutgoing {
		if err := vda5050.Validate(topic, msg); err != nil {
			return err
		}
	}
	payload, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("vda5050: marshalling %s for %s: %w", topic, id, err)
	}
	t := f.scheme.Build(id, topic)
	f.log.Debugf("[vda5050] -> %s (%d bytes)", t, len(payload))
	return f.broker.Publish(ctx, t, topic.QoS(), topic.Retained(), payload)
}

// ---------------------------------------------------------------------------
// Receiving
// ---------------------------------------------------------------------------

func (f *Fleet) handler(topic vda5050.Topic) transport.Handler {
	return func(m transport.Message) {
		id, parsed, err := f.scheme.Parse(m.Topic)
		if err != nil {
			// Another system's traffic on a shared broker; ignore quietly.
			f.log.Debugf("[vda5050] ignoring message on %s: %v", m.Topic, err)
			return
		}
		if parsed != topic {
			return
		}
		if f.opts.ValidateIncoming {
			if err := vda5050.ValidateRaw(topic, m.Payload); err != nil {
				f.log.Warnf("[vda5050] %s from %s does not match the schema, processing anyway: %v", topic, id, err)
			}
		}
		if err := f.dispatch(id, topic, m); err != nil {
			f.log.Errorf("[vda5050] handling %s from %s: %v", topic, id, err)
		}
	}
}

func (f *Fleet) dispatch(id vda5050.Identity, topic vda5050.Topic, m transport.Message) error {
	v := f.Vehicle(id)
	if v == nil {
		if !f.opts.Discover {
			// Not ours to manage.
			return nil
		}
		var err error
		v, err = f.Register(context.Background(), id)
		if err != nil {
			return err
		}
		f.log.Infof("[vda5050] discovered vehicle %s", id)
	}

	switch topic {
	case vda5050.TopicState:
		var s vda5050.State
		if err := json.Unmarshal(m.Payload, &s); err != nil {
			return fmt.Errorf("decoding state: %w", err)
		}
		v.mu.Lock()
		v.state, v.stateAt = &s, time.Now()
		v.mu.Unlock()
		f.answerRequests(v, &s)
		if f.opts.OnState != nil {
			f.opts.OnState(v, &s)
		}

	case vda5050.TopicConnection:
		var c vda5050.Connection
		if err := json.Unmarshal(m.Payload, &c); err != nil {
			return fmt.Errorf("decoding connection: %w", err)
		}
		v.mu.Lock()
		prev := v.conn
		v.conn, v.connAt = c.ConnectionState, time.Now()
		v.mu.Unlock()
		if prev != c.ConnectionState {
			f.log.Infof("[vda5050] %s connection %s -> %s", id, prev, c.ConnectionState)
		}
		if f.opts.OnConnection != nil {
			f.opts.OnConnection(v, c.ConnectionState)
		}

	case vda5050.TopicFactsheet:
		var fs vda5050.Factsheet
		if err := json.Unmarshal(m.Payload, &fs); err != nil {
			return fmt.Errorf("decoding factsheet: %w", err)
		}
		v.mu.Lock()
		v.factsheet, v.factsheetAt = &fs, time.Now()
		v.mu.Unlock()
		f.log.Infof("[vda5050] %s factsheet: series=%q class=%s kinematics=%s protocol=%s",
			id, fs.TypeSpecification.SeriesName, fs.TypeSpecification.MobileRobotClass,
			fs.TypeSpecification.MobileRobotKinematics, fs.Version)
		if f.opts.OnFactsheet != nil {
			f.opts.OnFactsheet(v, &fs)
		}

	case vda5050.TopicVisualization:
		var vis vda5050.Visualization
		if err := json.Unmarshal(m.Payload, &vis); err != nil {
			return fmt.Errorf("decoding visualization: %w", err)
		}
		v.mu.Lock()
		v.vis, v.visAt = &vis, time.Now()
		v.mu.Unlock()
		if f.opts.OnVisualization != nil {
			f.opts.OnVisualization(v, &vis)
		}
	}
	return nil
}

// decidable reports whether a request in this state still wants a decision
// from the fleet control.
//
// REQUESTED always does. QUEUED and GRANTED do only when the fleet control
// asked to review them, which is what makes revocation and lease extension
// possible. REJECTED, REVOKED and EXPIRED are terminal: the vehicle is about
// to drop the request from its state, and answering again would churn it.
func (f *Fleet) decidable(status vda5050.RequestStatus) bool {
	switch status {
	case vda5050.RequestStatusRequested:
		return true
	case vda5050.RequestStatusQueued, vda5050.RequestStatusGranted:
		return f.opts.ReviewActiveRequests
	}
	return false
}

// answerRequests replies to zone and corridor requests carried in a state
// message (§6.9).
func (f *Fleet) answerRequests(v *Vehicle, s *vda5050.State) {
	if f.opts.OnZoneRequest == nil && f.opts.OnEdgeRequest == nil {
		return
	}
	var responses []vda5050.Response

	if f.opts.OnZoneRequest != nil {
		for _, r := range s.ZoneRequests {
			if !f.decidable(r.RequestStatus) {
				continue
			}
			if resp, ok := toResponse(r.RequestID, f.opts.OnZoneRequest(v, r)); ok {
				responses = append(responses, resp)
			}
		}
	}
	if f.opts.OnEdgeRequest != nil {
		for _, r := range s.EdgeRequests {
			if !f.decidable(r.RequestStatus) {
				continue
			}
			if resp, ok := toResponse(r.RequestID, f.opts.OnEdgeRequest(v, r)); ok {
				responses = append(responses, resp)
			}
		}
	}
	if len(responses) == 0 {
		return
	}
	if err := f.SendResponses(context.Background(), v.ID, responses...); err != nil {
		f.log.Errorf("[vda5050] answering requests from %s: %v", v.ID, err)
	}
}

// RevokeRequests withdraws permissions the fleet control granted earlier
// (§6.9). The vehicle reacts according to the releaseLossBehavior defined for
// the resource: stopping, continuing, or evacuating a zone; returning to the
// predefined trajectory of an edge.
//
// A revoked grant is not instantaneous. The fleet control "shall assume a
// REVOKED request as still being GRANTED until the requestStatus of the
// mobile robot is set to REVOKED", so the space stays committed until the
// vehicle's own state confirms it has let go.
func (f *Fleet) RevokeRequests(ctx context.Context, id vda5050.Identity, requestIDs ...string) error {
	if len(requestIDs) == 0 {
		return nil
	}
	responses := make([]vda5050.Response, 0, len(requestIDs))
	for _, rid := range requestIDs {
		responses = append(responses, vda5050.Response{
			RequestID: rid,
			GrantType: vda5050.GrantTypeRevoked,
		})
	}
	return f.SendResponses(ctx, id, responses...)
}

// ExtendLease re-grants a request with a later expiry (§6.9). Sending an
// updated response with the same requestId and a new leaseExpiry is the only
// way to keep a vehicle inside a RELEASE zone past its original lease.
func (f *Fleet) ExtendLease(ctx context.Context, id vda5050.Identity, requestID string, until time.Time) error {
	resp, ok := toResponse(requestID, Grant(until))
	if !ok {
		return errors.New("vda5050: lease extension needs a grant decision")
	}
	return f.SendResponses(ctx, id, resp)
}

func toResponse(requestID string, d Decision) (vda5050.Response, bool) {
	if d.Grant == "" {
		return vda5050.Response{}, false
	}
	r := vda5050.Response{RequestID: requestID, GrantType: d.Grant}
	// A lease is only meaningful on a grant; §7.5 restricts leaseExpiry to
	// responses that grant a request.
	if d.Grant == vda5050.GrantTypeGranted && !d.LeaseExpiry.IsZero() {
		r.LeaseExpiry = vda5050.Str(vda5050.Timestamp(d.LeaseExpiry))
	}
	return r, true
}
