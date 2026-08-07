package vda5050

import (
	"errors"
	"fmt"
	"sync"
)

// Route is a fleet-control plan for one vehicle, before it is split into the
// released "base" and the not-yet-released "horizon".
//
// Nodes and Edges interleave: Edges[i] connects Nodes[i] to Nodes[i+1], so a
// well-formed route with n nodes has exactly n-1 edges. A single-node route
// (n=1, no edges) is valid and is how a pure "do an action here" order is
// expressed.
type Route struct {
	Nodes []RouteNode
	Edges []RouteEdge
}

// RouteNode is one waypoint of a planned route, without the protocol
// bookkeeping (sequenceId, released) that the tracker fills in.
type RouteNode struct {
	NodeID     string
	Descriptor string
	Position   *NodePosition
	Actions    []Action
}

// RouteEdge is one leg of a planned route.
type RouteEdge struct {
	EdgeID     string
	Descriptor string
	Actions    []Action

	MaximumSpeed *float64
	// MaxRotationSpeed follows the JSON schema, which names this field
	// `maxRotationSpeed`. Table 7.3 of the prose spells it
	// `maximumRotationSpeed`; the schema wins, because that is what goes on
	// the wire and what vehicles validate against.
	MaxRotationSpeed                *float64
	MaximumMobileRobotHeight        *float64
	MinimumLoadHandlingDeviceHeight *float64
	Orientation                     *float64
	OrientationType                 *OrientationType
	Direction                       *string
	ReachOrientationBeforeEntering  *bool
	Length                          *float64
	Trajectory                      *Trajectory
	Corridor                        *Corridor
}

// Validate checks the structural invariant that ties nodes to edges.
func (r Route) Validate() error {
	if len(r.Nodes) == 0 {
		return errors.New("vda5050: route has no nodes")
	}
	if len(r.Edges) != len(r.Nodes)-1 {
		return fmt.Errorf("vda5050: route has %d nodes and %d edges, want %d edges",
			len(r.Nodes), len(r.Edges), len(r.Nodes)-1)
	}
	seen := make(map[string]struct{}, len(r.Nodes))
	for _, n := range r.Nodes {
		if n.NodeID == "" {
			return errors.New("vda5050: route contains a node with an empty nodeId")
		}
		for _, a := range n.Actions {
			if a.ActionID == "" {
				return fmt.Errorf("vda5050: node %q has an action with an empty actionId", n.NodeID)
			}
			if _, dup := seen[a.ActionID]; dup {
				return fmt.Errorf("vda5050: actionId %q appears more than once in the route", a.ActionID)
			}
			seen[a.ActionID] = struct{}{}
		}
	}
	for _, e := range r.Edges {
		if e.EdgeID == "" {
			return errors.New("vda5050: route contains an edge with an empty edgeId")
		}
		for _, a := range e.Actions {
			if _, dup := seen[a.ActionID]; dup {
				return fmt.Errorf("vda5050: actionId %q appears more than once in the route", a.ActionID)
			}
			seen[a.ActionID] = struct{}{}
		}
	}
	return nil
}

// Sequence IDs are shared between nodes and edges and encode traversal order
// (§7.3). Interleaving them as node=2i, edge=2i+1 gives a stable numbering
// that never has to be recomputed when the horizon changes — which matters,
// because the specification forbids a released node's sequenceId from moving.
func nodeSeq(i int) uint32 { return uint32(2 * i) }
func edgeSeq(i int) uint32 { return uint32(2*i + 1) }

// ErrBaseChanged is returned when a caller tries to alter part of the route
// that has already been released to the vehicle. The base is immutable
// (§6.1.2): MQTT gives no delivery guarantee, so fleet control must assume
// the vehicle is already acting on everything it released. Recovering from
// this means cancelling the order and issuing a new one.
var ErrBaseChanged = errors.New("vda5050: cannot modify the released base of an active order")

// OrderTracker owns the order lifecycle for a single vehicle: it assigns
// orderIds and orderUpdateIds, splits the plan into base and horizon, and
// builds the stitching node that every order update must begin with.
//
// An OrderTracker is safe for concurrent use.
type OrderTracker struct {
	mu sync.Mutex

	id      Identity
	headers *HeaderCounter

	orderID  string
	updateID uint32
	active   bool

	route Route
	// released is the index of the last node released to the vehicle: the
	// decision point. -1 before the first order is issued.
	released int
	// sent is the decision-point index carried by the last message published.
	// The next update must begin with this node, resent unchanged.
	sent int
	// cancelled records that a cancelOrder was issued; no further updates may
	// be sent for this orderId (§6.1.3.1).
	cancelled bool
}

// NewOrderTracker returns a tracker for one vehicle. The headers counter is
// normally shared across all vehicles handled by the same fleet control.
func NewOrderTracker(id Identity, headers *HeaderCounter) *OrderTracker {
	if headers == nil {
		headers = NewHeaderCounter()
	}
	return &OrderTracker{id: id, headers: headers, released: -1, sent: -1}
}

// OrderID returns the identifier of the active order, or "" when idle.
func (t *OrderTracker) OrderID() string {
	t.mu.Lock()
	defer t.mu.Unlock()
	if !t.active {
		return ""
	}
	return t.orderID
}

// UpdateID returns the orderUpdateId of the most recent message.
func (t *OrderTracker) UpdateID() uint32 {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.updateID
}

// Active reports whether an order is in flight.
func (t *OrderTracker) Active() bool {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.active && !t.cancelled
}

// Start begins a new order. releaseThrough is the index of the last node the
// traffic layer has cleared the vehicle to reach — the decision point. Pass
// len(route.Nodes)-1 to release the whole route, or a smaller index to hold
// the remainder back as horizon.
//
// The vehicle must already be standing on (or within the allowed deviation
// of) route.Nodes[0]; §6.1.4 has the vehicle reject the order otherwise.
func (t *OrderTracker) Start(route Route, releaseThrough int) (*Order, error) {
	if err := route.Validate(); err != nil {
		return nil, err
	}
	if releaseThrough < 0 || releaseThrough >= len(route.Nodes) {
		return nil, fmt.Errorf("vda5050: releaseThrough %d out of range for %d nodes", releaseThrough, len(route.Nodes))
	}

	t.mu.Lock()
	defer t.mu.Unlock()

	t.orderID = NewOrderID()
	t.updateID = 0
	t.active = true
	t.cancelled = false
	t.route = route
	t.released = releaseThrough
	t.sent = releaseThrough

	return t.build(0, releaseThrough), nil
}

// Extend releases more of the current route and/or replaces the horizon.
//
// route must be identical to the tracked route up to and including the
// current decision point; only the horizon beyond it may differ. Anything
// else returns ErrBaseChanged, because the vehicle may already have consumed
// the base. releaseThrough may only grow.
//
// The returned order begins with the previous decision point, resent
// unchanged as the stitching node, which is how the vehicle confirms it did
// not miss an update (§6.1.2).
func (t *OrderTracker) Extend(route Route, releaseThrough int) (*Order, error) {
	if err := route.Validate(); err != nil {
		return nil, err
	}

	t.mu.Lock()
	defer t.mu.Unlock()

	if !t.active {
		return nil, errors.New("vda5050: no active order to extend")
	}
	if t.cancelled {
		return nil, errors.New("vda5050: order has been cancelled; start a new order instead")
	}
	if releaseThrough < t.released {
		return nil, fmt.Errorf("vda5050: releaseThrough %d is behind the current decision point %d", releaseThrough, t.released)
	}
	if releaseThrough >= len(route.Nodes) {
		return nil, fmt.Errorf("vda5050: releaseThrough %d out of range for %d nodes", releaseThrough, len(route.Nodes))
	}
	if err := sameThrough(t.route, route, t.released); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrBaseChanged, err)
	}

	t.route = route
	from := t.sent
	t.released = releaseThrough
	t.sent = releaseThrough
	t.updateID++

	return t.build(from, releaseThrough), nil
}

// build assembles the order message covering nodes [from, end of route],
// marking everything up to releaseThrough as released. Callers hold t.mu.
func (t *OrderTracker) build(from, releaseThrough int) *Order {
	o := &Order{
		OrderID:       t.orderID,
		OrderUpdateID: t.updateID,
		Nodes:         make([]Node, 0, len(t.route.Nodes)-from),
		Edges:         make([]Edge, 0, len(t.route.Edges)),
	}

	for i := from; i < len(t.route.Nodes); i++ {
		rn := t.route.Nodes[i]
		n := Node{
			NodeID:     rn.NodeID,
			SequenceID: nodeSeq(i),
			Released:   i <= releaseThrough,
			// The specification requires an array here, not null, when a node
			// carries no actions.
			Actions: rn.Actions,
		}
		if n.Actions == nil {
			n.Actions = []Action{}
		}
		if rn.Descriptor != "" {
			n.NodeDescriptor = Str(rn.Descriptor)
		}
		if rn.Position != nil {
			p := *rn.Position
			n.NodePosition = &p
		}
		o.Nodes = append(o.Nodes, n)
	}

	// Edge i leads out of node i, so the edges accompanying nodes [from..] are
	// edges [from..]. An edge is released only if the node it leads to is.
	for i := from; i < len(t.route.Edges); i++ {
		re := t.route.Edges[i]
		// 3.0.0 removed startNodeId/endNodeId from the edge: an edge's
		// endpoints are now implied by its sequenceId sitting between the two
		// node sequenceIds. Code ported from a 2.x integration will need those
		// fields deleted rather than renamed.
		e := Edge{
			EdgeID:                          re.EdgeID,
			SequenceID:                      edgeSeq(i),
			Released:                        i+1 <= releaseThrough,
			Actions:                         re.Actions,
			MaximumSpeed:                    re.MaximumSpeed,
			MaxRotationSpeed:                re.MaxRotationSpeed,
			MaximumMobileRobotHeight:        re.MaximumMobileRobotHeight,
			MinimumLoadHandlingDeviceHeight: re.MinimumLoadHandlingDeviceHeight,
			Orientation:                     re.Orientation,
			OrientationType:                 re.OrientationType,
			Direction:                       re.Direction,
			ReachOrientationBeforeEntering:  re.ReachOrientationBeforeEntering,
			Length:                          re.Length,
			Trajectory:                      re.Trajectory,
			Corridor:                        re.Corridor,
		}
		if e.Actions == nil {
			e.Actions = []Action{}
		}
		if re.Descriptor != "" {
			e.EdgeDescriptor = Str(re.Descriptor)
		}
		o.Edges = append(o.Edges, e)
	}

	t.headers.NewHeader(t.id, TopicOrder).Apply(o)
	return o
}

// sameThrough verifies that two routes agree on everything up to and
// including node index n — the part that has already been released.
func sameThrough(old, next Route, n int) error {
	if n < 0 {
		return nil
	}
	if len(next.Nodes) <= n {
		return fmt.Errorf("new route has %d nodes, fewer than the %d already released", len(next.Nodes), n+1)
	}
	for i := 0; i <= n; i++ {
		if old.Nodes[i].NodeID != next.Nodes[i].NodeID {
			return fmt.Errorf("node %d changed from %q to %q", i, old.Nodes[i].NodeID, next.Nodes[i].NodeID)
		}
	}
	for i := 0; i < n; i++ {
		if old.Edges[i].EdgeID != next.Edges[i].EdgeID {
			return fmt.Errorf("edge %d changed from %q to %q", i, old.Edges[i].EdgeID, next.Edges[i].EdgeID)
		}
	}
	return nil
}

// Cancel marks the order cancelled and returns the cancelOrder instant action
// to publish. The caller is responsible for sending it; the tracker only
// stops issuing further updates for this orderId, as §6.1.3.1 requires.
//
// The order is not considered finished until the vehicle reports the
// cancelOrder action as FINISHED; call Finish then.
func (t *OrderTracker) Cancel() (InstantAction, error) {
	t.mu.Lock()
	defer t.mu.Unlock()
	if !t.active {
		return InstantAction{}, errors.New("vda5050: no active order to cancel")
	}
	t.cancelled = true
	return CancelOrder(t.orderID), nil
}

// Finish clears the tracker once the vehicle reports the order complete —
// an empty nodeStates and edgeStates with all actions in a terminal state, or
// a FINISHED cancelOrder. The vehicle keeps reporting the last orderId and
// orderUpdateId after completion, so the fleet control, not the vehicle, is
// what decides an order is done.
func (t *OrderTracker) Finish() {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.active = false
	t.cancelled = false
	t.route = Route{}
	t.released = -1
	t.sent = -1
}

// Reconcile compares a vehicle state message against the tracker and reports
// whether the vehicle is executing the order the fleet control believes it
// is. A mismatch usually means an order or an update was lost in transit, and
// the caller should stop extending and re-plan from the vehicle's actual
// position.
func (t *OrderTracker) Reconcile(s *State) (ok bool, reason string) {
	t.mu.Lock()
	defer t.mu.Unlock()
	if !t.active {
		return true, ""
	}
	if s.OrderID != t.orderID {
		return false, fmt.Sprintf("vehicle reports orderId %q, fleet control holds %q", s.OrderID, t.orderID)
	}
	if s.OrderUpdateID < t.updateID {
		return false, fmt.Sprintf("vehicle is on orderUpdateId %d, fleet control has sent %d", s.OrderUpdateID, t.updateID)
	}
	if s.OrderUpdateID > t.updateID {
		return false, fmt.Sprintf("vehicle reports orderUpdateId %d, ahead of the %d sent by this fleet control", s.OrderUpdateID, t.updateID)
	}
	return true, ""
}

// OrderComplete reports whether a state message shows the order fully
// executed: nothing left to traverse and every action in a terminal state.
func OrderComplete(s *State) bool {
	if len(s.NodeStates) > 0 || len(s.EdgeStates) > 0 {
		return false
	}
	for _, a := range s.ActionStates {
		if !a.ActionStatus.IsTerminal() {
			return false
		}
	}
	return true
}

// NeedsBaseExtension reports whether the vehicle has asked for more base
// (§6.6.3). Honouring this promptly is what keeps a vehicle from braking at
// the decision point.
func NeedsBaseExtension(s *State) bool {
	return BoolOr(s.NewBaseRequest, false)
}
