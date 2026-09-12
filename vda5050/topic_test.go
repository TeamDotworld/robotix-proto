package vda5050

import "testing"

func TestBuildFollowsTheSuggestedLayout(t *testing.T) {
	s := NewTopicScheme("")
	id := Identity{Manufacturer: "KIT", SerialNumber: "0001"}

	if got, want := s.Build(id, TopicOrder), "vda5050/v3/KIT/0001/order"; got != want {
		t.Errorf("Build = %q, want %q", got, want)
	}
}

func TestBuildWithABrokerPrefix(t *testing.T) {
	// Cloud brokers impose their own root, which is why the leading levels
	// are configurable while the final one is not.
	s := NewTopicScheme("/plant-a/")
	id := Identity{Manufacturer: "KIT", SerialNumber: "0001"}

	if got, want := s.Build(id, TopicState), "plant-a/vda5050/v3/KIT/0001/state"; got != want {
		t.Errorf("Build = %q, want %q", got, want)
	}
}

func TestParseRoundTrip(t *testing.T) {
	for _, prefix := range []string{"", "plant-a", "a/b"} {
		s := NewTopicScheme(prefix)
		id := Identity{Manufacturer: "GOAT", SerialNumber: "amr-07"}
		for _, topic := range AllTopics {
			built := s.Build(id, topic)
			gotID, gotTopic, err := s.Parse(built)
			if err != nil {
				t.Fatalf("Parse(%q): %v", built, err)
			}
			if gotID != id || gotTopic != topic {
				t.Errorf("round trip of %q gave %v/%v", built, gotID, gotTopic)
			}
		}
	}
}

func TestParseRejectsForeignTraffic(t *testing.T) {
	// On a shared plant broker, unrelated messages must be ignored rather
	// than misread as vehicle data.
	s := NewTopicScheme("")
	for _, topic := range []string{
		"some/other/system/state",
		"vda5050/v2/KIT/0001/state", // wrong major version
		"vda5050/v3/KIT/0001/telemetry",
		"vda5050/v3/KIT/0001",
		"vda5050/v3/KIT/0001/state/extra",
	} {
		if _, _, err := s.Parse(topic); err == nil {
			t.Errorf("Parse(%q) should have failed", topic)
		}
	}
}

func TestSubscribeAllUsesWildcards(t *testing.T) {
	s := NewTopicScheme("")
	if got, want := s.SubscribeAll(TopicState), "vda5050/v3/+/+/state"; got != want {
		t.Errorf("SubscribeAll = %q, want %q", got, want)
	}
	if got, want := s.SubscribeManufacturer("KIT", TopicConnection), "vda5050/v3/KIT/+/connection"; got != want {
		t.Errorf("SubscribeManufacturer = %q, want %q", got, want)
	}
}

func TestIdentityRejectsTopicBreakingCharacters(t *testing.T) {
	// '/' would create extra topic levels; '+' and '#' are wildcards, so a
	// vehicle named "#" would silently receive every other vehicle's orders.
	for _, id := range []Identity{
		{Manufacturer: "a/b", SerialNumber: "0001"},
		{Manufacturer: "KIT", SerialNumber: "00+1"},
		{Manufacturer: "KIT", SerialNumber: "#"},
		{Manufacturer: "$SYS", SerialNumber: "0001"},
		{Manufacturer: "", SerialNumber: "0001"},
		{Manufacturer: "KIT", SerialNumber: ""},
	} {
		if err := id.Valid(); err == nil {
			t.Errorf("Identity%+v should be rejected", id)
		}
	}

	// The characters the specification explicitly permits.
	ok := Identity{Manufacturer: "GOAT-Robotics", SerialNumber: "amr_07.a:1-x"}
	if err := ok.Valid(); err != nil {
		t.Errorf("Identity%+v should be accepted: %v", ok, err)
	}
}

func TestQoSAndRetainMatchTheSpecification(t *testing.T) {
	// §4.1: QoS 0 everywhere except connection, which is QoS 1.
	// §6.5: connection is retained, so a fleet control learns connectivity on
	// subscribe. §6.10: "all messages on this topic shall be sent with a
	// retained flag" for factsheet, for the same reason -- a fleet control
	// joining later must learn the vehicle's capabilities without asking.
	for _, topic := range AllTopics {
		wantQoS := byte(0)
		wantRetain := topic == TopicFactsheet
		if topic == TopicConnection {
			wantQoS, wantRetain = 1, true
		}
		if got := topic.QoS(); got != wantQoS {
			t.Errorf("%s QoS = %d, want %d", topic, got, wantQoS)
		}
		if got := topic.Retained(); got != wantRetain {
			t.Errorf("%s Retained = %v, want %v", topic, got, wantRetain)
		}
	}
}

func TestHeaderCounterIsPerVehicleAndPerTopic(t *testing.T) {
	h := NewHeaderCounter()
	a := Identity{Manufacturer: "GOAT", SerialNumber: "1"}
	b := Identity{Manufacturer: "GOAT", SerialNumber: "2"}

	if got := h.Next(a, TopicOrder); got != 0 {
		t.Errorf("first headerId should be 0, got %d", got)
	}
	if got := h.Next(a, TopicOrder); got != 1 {
		t.Errorf("second headerId should be 1, got %d", got)
	}
	// A different topic on the same vehicle counts independently, so a
	// receiver can spot a gap on one topic without traffic on another
	// confusing the picture.
	if got := h.Next(a, TopicInstantActions); got != 0 {
		t.Errorf("a new topic starts at 0, got %d", got)
	}
	if got := h.Next(b, TopicOrder); got != 0 {
		t.Errorf("a different vehicle starts at 0, got %d", got)
	}
}
