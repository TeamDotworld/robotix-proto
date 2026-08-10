package vda5050

import "github.com/google/uuid"

// Optional-field constructors.
//
// Every optional field in VDA5050 is a pointer, because the protocol treats
// "absent" and "present and zero" differently: an absent `maximumSpeed` means
// no limit, while `maximumSpeed: 0` means the vehicle must not move. These
// helpers keep call sites readable without losing that distinction.

// Float returns a pointer to v.
func Float(v float64) *float64 { return &v }

// Bool returns a pointer to v.
func Bool(v bool) *bool { return &v }

// Str returns a pointer to v.
func Str(v string) *string { return &v }

// Uint returns a pointer to v.
func Uint(v uint32) *uint32 { return &v }

// Int returns a pointer to v.
func Int(v int64) *int64 { return &v }

// FloatOr dereferences p, returning def when p is nil.
func FloatOr(p *float64, def float64) float64 {
	if p == nil {
		return def
	}
	return *p
}

// BoolOr dereferences p, returning def when p is nil.
func BoolOr(p *bool, def bool) bool {
	if p == nil {
		return def
	}
	return *p
}

// StrOr dereferences p, returning def when p is nil.
func StrOr(p *string, def string) string {
	if p == nil {
		return def
	}
	return *p
}

// NewActionID returns a fresh action identifier. §7.3 suggests UUIDs, and
// uniqueness matters: the actionId is the only key linking an action in an
// order to its progress in the vehicle's actionStates array.
func NewActionID() string { return uuid.NewString() }

// NewOrderID returns a fresh order identifier.
func NewOrderID() string { return uuid.NewString() }

// Param builds an action parameter. The value may be any JSON type: string,
// number, bool, array or object (§7.3.1).
func Param(key string, value any) ActionParameter {
	return ActionParameter{Key: key, Value: value}
}

// ParamValue looks up an action parameter by key.
func ParamValue(params []ActionParameter, key string) (any, bool) {
	for _, p := range params {
		if p.Key == key {
			return p.Value, true
		}
	}
	return nil, false
}

// NewAction builds an order action (for a node or an edge).
func NewAction(actionType string, blocking BlockingType, params ...ActionParameter) Action {
	a := Action{
		ActionType:   actionType,
		ActionID:     NewActionID(),
		BlockingType: blocking,
	}
	if len(params) > 0 {
		a.ActionParameters = params
	}
	return a
}

// NewInstantAction builds an instant action. Its blocking type is always NONE
// (§7.4): instant actions run in parallel with whatever the vehicle is doing.
func NewInstantAction(actionType string, params ...ActionParameter) InstantAction {
	a := InstantAction{
		ActionType:   actionType,
		ActionID:     NewActionID(),
		BlockingType: InstantActionBlockingTypeNone,
	}
	if len(params) > 0 {
		a.ActionParameters = params
	}
	return a
}

// CancelOrder builds the cancelOrder instant action. Passing an orderId is
// optional but strongly recommended: without it a vehicle that has already
// moved on to a different order would cancel the wrong one.
func CancelOrder(orderID string) InstantAction {
	if orderID == "" {
		return NewInstantAction(ActionCancelOrder)
	}
	return NewInstantAction(ActionCancelOrder, Param("orderId", orderID))
}

// StartPause builds the startPause instant action.
func StartPause() InstantAction { return NewInstantAction(ActionStartPause) }

// StopPause builds the stopPause instant action.
func StopPause() InstantAction { return NewInstantAction(ActionStopPause) }

// StateRequest asks the vehicle to publish a fresh state message.
func StateRequest() InstantAction { return NewInstantAction(ActionStateRequest) }

// FactsheetRequest asks the vehicle to publish its factsheet. A fleet control
// sends this on first contact to learn the vehicle's capabilities and limits.
func FactsheetRequest() InstantAction { return NewInstantAction(ActionFactsheetRequest) }

// InitializePosition overrides the vehicle's pose, e.g. after a manual move.
func InitializePosition(x, y, theta float64, mapID, lastNodeID string) InstantAction {
	// lastNodeSequenceId is part of the action per the standard and defaults to
	// zero. Sending it explicitly avoids relying on a vehicle to default it.
	return NewInstantAction(ActionInitializePosition,
		Param("x", x), Param("y", y), Param("theta", theta),
		Param("mapId", mapID), Param("lastNodeId", lastNodeID),
		Param("lastNodeSequenceId", 0),
	)
}

// Retry re-runs an action currently sitting in RETRIABLE.
func Retry(actionID string) InstantAction {
	return NewInstantAction(ActionRetry, Param("actionId", actionID))
}

// SkipRetry abandons a RETRIABLE action, moving it to FAILED.
func SkipRetry(actionID string) InstantAction {
	return NewInstantAction(ActionSkipRetry, Param("actionId", actionID))
}

// Trigger releases a waitForTrigger action the vehicle is blocked on.
func Trigger() InstantAction { return NewInstantAction(ActionTrigger) }
