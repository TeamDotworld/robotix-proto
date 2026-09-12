package vda5050

import (
	"errors"
	"testing"
)

// route builds a straight run of n nodes with n-1 edges, named n0..n{n-1}.
func straightRoute(n int) Route {
	r := Route{}
	for i := 0; i < n; i++ {
		r.Nodes = append(r.Nodes, RouteNode{
			NodeID:   nodeName(i),
			Position: &NodePosition{X: float64(i), Y: 0, MapID: "level_1"},
		})
	}
	for i := 0; i < n-1; i++ {
		r.Edges = append(r.Edges, RouteEdge{EdgeID: "e" + nodeName(i)})
	}
	return r
}

func seqOfNode(o *Order, nodeID string) []uint32 {
	var out []uint32
	for _, n := range o.Nodes {
		if n.NodeID == nodeID {
			out = append(out, n.SequenceID)
		}
	}
	return out
}

func TestOrdinaryRouteKeepsTheClassicNumbering(t *testing.T) {
	// Nothing about the new assignable scheme may change the numbering of a
	// plain route: nodes on even sequenceIds, edges on the odd one between
	// their endpoints (§7.3).
	tr := newTracker()
	o, err := tr.Start(straightRoute(4), 3)
	if err != nil {
		t.Fatalf("Start: %v", err)
	}
	for i, n := range o.Nodes {
		if want := uint32(2 * i); n.SequenceID != want {
			t.Errorf("node %d (%s) sequenceId = %d, want %d", i, n.NodeID, n.SequenceID, want)
		}
	}
	for i, e := range o.Edges {
		if want := uint32(2*i + 1); e.SequenceID != want {
			t.Errorf("edge %d (%s) sequenceId = %d, want %d", i, e.EdgeID, e.SequenceID, want)
		}
	}
}

func TestReleasedSequenceIDsSurviveAnExtension(t *testing.T) {
	// §6.1.2: "Once a sequenceId is assigned and the node is released, it does
	// not change with order updates."
	tr := newTracker()
	base := straightRoute(5)
	first, err := tr.Start(base, 2)
	if err != nil {
		t.Fatalf("Start: %v", err)
	}
	before := map[string]uint32{}
	for _, n := range first.Nodes {
		if n.Released {
			before[n.NodeID] = n.SequenceID
		}
	}

	longer := straightRoute(7)
	second, err := tr.Extend(longer, 4)
	if err != nil {
		t.Fatalf("Extend: %v", err)
	}
	for _, n := range second.Nodes {
		if was, ok := before[n.NodeID]; ok && n.SequenceID != was {
			t.Errorf("released node %s moved from sequenceId %d to %d", n.NodeID, was, n.SequenceID)
		}
	}
}

func TestStitchingNodeReproducesFigure7(t *testing.T) {
	// Figure 7 of the specification, reproduced exactly.
	//
	// Initial order: d(2) g(4) b(6) h(8), decision point at g.
	// The update must resend g at sequenceId 4 unchanged, insert a second g
	// at sequenceId 6 carrying the newly released actions, and shift b to 8
	// and h to 10. The joining edge sits at 5.
	tr := newTracker()
	r := Route{
		Nodes: []RouteNode{
			{NodeID: "d", Position: &NodePosition{X: 0, MapID: "m"}},
			{NodeID: "g", Position: &NodePosition{X: 1, MapID: "m"}},
			{NodeID: "b", Position: &NodePosition{X: 2, MapID: "m"}},
			{NodeID: "h", Position: &NodePosition{X: 3, MapID: "m"}},
		},
		Edges: []RouteEdge{{EdgeID: "e3"}, {EdgeID: "e8"}, {EdgeID: "e9"}},
	}
	// Shift the whole thing so d starts at 2, as the figure has it, by
	// starting one node earlier and extending. Simpler: assert relative to
	// what Start produces, which puts d at 0 and g at 2.
	first, err := tr.Start(r, 1) // decision point at g
	if err != nil {
		t.Fatalf("Start: %v", err)
	}
	gSeq := seqOfNode(first, "g")
	if len(gSeq) != 1 || gSeq[0] != 2 {
		t.Fatalf("decision node g should be at sequenceId 2, got %v", gSeq)
	}

	pick := NewAction(ActionPick, BlockingTypeHard, Param("lhd", "LHD1"))
	upd, err := tr.ReleaseActionsAtDecisionPoint("g-actions", []Action{pick}, 1)
	if err != nil {
		t.Fatalf("ReleaseActionsAtDecisionPoint: %v", err)
	}

	if upd.OrderID != first.OrderID {
		t.Errorf("orderId changed: %q -> %q", first.OrderID, upd.OrderID)
	}
	if upd.OrderUpdateID != 1 {
		t.Errorf("orderUpdateId = %d, want 1", upd.OrderUpdateID)
	}

	// The update begins with the decision node, resent unchanged.
	if len(upd.Nodes) == 0 || upd.Nodes[0].NodeID != "g" || upd.Nodes[0].SequenceID != 2 {
		t.Fatalf("update must begin with the decision node g at sequenceId 2, got %+v", upd.Nodes[0])
	}
	if len(upd.Nodes[0].Actions) != 0 {
		t.Errorf("the decision node must be resent with its original (empty) action list, got %d actions",
			len(upd.Nodes[0].Actions))
	}

	// The stitching node lands on decision + 2 and carries the new actions.
	stitch := upd.Nodes[1]
	if stitch.SequenceID != 4 {
		t.Errorf("stitching node sequenceId = %d, want decision+2 = 4", stitch.SequenceID)
	}
	if !stitch.Released {
		t.Error("the stitching node must be released; releasing its actions is the point")
	}
	if len(stitch.Actions) != 1 || stitch.Actions[0].ActionType != ActionPick {
		t.Errorf("stitching node should carry the pick action, got %+v", stitch.Actions)
	}
	if stitch.NodePosition == nil || upd.Nodes[0].NodePosition == nil ||
		stitch.NodePosition.X != upd.Nodes[0].NodePosition.X {
		t.Error("the stitching node must sit at the same position as the decision node")
	}

	// The joining edge sits between them, at decision + 1.
	if len(upd.Edges) == 0 || upd.Edges[0].SequenceID != 3 {
		t.Fatalf("joining edge should be at sequenceId 3, got %+v", upd.Edges[0])
	}
	if len(upd.Edges[0].Actions) != 0 {
		t.Error("the joining edge must carry no actions")
	}

	// Everything downstream shifted by exactly 2.
	for _, n := range upd.Nodes[2:] {
		switch n.NodeID {
		case "b":
			if n.SequenceID != 6 {
				t.Errorf("b should shift to sequenceId 6, got %d", n.SequenceID)
			}
		case "h":
			if n.SequenceID != 8 {
				t.Errorf("h should shift to sequenceId 8, got %d", n.SequenceID)
			}
		}
	}

	// Node and edge sequenceIds interleave and strictly increase across the
	// whole message (§6.1.1): node, edge, node, edge, ...
	var last int64 = -1
	for i := range upd.Nodes {
		if int64(upd.Nodes[i].SequenceID) <= last {
			t.Errorf("node sequenceIds must strictly increase, saw %d after %d", upd.Nodes[i].SequenceID, last)
		}
		last = int64(upd.Nodes[i].SequenceID)
		if i < len(upd.Edges) {
			if got, want := upd.Edges[i].SequenceID, upd.Nodes[i].SequenceID+1; got != want {
				t.Errorf("edge %d sequenceId = %d, want %d (between its endpoints)", i, got, want)
			}
		}
	}

	if err := Validate(TopicOrder, upd); err != nil {
		t.Errorf("the stitched update does not conform to the schema: %v", err)
	}
}

func TestStitchedRouteCanStillBeExtended(t *testing.T) {
	// After a stitch the tracker's own route is authoritative; a caller that
	// refreshes from Route() must be able to carry on extending.
	tr := newTracker()
	if _, err := tr.Start(straightRoute(4), 1); err != nil {
		t.Fatalf("Start: %v", err)
	}
	if _, err := tr.ReleaseActionsAtDecisionPoint("", []Action{NewAction(ActionDrop, BlockingTypeSoft)}, 1); err != nil {
		t.Fatalf("stitch: %v", err)
	}
	r := tr.Route()
	if len(r.Nodes) != 5 {
		t.Fatalf("route should have grown to 5 nodes, got %d", len(r.Nodes))
	}
	if _, err := tr.Extend(r, 3); err != nil {
		t.Fatalf("Extend after stitch: %v", err)
	}
}

func TestBaseImmutabilityCatchesChangedActions(t *testing.T) {
	// An id-only check passes this; the base is immutable in its entirety.
	tr := newTracker()
	r := straightRoute(4)
	r.Nodes[1].Actions = []Action{NewAction(ActionPick, BlockingTypeHard)}
	if _, err := tr.Start(r, 2); err != nil {
		t.Fatalf("Start: %v", err)
	}

	tampered := tr.Route()
	tampered.Nodes[1].Actions = []Action{NewAction(ActionDrop, BlockingTypeHard)}
	_, err := tr.Extend(tampered, 3)
	if !errors.Is(err, ErrBaseChanged) {
		t.Fatalf("changing a released node's actions must be rejected, got %v", err)
	}
}

func TestBaseImmutabilityCatchesChangedEdgeSpeed(t *testing.T) {
	tr := newTracker()
	r := straightRoute(4)
	r.Edges[0].MaximumSpeed = Float(1.0)
	if _, err := tr.Start(r, 2); err != nil {
		t.Fatalf("Start: %v", err)
	}
	tampered := tr.Route()
	tampered.Edges[0].MaximumSpeed = Float(2.0)
	if _, err := tr.Extend(tampered, 3); !errors.Is(err, ErrBaseChanged) {
		t.Fatalf("changing a released edge's maximumSpeed must be rejected, got %v", err)
	}
}

func TestHorizonMayBeReplacedFreely(t *testing.T) {
	// The other half of the same rule: everything past the decision point is
	// the fleet control's to change, including deleting it entirely.
	tr := newTracker()
	if _, err := tr.Start(straightRoute(5), 1); err != nil {
		t.Fatalf("Start: %v", err)
	}
	replaced := tr.Route()
	replaced.Nodes = replaced.Nodes[:3]
	replaced.Edges = replaced.Edges[:2]
	replaced.Nodes[2].Actions = []Action{NewAction(ActionDetectObject, BlockingTypeNone)}
	if _, err := tr.Extend(replaced, 2); err != nil {
		t.Fatalf("replacing the horizon must be allowed: %v", err)
	}
}
