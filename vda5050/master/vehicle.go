// Package master implements the VDA5050 fleet-control (master control) role.
//
// It subscribes to the topics vehicles publish (state, connection, factsheet,
// visualization), keeps a live view of every vehicle, and publishes the topics
// fleet control owns (order, instantActions, zoneSet, responses).
//
// The package deliberately knows nothing about herdIQ's traffic control,
// pathfinding or task model. herdIQ drives it from the outside: it decides
// where a vehicle should go, hands over a vda5050.Route, and reads back the
// vehicle's reported state. That boundary keeps protocol conformance testable
// on its own and stops VDA5050 concepts leaking into the fleet logic.
package master

import (
	"sync"
	"time"

	"github.com/GOAT-Robotics/gtstudio-proto/vda5050"
)

// Vehicle is the fleet control's view of one VDA5050 mobile robot: the last
// message received on each topic, plus the order lifecycle state.
//
// A Vehicle is safe for concurrent use. Accessors return copies or immutable
// snapshots so that callers cannot mutate the stored message.
type Vehicle struct {
	// ID is the manufacturer/serialNumber pair addressing this vehicle.
	ID vda5050.Identity

	// Orders tracks orderIds, orderUpdateIds and the base/horizon split.
	Orders *vda5050.OrderTracker

	mu sync.RWMutex

	state   *vda5050.State
	stateAt time.Time

	conn   vda5050.ConnectionState
	connAt time.Time

	factsheet   *vda5050.Factsheet
	factsheetAt time.Time

	vis   *vda5050.Visualization
	visAt time.Time

	// firstSeen is when this vehicle was first observed on the bus.
	firstSeen time.Time
}

func newVehicle(id vda5050.Identity, headers *vda5050.HeaderCounter) *Vehicle {
	return &Vehicle{
		ID:        id,
		Orders:    vda5050.NewOrderTracker(id, headers),
		firstSeen: time.Now(),
		// A vehicle that has not published a connection message yet is
		// treated as offline rather than online: assuming the optimistic case
		// would let herdIQ dispatch to a vehicle that is not listening.
		conn: vda5050.ConnectionStateOffline,
	}
}

// State returns the most recent state message, or nil if none has arrived.
func (v *Vehicle) State() *vda5050.State {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.state
}

// StateAge reports how long ago the last state message arrived. Vehicles must
// publish at least every 30 seconds (§6.6), so an age well beyond that means
// the vehicle has stopped reporting even if MQTT still shows it connected.
func (v *Vehicle) StateAge() time.Duration {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.stateAt.IsZero() {
		return time.Since(v.firstSeen)
	}
	return time.Since(v.stateAt)
}

// ConnectionState returns the last reported MQTT-level connection state.
func (v *Vehicle) ConnectionState() vda5050.ConnectionState {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.conn
}

// Factsheet returns the vehicle's factsheet, or nil if it has not published
// one. The factsheet is retained on the broker, so it normally arrives on
// subscribe; if it does not, send a factsheetRequest instant action.
func (v *Vehicle) Factsheet() *vda5050.Factsheet {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.factsheet
}

// Visualization returns the last high-frequency pose update, or nil.
func (v *Vehicle) Visualization() *vda5050.Visualization {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.vis
}

// Online reports whether the vehicle is connected and reporting recently
// enough to be dispatched to. staleAfter bounds the acceptable state age;
// pass 0 for the specification's 30-second interval with a safety margin.
func (v *Vehicle) Online(staleAfter time.Duration) bool {
	if staleAfter <= 0 {
		staleAfter = 3 * vda5050.StateIntervalSeconds * time.Second
	}
	v.mu.RLock()
	conn, state, at := v.conn, v.state, v.stateAt
	v.mu.RUnlock()

	if !conn.IsConnected() || state == nil || at.IsZero() {
		return false
	}
	return time.Since(at) <= staleAfter
}

// Dispatchable reports whether the fleet control may send this vehicle an
// order right now: connected, reporting, in an automatic operating mode, not
// paused, not emergency-stopped and free of blocking errors.
func (v *Vehicle) Dispatchable(staleAfter time.Duration) (bool, string) {
	if !v.Online(staleAfter) {
		return false, "vehicle is offline or its state is stale"
	}
	s := v.State()
	if !s.OperatingMode.CanDrive() {
		return false, "vehicle is in operating mode " + string(s.OperatingMode)
	}
	if vda5050.BoolOr(s.Paused, false) {
		return false, "vehicle is paused"
	}
	// Block only on a *known* active emergency stop. safetyState is a required
	// field, so an empty value means the vehicle sent a non-conformant message
	// rather than that its e-stop is engaged — and grounding a fleet because a
	// vendor omitted a field, with no explanation of why, is worse than
	// letting incoming schema validation report the real problem. A genuinely
	// e-stopped vehicle also leaves AUTOMATIC mode, which is checked above.
	switch s.SafetyState.ActiveEmergencyStop {
	case vda5050.ActiveEmergencyStopManual, vda5050.ActiveEmergencyStopRemote:
		return false, "vehicle emergency stop is active (" + string(s.SafetyState.ActiveEmergencyStop) + ")"
	}
	if s.SafetyState.FieldViolation {
		return false, "vehicle reports a protective field violation"
	}
	for _, e := range s.Errors {
		if e.ErrorLevel.IsBlocking() {
			return false, "vehicle reports a " + string(e.ErrorLevel) + " error: " + e.ErrorType
		}
	}
	return true, ""
}

// BatteryCharge returns the state of charge in percent, and whether the
// vehicle has reported a state at all.
func (v *Vehicle) BatteryCharge() (float64, bool) {
	s := v.State()
	if s == nil {
		return 0, false
	}
	return s.PowerSupply.StateOfCharge, true
}

// Charging reports whether the vehicle is currently charging.
func (v *Vehicle) Charging() bool {
	s := v.State()
	return s != nil && s.PowerSupply.Charging
}

// Position returns the vehicle's pose. ok is false when the vehicle does not
// publish a position (line-guided vehicles often do not) or when its
// localisation is not yet initialised — publishing a stale pose as if it were
// current is worse than reporting none.
func (v *Vehicle) Position() (x, y, theta float64, mapID string, ok bool) {
	s := v.State()
	if s == nil || s.MobileRobotPosition == nil {
		return 0, 0, 0, "", false
	}
	p := s.MobileRobotPosition
	if !p.Localized {
		return 0, 0, 0, "", false
	}
	return p.X, p.Y, p.Theta, p.MapID, true
}

// Driving reports whether the vehicle is in motion.
func (v *Vehicle) Driving() bool {
	s := v.State()
	return s != nil && s.Driving
}

// WorstError returns the highest-severity error the vehicle is reporting.
func (v *Vehicle) WorstError() *vda5050.Error {
	s := v.State()
	if s == nil {
		return nil
	}
	var worst *vda5050.Error
	for i := range s.Errors {
		if worst == nil || s.Errors[i].ErrorLevel.Severity() > worst.ErrorLevel.Severity() {
			worst = &s.Errors[i]
		}
	}
	return worst
}

// ActionsNeedingIntervention returns actions sitting in RETRIABLE, which wait
// for a retry or skipRetry instant action from the fleet control and will
// otherwise stall the order indefinitely.
func (v *Vehicle) ActionsNeedingIntervention() []vda5050.ActionState {
	s := v.State()
	if s == nil {
		return nil
	}
	var out []vda5050.ActionState
	for _, a := range s.ActionStates {
		if a.ActionStatus.NeedsIntervention() {
			out = append(out, a)
		}
	}
	for _, a := range s.InstantActionStates {
		if a.ActionStatus.NeedsIntervention() {
			out = append(out, a)
		}
	}
	return out
}

// SupportsAction reports whether the factsheet advertises an action type. It
// returns true for the three mandatory actions even without a factsheet, and
// false for everything else until one arrives — an unknown capability is
// safer treated as absent.
func (v *Vehicle) SupportsAction(actionType string) bool {
	for _, m := range vda5050.MandatoryActions {
		if m == actionType {
			return true
		}
	}
	fs := v.Factsheet()
	if fs == nil {
		return false
	}
	for _, a := range fs.ProtocolFeatures.MobileRobotActions {
		if a.ActionType == actionType {
			return true
		}
	}
	return false
}

// MinStateInterval returns the shortest interval between state messages the
// vehicle guarantees, taken from its factsheet. It bounds how quickly fleet
// control can expect to observe the effect of an order.
func (v *Vehicle) MinStateInterval() (time.Duration, bool) {
	fs := v.Factsheet()
	if fs == nil || fs.ProtocolLimits.Timing.MinimumStateInterval == 0 {
		return 0, false
	}
	return time.Duration(fs.ProtocolLimits.Timing.MinimumStateInterval * float64(time.Second)), true
}

// MinOrderInterval returns the shortest interval the vehicle accepts between
// two order messages. Publishing order updates faster than this risks the
// vehicle rejecting them, so herdIQ's base-extension loop should rate-limit
// itself to this value.
func (v *Vehicle) MinOrderInterval() (time.Duration, bool) {
	fs := v.Factsheet()
	if fs == nil || fs.ProtocolLimits.Timing.MinimumOrderInterval == 0 {
		return 0, false
	}
	return time.Duration(fs.ProtocolLimits.Timing.MinimumOrderInterval * float64(time.Second)), true
}
