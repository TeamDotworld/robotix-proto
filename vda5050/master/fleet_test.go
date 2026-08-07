package master

import (
	"context"
	"encoding/json"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/GOAT-Robotics/gtstudio-proto/vda5050"
	"github.com/GOAT-Robotics/gtstudio-proto/vda5050/transport"
)

// fakeBroker is an in-memory MQTT stand-in. It records what the fleet control
// published and lets a test feed messages back as if a vehicle had sent them,
// so the whole master-control loop can be exercised without a broker.
type fakeBroker struct {
	mu        sync.Mutex
	published []publication
	subs      map[string]transport.Handler
	connected bool
}

type publication struct {
	Topic   string
	QoS     byte
	Retain  bool
	Payload []byte
}

func newFakeBroker() *fakeBroker {
	return &fakeBroker{subs: map[string]transport.Handler{}}
}

func (b *fakeBroker) Connect(context.Context) error {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.connected = true
	return nil
}

func (b *fakeBroker) Publish(_ context.Context, topic string, qos byte, retain bool, payload []byte) error {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.published = append(b.published, publication{topic, qos, retain, append([]byte(nil), payload...)})
	return nil
}

func (b *fakeBroker) Subscribe(_ context.Context, filter string, _ byte, h transport.Handler) error {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.subs[filter] = h
	return nil
}

func (b *fakeBroker) Unsubscribe(_ context.Context, filter string) error {
	b.mu.Lock()
	defer b.mu.Unlock()
	delete(b.subs, filter)
	return nil
}

func (b *fakeBroker) Connected() bool {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.connected
}

func (b *fakeBroker) Close() error { return nil }

// deliver routes a message to every handler whose filter matches, applying
// MQTT '+' single-level wildcard semantics.
func (b *fakeBroker) deliver(topic string, payload []byte) {
	b.mu.Lock()
	handlers := make([]transport.Handler, 0, len(b.subs))
	for filter, h := range b.subs {
		if topicMatches(filter, topic) {
			handlers = append(handlers, h)
		}
	}
	b.mu.Unlock()
	for _, h := range handlers {
		h(transport.Message{Topic: topic, Payload: payload})
	}
}

func topicMatches(filter, topic string) bool {
	f, t := strings.Split(filter, "/"), strings.Split(topic, "/")
	if len(f) != len(t) {
		return false
	}
	for i := range f {
		if f[i] != "+" && f[i] != t[i] {
			return false
		}
	}
	return true
}

func (b *fakeBroker) sent() []publication {
	b.mu.Lock()
	defer b.mu.Unlock()
	return append([]publication(nil), b.published...)
}

// ---------------------------------------------------------------------------

var testID = vda5050.Identity{Manufacturer: "KIT", SerialNumber: "0001"}

func stateJSON(t *testing.T, s vda5050.State) []byte {
	t.Helper()
	s.Version = vda5050.ProtocolVersion
	s.Manufacturer, s.SerialNumber = testID.Manufacturer, testID.SerialNumber
	s.Timestamp = vda5050.Now()
	if s.NodeStates == nil {
		s.NodeStates = []vda5050.NodeState{}
	}
	if s.EdgeStates == nil {
		s.EdgeStates = []vda5050.EdgeState{}
	}
	if s.ActionStates == nil {
		s.ActionStates = []vda5050.ActionState{}
	}
	if s.InstantActionStates == nil {
		s.InstantActionStates = []vda5050.ActionState{}
	}
	if s.Errors == nil {
		s.Errors = []vda5050.Error{}
	}
	raw, err := json.Marshal(s)
	if err != nil {
		t.Fatalf("marshal state: %v", err)
	}
	return raw
}

func TestStateIsIngestedAndSurfaced(t *testing.T) {
	b := newFakeBroker()
	f := NewFleet(b, Options{})
	ctx := context.Background()

	if _, err := f.Register(ctx, testID); err != nil {
		t.Fatalf("Register: %v", err)
	}
	if err := f.Start(ctx); err != nil {
		t.Fatalf("Start: %v", err)
	}

	b.deliver("vda5050/v3/KIT/0001/state", stateJSON(t, vda5050.State{
		OperatingMode: vda5050.OperatingModeAutomatic,
		Driving:       true,
		PowerSupply:   vda5050.PowerSupply{StateOfCharge: 82.5},
		MobileRobotPosition: &vda5050.MobileRobotPosition{
			X: 3.5, Y: -1.25, Theta: 1.57, MapID: "level_1", Localized: true,
		},
	}))

	v := f.Vehicle(testID)
	if v == nil {
		t.Fatal("vehicle should be registered")
	}
	if v.State() == nil {
		t.Fatal("state should have been ingested")
	}
	if soc, ok := v.BatteryCharge(); !ok || soc != 82.5 {
		t.Errorf("battery = %v (ok=%v), want 82.5", soc, ok)
	}
	x, y, _, mapID, ok := v.Position()
	if !ok || x != 3.5 || y != -1.25 || mapID != "level_1" {
		t.Errorf("position = (%v,%v,%q) ok=%v", x, y, mapID, ok)
	}
	if !v.Driving() {
		t.Error("vehicle should report driving")
	}
}

func TestUnlocalisedPositionIsNotReported(t *testing.T) {
	// A vehicle that is not localised is publishing a pose it does not trust.
	// Passing it on would let the traffic layer reserve lanes around a
	// position the vehicle is not at.
	b := newFakeBroker()
	f := NewFleet(b, Options{})
	ctx := context.Background()
	_, _ = f.Register(ctx, testID)
	_ = f.Start(ctx)

	b.deliver("vda5050/v3/KIT/0001/state", stateJSON(t, vda5050.State{
		OperatingMode:       vda5050.OperatingModeAutomatic,
		MobileRobotPosition: &vda5050.MobileRobotPosition{X: 9, Y: 9, MapID: "level_1", Localized: false},
	}))

	if _, _, _, _, ok := f.Vehicle(testID).Position(); ok {
		t.Error("an unlocalised position must not be reported as usable")
	}
}

func TestDispatchableGuards(t *testing.T) {
	b := newFakeBroker()
	f := NewFleet(b, Options{})
	ctx := context.Background()
	_, _ = f.Register(ctx, testID)
	_ = f.Start(ctx)
	v := f.Vehicle(testID)

	// safe is what a healthy, conformant vehicle reports.
	safe := vda5050.SafetyState{ActiveEmergencyStop: vda5050.ActiveEmergencyStopNone}

	cases := []struct {
		name  string
		state vda5050.State
		want  bool
	}{
		{"automatic and healthy", vda5050.State{
			OperatingMode: vda5050.OperatingModeAutomatic, SafetyState: safe,
		}, true},
		{"manual mode", vda5050.State{
			OperatingMode: vda5050.OperatingModeManual, SafetyState: safe,
		}, false},
		{"service mode", vda5050.State{
			OperatingMode: vda5050.OperatingModeService, SafetyState: safe,
		}, false},
		{"paused", vda5050.State{
			OperatingMode: vda5050.OperatingModeAutomatic, SafetyState: safe, Paused: vda5050.Bool(true),
		}, false},
		{"emergency stop", vda5050.State{
			OperatingMode: vda5050.OperatingModeAutomatic,
			SafetyState:   vda5050.SafetyState{ActiveEmergencyStop: vda5050.ActiveEmergencyStopManual},
		}, false},
		{"remote emergency stop", vda5050.State{
			OperatingMode: vda5050.OperatingModeAutomatic,
			SafetyState:   vda5050.SafetyState{ActiveEmergencyStop: vda5050.ActiveEmergencyStopRemote},
		}, false},
		{"protective field violated", vda5050.State{
			OperatingMode: vda5050.OperatingModeAutomatic,
			SafetyState:   vda5050.SafetyState{ActiveEmergencyStop: vda5050.ActiveEmergencyStopNone, FieldViolation: true},
		}, false},
		{"fatal error", vda5050.State{
			OperatingMode: vda5050.OperatingModeAutomatic, SafetyState: safe,
			Errors: []vda5050.Error{{ErrorType: "motorFault", ErrorLevel: vda5050.ErrorLevelFatal}},
		}, false},
		{"critical error", vda5050.State{
			OperatingMode: vda5050.OperatingModeAutomatic, SafetyState: safe,
			Errors: []vda5050.Error{{ErrorType: "lidarBlind", ErrorLevel: vda5050.ErrorLevelCritical}},
		}, false},
		// WARNING and URGENT are informational: 3.0.0 introduced URGENT
		// precisely so that "needs attention soon" is distinguishable from
		// "stop using this vehicle".
		{"warning error", vda5050.State{
			OperatingMode: vda5050.OperatingModeAutomatic, SafetyState: safe,
			Errors: []vda5050.Error{{ErrorType: "lowInk", ErrorLevel: vda5050.ErrorLevelWarning}},
		}, true},
		{"urgent error", vda5050.State{
			OperatingMode: vda5050.OperatingModeAutomatic, SafetyState: safe,
			Errors: []vda5050.Error{{ErrorType: "batteryAging", ErrorLevel: vda5050.ErrorLevelUrgent}},
		}, true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			b.deliver("vda5050/v3/KIT/0001/connection", connectionJSON(t, vda5050.ConnectionStateOnline))
			b.deliver("vda5050/v3/KIT/0001/state", stateJSON(t, tc.state))
			got, reason := v.Dispatchable(time.Minute)
			if got != tc.want {
				t.Errorf("Dispatchable = %v (%s), want %v", got, reason, tc.want)
			}
		})
	}
}

func connectionJSON(t *testing.T, s vda5050.ConnectionState) []byte {
	t.Helper()
	raw, err := json.Marshal(vda5050.Connection{
		HeaderID: 1, Timestamp: vda5050.Now(), Version: vda5050.ProtocolVersion,
		Manufacturer: testID.Manufacturer, SerialNumber: testID.SerialNumber,
		ConnectionState: s,
	})
	if err != nil {
		t.Fatalf("marshal connection: %v", err)
	}
	return raw
}

func TestSendOrderPublishesToTheRightTopicAtTheRightQoS(t *testing.T) {
	b := newFakeBroker()
	f := NewFleet(b, Options{ValidateOutgoing: true})
	ctx := context.Background()
	v, _ := f.Register(ctx, testID)
	_ = f.Start(ctx)

	r := vda5050.Route{Nodes: []vda5050.RouteNode{
		{NodeID: "a", Position: &vda5050.NodePosition{X: 0, Y: 0, MapID: "level_1"}},
		{NodeID: "b", Position: &vda5050.NodePosition{X: 5, Y: 0, MapID: "level_1"}},
	}, Edges: []vda5050.RouteEdge{{EdgeID: "e1"}}}

	order, err := v.Orders.Start(r, 1)
	if err != nil {
		t.Fatalf("Start: %v", err)
	}
	if err := f.SendOrder(ctx, testID, order); err != nil {
		t.Fatalf("SendOrder: %v", err)
	}

	sent := b.sent()
	if len(sent) != 1 {
		t.Fatalf("want exactly 1 publication, got %d", len(sent))
	}
	p := sent[0]
	if p.Topic != "vda5050/v3/KIT/0001/order" {
		t.Errorf("topic = %q", p.Topic)
	}
	if p.QoS != 0 {
		t.Errorf("order must go out at QoS 0, got %d", p.QoS)
	}
	if p.Retain {
		t.Error("orders must not be retained: a reconnecting vehicle would replay a stale order")
	}
	if err := vda5050.ValidateRaw(vda5050.TopicOrder, p.Payload); err != nil {
		t.Errorf("published order does not conform: %v", err)
	}
}

func TestPublishToAnUnknownVehicleIsRejected(t *testing.T) {
	b := newFakeBroker()
	f := NewFleet(b, Options{})
	ctx := context.Background()
	_ = f.Start(ctx)

	err := f.SendInstantActions(ctx, testID, vda5050.StartPause())
	if err == nil {
		t.Fatal("want an error when publishing to an unregistered vehicle")
	}
}

func TestRequestsAreAnsweredOnTheResponsesTopic(t *testing.T) {
	b := newFakeBroker()
	expiry := time.Now().Add(30 * time.Second)
	f := NewFleet(b, Options{
		OnEdgeRequest: func(*Vehicle, vda5050.EdgeRequest) Decision { return Grant(expiry) },
		OnZoneRequest: func(*Vehicle, vda5050.ZoneRequest) Decision { return Reject() },
	})
	ctx := context.Background()
	_, _ = f.Register(ctx, testID)
	_ = f.Start(ctx)

	b.deliver("vda5050/v3/KIT/0001/state", stateJSON(t, vda5050.State{
		OperatingMode: vda5050.OperatingModeAutomatic,
		EdgeRequests: []vda5050.EdgeRequest{{
			RequestID: "edge-1", RequestType: vda5050.EdgeRequestRequestTypeCorridor,
			EdgeID: "e1", SequenceID: 1, RequestStatus: vda5050.RequestStatusRequested,
		}},
		ZoneRequests: []vda5050.ZoneRequest{
			{
				RequestID: "zone-1", RequestType: vda5050.RequestTypeAccess,
				ZoneID: "z1", ZoneSetID: "zs1", RequestStatus: vda5050.RequestStatusRequested,
			},
			{
				// Already decided. Answering it again would churn the
				// vehicle's state for no reason, so it must be skipped.
				RequestID: "zone-2", RequestType: vda5050.RequestTypeAccess,
				ZoneID: "z2", ZoneSetID: "zs1", RequestStatus: vda5050.RequestStatusGranted,
			},
		},
	}))

	sent := b.sent()
	if len(sent) != 1 {
		t.Fatalf("want one responses publication, got %d", len(sent))
	}
	if sent[0].Topic != "vda5050/v3/KIT/0001/responses" {
		t.Errorf("topic = %q", sent[0].Topic)
	}

	var msg vda5050.Responses
	if err := json.Unmarshal(sent[0].Payload, &msg); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if len(msg.Responses) != 2 {
		t.Fatalf("want 2 responses, got %d", len(msg.Responses))
	}
	byID := map[string]vda5050.Response{}
	for _, r := range msg.Responses {
		byID[r.RequestID] = r
	}
	if got := byID["edge-1"]; got.GrantType != vda5050.GrantTypeGranted || got.LeaseExpiry == nil {
		t.Errorf("edge request should be granted with a lease, got %+v", got)
	}
	if got := byID["zone-1"]; got.GrantType != vda5050.GrantTypeRejected {
		t.Errorf("zone request should be rejected, got %+v", got)
	}
	if got := byID["zone-1"]; got.LeaseExpiry != nil {
		t.Error("a lease is only meaningful on a grant")
	}
	if _, answered := byID["zone-2"]; answered {
		t.Error("a request that is no longer REQUESTED must not be answered again")
	}
	if err := vda5050.ValidateRaw(vda5050.TopicResponses, sent[0].Payload); err != nil {
		t.Errorf("responses message does not conform: %v", err)
	}
}

func TestConnectionStateTransitions(t *testing.T) {
	b := newFakeBroker()
	var seen []vda5050.ConnectionState
	var mu sync.Mutex
	f := NewFleet(b, Options{
		OnConnection: func(_ *Vehicle, s vda5050.ConnectionState) {
			mu.Lock()
			seen = append(seen, s)
			mu.Unlock()
		},
	})
	ctx := context.Background()
	_, _ = f.Register(ctx, testID)
	_ = f.Start(ctx)

	for _, s := range []vda5050.ConnectionState{
		vda5050.ConnectionStateOnline,
		vda5050.ConnectionStateConnectionBroken,
		vda5050.ConnectionStateOnline,
		vda5050.ConnectionStateOffline,
	} {
		b.deliver("vda5050/v3/KIT/0001/connection", connectionJSON(t, s))
	}

	mu.Lock()
	defer mu.Unlock()
	if len(seen) != 4 {
		t.Fatalf("want 4 connection callbacks, got %d", len(seen))
	}
	if f.Vehicle(testID).ConnectionState() != vda5050.ConnectionStateOffline {
		t.Error("the vehicle should end up OFFLINE")
	}
}

func TestForeignTrafficIsIgnored(t *testing.T) {
	// Discovery on a shared broker must not adopt messages from another
	// system, or a vehicle belonging to someone else joins the fleet.
	b := newFakeBroker()
	f := NewFleet(b, Options{Discover: true})
	ctx := context.Background()
	if err := f.Start(ctx); err != nil {
		t.Fatalf("Start: %v", err)
	}

	b.deliver("someothersystem/telemetry/robot1", []byte(`{"x":1}`))
	if n := len(f.Vehicles()); n != 0 {
		t.Errorf("no vehicle should have been created from foreign traffic, got %d", n)
	}

	b.deliver("vda5050/v3/ACME/9/state", stateJSON(t, vda5050.State{OperatingMode: vda5050.OperatingModeAutomatic}))
	if n := len(f.Vehicles()); n != 1 {
		t.Errorf("discovery should have adopted the VDA5050 vehicle, got %d", n)
	}
}

func TestDisabledFleetIsInert(t *testing.T) {
	// The path a deployment with no MQTT broker takes: everything must be
	// callable and do nothing, so herdIQ needs no conditional wiring.
	f := NewFleet(transport.Disabled{}, Options{})
	if f.Enabled() {
		t.Fatal("a fleet over a disabled broker must report itself disabled")
	}
	if err := f.Start(context.Background()); err != nil {
		t.Errorf("Start on a disabled fleet should be a silent no-op, got %v", err)
	}
	if n := len(f.Vehicles()); n != 0 {
		t.Errorf("want no vehicles, got %d", n)
	}
	if err := f.Stop(); err != nil {
		t.Errorf("Stop: %v", err)
	}
}
