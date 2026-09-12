package vda5050

import (
	"encoding/json"
	"strings"
	"testing"
	"time"
)

// Every message this package builds is validated against the official
// VDA5050 3.0.0 JSON schemas. These tests are the reason the schemas are
// embedded rather than merely used at code-generation time: they prove that
// what herdIQ puts on the wire is what a third-party vehicle will accept.

func TestOrderValidatesAgainstTheOfficialSchema(t *testing.T) {
	tr := newTracker()
	r := Route{
		Nodes: []RouteNode{
			{
				NodeID:   "start",
				Position: &NodePosition{X: 0, Y: 0, MapID: "level_1", Theta: Float(0)},
			},
			{
				NodeID:   "pick-station",
				Position: &NodePosition{X: 10.5, Y: -3.25, MapID: "level_1", AllowedDeviationXY: &AllowedDeviationXY{A: 0.05, B: 0.05}},
				Actions: []Action{
					NewAction(ActionPick, BlockingTypeHard,
						Param("lhd", "LHD1"),
						Param("stationType", "floor"),
						Param("loadType", "pallet_eu"),
						Param("height", 0.12),
					),
				},
			},
		},
		Edges: []RouteEdge{{
			EdgeID:          "start->pick",
			MaximumSpeed:    Float(1.2),
			OrientationType: orientation(OrientationTypeTangential),
			Length:          Float(10.9),
		}},
	}

	o, err := tr.Start(r, 1)
	if err != nil {
		t.Fatalf("Start: %v", err)
	}
	if err := Validate(TopicOrder, o); err != nil {
		t.Fatalf("generated order does not conform to the specification: %v", err)
	}
}

func orientation(v OrientationType) *OrientationType { return &v }

func TestInstantActionsValidateAgainstTheOfficialSchema(t *testing.T) {
	h := NewHeaderCounter()
	id := Identity{Manufacturer: "GOAT", SerialNumber: "0001"}

	for name, action := range map[string]InstantAction{
		"cancelOrder":      CancelOrder("order-123"),
		"startPause":       StartPause(),
		"stopPause":        StopPause(),
		"stateRequest":     StateRequest(),
		"factsheetRequest": FactsheetRequest(),
		"initializePos":    InitializePosition(1.5, 2.5, 0.1, "level_1", "node-7"),
		"retry":            Retry("action-42"),
		"skipRetry":        SkipRetry("action-42"),
	} {
		msg := &InstantActions{Actions: []InstantAction{action}}
		h.NewHeader(id, TopicInstantActions).ApplyInstantActions(msg)
		if err := Validate(TopicInstantActions, msg); err != nil {
			t.Errorf("%s does not conform: %v", name, err)
		}
	}
}

func TestResponsesValidateAgainstTheOfficialSchema(t *testing.T) {
	h := NewHeaderCounter()
	id := Identity{Manufacturer: "GOAT", SerialNumber: "0001"}

	msg := &Responses{Responses: []Response{
		{RequestID: "r1", GrantType: GrantTypeGranted, LeaseExpiry: Str(Timestamp(time.Now().Add(time.Minute)))},
		{RequestID: "r2", GrantType: GrantTypeQueued},
		{RequestID: "r3", GrantType: GrantTypeRejected},
	}}
	h.NewHeader(id, TopicResponses).ApplyResponses(msg)

	if err := Validate(TopicResponses, msg); err != nil {
		t.Fatalf("responses message does not conform: %v", err)
	}
}

func TestSchemaValidationRejectsAMalformedOrder(t *testing.T) {
	// orderId is required; leaving it out must be caught. This guards the
	// validator itself: a validator that accepts everything would make the
	// tests above meaningless.
	bad := []byte(`{
		"headerId": 1, "timestamp": "2026-08-06T10:00:00.000Z", "version": "3.0.0",
		"manufacturer": "GOAT", "serialNumber": "0001",
		"orderUpdateId": 0, "nodes": [], "edges": []
	}`)
	err := ValidateRaw(TopicOrder, bad)
	if err == nil {
		t.Fatal("want a validation error for an order with no orderId")
	}
	var ve *ValidationError
	if !asValidationError(err, &ve) {
		t.Fatalf("want a *ValidationError, got %T", err)
	}
}

func asValidationError(err error, target **ValidationError) bool {
	v, ok := err.(*ValidationError)
	if ok {
		*target = v
	}
	return ok
}

func TestAllSchemasCompile(t *testing.T) {
	// Catches a schema file that failed to copy or that a future spec bump
	// left in a state this package cannot load.
	for _, topic := range AllTopics {
		if _, err := SchemaFor(topic); err != nil {
			t.Errorf("schema for %s: %v", topic, err)
		}
	}
}

func TestTimestampUsesMillisecondPrecision(t *testing.T) {
	// 3.0.0 moved from 1/100 s to 1/1000 s. A 2.x implementation emitting two
	// fractional digits is no longer compliant, so this is worth pinning.
	ts := Timestamp(time.Date(2026, 8, 6, 11, 40, 3, 123456789, time.UTC))
	if ts != "2026-08-06T11:40:03.123Z" {
		t.Errorf("timestamp = %q, want millisecond precision ending in Z", ts)
	}
	if n := strings.Count(ts, "."); n != 1 {
		t.Errorf("timestamp should have exactly one fractional separator, got %d", n)
	}
}

func TestTimestampRoundTrip(t *testing.T) {
	now := time.Now().UTC().Truncate(time.Millisecond)
	got, err := ParseTimestamp(Timestamp(now))
	if err != nil {
		t.Fatalf("ParseTimestamp: %v", err)
	}
	if !got.Equal(now) {
		t.Errorf("round trip changed the value: %v -> %v", now, got)
	}
}

func TestParseTimestampToleratesOtherPrecisions(t *testing.T) {
	// Vehicles in the field emit 2, 6 or 9 fractional digits. Rejecting those
	// would mean discarding otherwise usable state messages.
	for _, s := range []string{
		"2026-08-06T11:40:03.12Z",
		"2026-08-06T11:40:03.123456Z",
		"2026-08-06T11:40:03Z",
	} {
		if _, err := ParseTimestamp(s); err != nil {
			t.Errorf("ParseTimestamp(%q): %v", s, err)
		}
	}
}

func TestOptionalFieldsAreOmittedNotZeroed(t *testing.T) {
	// VDA5050 distinguishes an absent optional field from one set to zero:
	// an absent maximumSpeed means "no limit", while 0 means "do not move".
	// Emitting a zero where the field should be absent would stop a vehicle.
	e := Edge{EdgeID: "e1", SequenceID: 1, Released: true, Actions: []Action{}}
	raw, err := json.Marshal(e)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var decoded map[string]any
	if err := json.Unmarshal(raw, &decoded); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	for _, absent := range []string{"maximumSpeed", "orientation", "trajectory", "corridor", "maximumRotationSpeed"} {
		if _, present := decoded[absent]; present {
			t.Errorf("optional field %q must be omitted when unset, not serialised as a zero value", absent)
		}
	}
	// Required fields are always present, even at their zero value.
	for _, required := range []string{"edgeId", "sequenceId", "released", "actions"} {
		if _, present := decoded[required]; !present {
			t.Errorf("required field %q must always be serialised", required)
		}
	}
}
