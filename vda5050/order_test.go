package vda5050

import (
	"encoding/json"
	"errors"
	"testing"
)

// route builds a straight n-node route with edges between consecutive nodes.
func route(t *testing.T, n int) Route {
	t.Helper()
	r := Route{}
	for i := 0; i < n; i++ {
		r.Nodes = append(r.Nodes, RouteNode{
			NodeID:   nodeName(i),
			Position: &NodePosition{X: float64(i), Y: 0, MapID: "level_1"},
		})
		if i > 0 {
			r.Edges = append(r.Edges, RouteEdge{EdgeID: edgeName(i - 1)})
		}
	}
	if err := r.Validate(); err != nil {
		t.Fatalf("test route is invalid: %v", err)
	}
	return r
}

func nodeName(i int) string { return string(rune('a' + i)) }
func edgeName(i int) string { return "e" + string(rune('1'+i)) }

func newTracker() *OrderTracker {
	return NewOrderTracker(Identity{Manufacturer: "GOAT", SerialNumber: "0001"}, NewHeaderCounter())
}

func TestStartSplitsBaseAndHorizon(t *testing.T) {
	tr := newTracker()
	// 5 nodes, released through index 2: a,b,c are base; d,e are horizon.
	o, err := tr.Start(route(t, 5), 2)
	if err != nil {
		t.Fatalf("Start: %v", err)
	}

	if o.OrderUpdateID != 0 {
		t.Errorf("a new order must start at orderUpdateId 0, got %d", o.OrderUpdateID)
	}
	if len(o.Nodes) != 5 {
		t.Fatalf("want all 5 nodes in the first order, got %d", len(o.Nodes))
	}
	wantReleased := []bool{true, true, true, false, false}
	for i, n := range o.Nodes {
		if n.Released != wantReleased[i] {
			t.Errorf("node %d (%s): released=%v, want %v", i, n.NodeID, n.Released, wantReleased[i])
		}
		if got := n.SequenceID; got != uint32(2*i) {
			t.Errorf("node %d: sequenceId=%d, want %d", i, got, 2*i)
		}
	}

	// An edge is part of the base only when the node it leads to is: edge 2
	// leads from c to d, and d is horizon, so edge 2 must not be released.
	wantEdgeReleased := []bool{true, true, false, false}
	for i, e := range o.Edges {
		if e.Released != wantEdgeReleased[i] {
			t.Errorf("edge %d (%s): released=%v, want %v", i, e.EdgeID, e.Released, wantEdgeReleased[i])
		}
		if got := e.SequenceID; got != uint32(2*i+1) {
			t.Errorf("edge %d: sequenceId=%d, want %d", i, got, 2*i+1)
		}
	}
}

func TestExtendStartsAtTheDecisionPoint(t *testing.T) {
	tr := newTracker()
	r := route(t, 5)
	if _, err := tr.Start(r, 2); err != nil {
		t.Fatalf("Start: %v", err)
	}

	upd, err := tr.Extend(r, 4)
	if err != nil {
		t.Fatalf("Extend: %v", err)
	}

	if upd.OrderUpdateID != 1 {
		t.Errorf("orderUpdateId must increment to 1, got %d", upd.OrderUpdateID)
	}
	if upd.OrderID != tr.OrderID() {
		t.Errorf("an update must keep the same orderId")
	}
	// §6.1.2: the update begins with the previous decision point (index 2),
	// resent unchanged, and does not repeat the base before it.
	if len(upd.Nodes) != 3 {
		t.Fatalf("want nodes c,d,e in the update, got %d", len(upd.Nodes))
	}
	if upd.Nodes[0].NodeID != nodeName(2) {
		t.Errorf("update must start at the decision point %q, got %q", nodeName(2), upd.Nodes[0].NodeID)
	}
	if upd.Nodes[0].SequenceID != 4 {
		t.Errorf("the stitching node keeps its sequenceId 4, got %d", upd.Nodes[0].SequenceID)
	}
	for i, n := range upd.Nodes {
		if !n.Released {
			t.Errorf("node %d (%s) should be released after extending through index 4", i, n.NodeID)
		}
	}
}

func TestSequenceIDsSurviveOrderUpdates(t *testing.T) {
	tr := newTracker()
	r := route(t, 6)
	first, _ := tr.Start(r, 1)
	second, err := tr.Extend(r, 3)
	if err != nil {
		t.Fatalf("Extend: %v", err)
	}

	// §6.1.2 forbids a released node's sequenceId from changing across
	// updates; a vehicle uses it to detect a missed message.
	seq := map[string]uint32{}
	for _, n := range first.Nodes {
		seq[n.NodeID] = n.SequenceID
	}
	for _, n := range second.Nodes {
		if prev, ok := seq[n.NodeID]; ok && prev != n.SequenceID {
			t.Errorf("node %s sequenceId changed from %d to %d across an update", n.NodeID, prev, n.SequenceID)
		}
	}
}

func TestExtendRejectsChangesToTheBase(t *testing.T) {
	tr := newTracker()
	r := route(t, 5)
	if _, err := tr.Start(r, 3); err != nil {
		t.Fatalf("Start: %v", err)
	}

	// Rewrite a node inside the released base.
	changed := route(t, 5)
	changed.Nodes[1].NodeID = "rerouted"

	if _, err := tr.Extend(changed, 4); !errors.Is(err, ErrBaseChanged) {
		t.Fatalf("want ErrBaseChanged when the base is rewritten, got %v", err)
	}
}

func TestExtendAllowsReplacingTheHorizon(t *testing.T) {
	tr := newTracker()
	if _, err := tr.Start(route(t, 5), 2); err != nil {
		t.Fatalf("Start: %v", err)
	}

	// Same base (a,b,c), a different route beyond the decision point. §6.1.2
	// explicitly permits this: only the base is immutable.
	replanned := route(t, 4)
	replanned.Nodes[3].NodeID = "detour"

	upd, err := tr.Extend(replanned, 3)
	if err != nil {
		t.Fatalf("replacing the horizon should be allowed, got %v", err)
	}
	if upd.Nodes[len(upd.Nodes)-1].NodeID != "detour" {
		t.Errorf("the new horizon should end at the re-planned node")
	}
}

func TestExtendRejectsARetreatingDecisionPoint(t *testing.T) {
	tr := newTracker()
	r := route(t, 5)
	if _, err := tr.Start(r, 3); err != nil {
		t.Fatalf("Start: %v", err)
	}
	if _, err := tr.Extend(r, 2); err == nil {
		t.Fatal("want an error when releaseThrough moves backwards: the base cannot be un-released")
	}
}

func TestCancelBlocksFurtherUpdates(t *testing.T) {
	tr := newTracker()
	r := route(t, 4)
	if _, err := tr.Start(r, 1); err != nil {
		t.Fatalf("Start: %v", err)
	}

	action, err := tr.Cancel()
	if err != nil {
		t.Fatalf("Cancel: %v", err)
	}
	if action.ActionType != ActionCancelOrder {
		t.Errorf("want a cancelOrder action, got %q", action.ActionType)
	}
	if action.BlockingType != InstantActionBlockingTypeNone {
		t.Errorf("an instant action's blockingType is always NONE, got %q", action.BlockingType)
	}
	// §6.1.3.1: no further updates may be sent for a cancelled order.
	if _, err := tr.Extend(r, 3); err == nil {
		t.Fatal("want an error when extending a cancelled order")
	}
}

func TestReconcileDetectsALostUpdate(t *testing.T) {
	tr := newTracker()
	r := route(t, 4)
	o, _ := tr.Start(r, 1)
	if _, err := tr.Extend(r, 3); err != nil {
		t.Fatalf("Extend: %v", err)
	}

	// The vehicle is still reporting the first message: the update was lost.
	s := &State{OrderID: o.OrderID, OrderUpdateID: 0}
	ok, reason := tr.Reconcile(s)
	if ok {
		t.Fatal("want a mismatch when the vehicle is behind on orderUpdateId")
	}
	if reason == "" {
		t.Error("a mismatch should explain itself")
	}

	// Caught up.
	s.OrderUpdateID = 1
	if ok, reason := tr.Reconcile(s); !ok {
		t.Errorf("want a match once the vehicle catches up, got %q", reason)
	}
}

func TestEmptyActionsSerialiseAsAnArray(t *testing.T) {
	tr := newTracker()
	o, err := tr.Start(route(t, 2), 1)
	if err != nil {
		t.Fatalf("Start: %v", err)
	}
	raw, err := json.Marshal(o)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	// `actions` is required, so a node with none must emit [] and not null.
	// Vehicles validating against the schema reject null here.
	var decoded map[string]any
	if err := json.Unmarshal(raw, &decoded); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	nodes := decoded["nodes"].([]any)
	for i, n := range nodes {
		actions, present := n.(map[string]any)["actions"]
		if !present || actions == nil {
			t.Errorf("node %d: actions must be an empty array, not null or absent", i)
		}
	}
}

func TestSingleNodeRouteIsValid(t *testing.T) {
	// An order that only performs an action where the vehicle already stands
	// has one node and no edges.
	tr := newTracker()
	r := Route{Nodes: []RouteNode{{
		NodeID:   "station-1",
		Position: &NodePosition{X: 1, Y: 2, MapID: "level_1"},
		Actions:  []Action{NewAction(ActionPick, BlockingTypeHard, Param("loadType", "pallet_eu"))},
	}}}
	o, err := tr.Start(r, 0)
	if err != nil {
		t.Fatalf("a single-node order must be valid: %v", err)
	}
	if len(o.Edges) != 0 {
		t.Errorf("want no edges, got %d", len(o.Edges))
	}
	if len(o.Nodes[0].Actions) != 1 {
		t.Errorf("want the pick action carried through")
	}
}

func TestRouteValidateCatchesDuplicateActionIDs(t *testing.T) {
	// actionId is the only key linking an action to its actionState, so a
	// duplicate makes progress impossible to attribute.
	a := NewAction(ActionPick, BlockingTypeHard)
	r := Route{Nodes: []RouteNode{
		{NodeID: "a", Actions: []Action{a}},
		{NodeID: "b", Actions: []Action{a}},
	}, Edges: []RouteEdge{{EdgeID: "e1"}}}

	if err := r.Validate(); err == nil {
		t.Fatal("want an error for a duplicated actionId")
	}
}

func TestOrderCompleteRequiresTerminalActions(t *testing.T) {
	s := &State{
		NodeStates: []NodeState{},
		EdgeStates: []EdgeState{},
		ActionStates: []ActionState{
			{ActionID: "1", ActionStatus: ActionStatusFinished},
			{ActionID: "2", ActionStatus: ActionStatusRunning},
		},
	}
	if OrderComplete(s) {
		t.Error("an order with a RUNNING action is not complete")
	}

	// RETRIABLE is not terminal: it is waiting on a fleet-control decision.
	s.ActionStates[1].ActionStatus = ActionStatusRetriable
	if OrderComplete(s) {
		t.Error("an order with a RETRIABLE action is not complete; it awaits retry or skipRetry")
	}

	s.ActionStates[1].ActionStatus = ActionStatusFailed
	if !OrderComplete(s) {
		t.Error("FAILED is terminal, so the order is complete")
	}
}
