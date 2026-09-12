package vda5050

// Predefined error types (§6.6.5.4, Table 9) and the reference keys that
// qualify them (§6.6.5.2).
//
// `errorType` is an extensible enumeration: a vehicle may report a
// manufacturer-specific string that appears nowhere below, and a fleet control
// must not choke on one. What the predefined set buys is the ability to react
// to a named condition rather than to an error level alone -- the difference
// between "this order failed" and "this order needs replanning, that one needs
// the pose re-initialised, and this third one will clear on its own".
//
// The specification is internally inconsistent in one place: the errorType
// enumeration inlined in §7.8 lists 'INVALID_ORDER', while Table 9 and
// §6.1.4.3 both say 'INVALID_ORDER_ACTION'. Table 9 is the normative list of
// predefined error types and is followed here; ErrInvalidOrder is kept as a
// recognised alias so a vehicle that followed the §7.8 spelling is still
// understood.
const (
	// --- order rejection (§6.1.4) --------------------------------------
	// Every one of these is reported until the vehicle accepts a new order.

	// ErrValidationFailure: the order was malformed (§6.1.4.1).
	ErrValidationFailure = "VALIDATION_FAILURE"
	// ErrUnsupportedParameter: the order carried an optional field the
	// vehicle cannot process (§6.1.4.2). The erroneous fields are named in
	// errorReferences.
	ErrUnsupportedParameter = "UNSUPPORTED_PARAMETER"
	// ErrInvalidOrderAction: the order contained actions the vehicle cannot
	// perform (§6.1.4.3).
	ErrInvalidOrderAction = "INVALID_ORDER_ACTION"
	// ErrInvalidOrder is the §7.8 spelling of ErrInvalidOrderAction. Accepted
	// on receipt, never emitted.
	ErrInvalidOrder = "INVALID_ORDER"
	// ErrOutdatedOrderUpdate: same orderId, lower orderUpdateId (§6.1.4.4).
	ErrOutdatedOrderUpdate = "OUTDATED_ORDER_UPDATE"
	// ErrSameOrderUpdateID: same orderId and orderUpdateId but different
	// content (§6.1.4.5). Identical content is ignored silently instead.
	ErrSameOrderUpdateID = "SAME_ORDER_UPDATE_ID"
	// ErrOtherOrderActive: a new orderId arrived while another order is still
	// running (§6.1.4.6).
	ErrOtherOrderActive = "OTHER_ORDER_ACTIVE"
	// ErrStartNodeOutOfRange: the first node of the order is not trivially
	// reachable (§6.1.4.7).
	ErrStartNodeOutOfRange = "START_NODE_OUT_OF_RANGE"
	// ErrNoRouteToTarget: at least one node in the order is unreachable
	// (§6.1.4.8).
	ErrNoRouteToTarget = "NO_ROUTE_TO_TARGET"
	// ErrMobileRobotNotAvailable: an order arrived in an operating mode that
	// does not accept orders (§6.1.4.9). Reported until the mode allows them.
	ErrMobileRobotNotAvailable = "MOBILE_ROBOT_NOT_AVAILABLE"
	// ErrUnknownMapID: the order references a mapId the vehicle does not hold
	// (§6.1.4.10, §6.3.1).
	ErrUnknownMapID = "UNKNOWN_MAP_ID"
	// ErrOrderUpdateFollowingCancel: an update arrived for an order that was
	// already cancelled (§6.1.3.1).
	ErrOrderUpdateFollowingCancel = "ORDER_UPDATE_FOLLOWING_CANCEL"
	// ErrInsufficientMemory: the vehicle cannot hold the incoming order.
	ErrInsufficientMemory = "INSUFFICIENT_MEMORY"

	// --- instant actions -----------------------------------------------

	// ErrNoOrderToCancel: cancelOrder arrived with no matching active order
	// (§6.1.3.2).
	ErrNoOrderToCancel = "NO_ORDER_TO_CANCEL"
	// ErrInvalidInstantAction: an instant action the vehicle cannot execute
	// (§6.2.1). Reported until a new instant action is accepted.
	ErrInvalidInstantAction = "INVALID_INSTANT_ACTION"

	// --- navigation and zones ------------------------------------------

	// ErrOutsideOfCorridor: the vehicle left the corridor defined for an edge
	// (§6.1.5).
	ErrOutsideOfCorridor = "OUTSIDE_OF_CORRIDOR"
	// ErrBlockedZoneViolation: the vehicle entered, or found itself inside, a
	// BLOCKED zone (§6.4.1.1).
	ErrBlockedZoneViolation = "BLOCKED_ZONE_VIOLATION"
	// ErrReleaseLost: the release for a RELEASE zone expired or was revoked
	// (§6.4.1.1).
	ErrReleaseLost = "RELEASE_LOST"
	// ErrZoneActionConflict: a zone's actions conflict with another zone's
	// behaviour (§6.4.4, Table 8 note 1).
	ErrZoneActionConflict = "ZONE_ACTION_CONFLICT"
	// ErrNodeUnreachable: a node in the active order turned out to be
	// unreachable during execution (§6.4.5).
	ErrNodeUnreachable = "NODE_UNREACHABLE"
	// ErrLocalizationError: the vehicle is not localized (§7.8,
	// mobileRobotPosition.localized).
	ErrLocalizationError = "LOCALIZATION_ERROR"

	// --- maps and zone sets --------------------------------------------

	// ErrDuplicateMap: downloadMap for a mapId/mapVersion already held
	// (§6.3.3).
	ErrDuplicateMap = "DUPLICATE_MAP"
	// ErrDuplicateZoneSet: a zone set arrived with a zoneSetId already held
	// (§6.4.2).
	ErrDuplicateZoneSet = "DUPLICATE_ZONE_SET"

	// --- state capacity (§7.8) -----------------------------------------

	// ErrInstantActionStatesFull: the instantActionStates array is growing
	// beyond what the vehicle can manage; fleet control should send
	// clearInstantActions.
	ErrInstantActionStatesFull = "INSTANT_ACTION_STATES_FULL"
	// ErrZoneActionStatesFull: as above for zoneActionStates; fleet control
	// should send clearZoneActions.
	ErrZoneActionStatesFull = "ZONE_ACTION_STATES_FULL"
)

// Reference keys for errorReferences and infoReferences (§6.6.5.2). The
// specification names the keys by example rather than by enumeration, so this
// list is a convention, not a closed set -- but using it consistently is what
// lets a fleet control pull the orderId out of an arbitrary vehicle's error.
const (
	RefHeaderID      = "headerId"
	RefTopic         = "topic"
	RefOrderID       = "orderId"
	RefOrderUpdateID = "orderUpdateId"
	RefActionID      = "actionId"
	RefNodeID        = "nodeId"
	RefEdgeID        = "edgeId"
	RefSequenceID    = "sequenceId"
	RefZoneID        = "zoneId"
	RefZoneSetID     = "zoneSetId"
	RefMapID         = "mapId"
	RefMapVersion    = "mapVersion"
	RefParameter     = "parameter"
	RefRequestID     = "requestId"
)

// predefinedErrorLevels is the level Table 9 assigns to each predefined type.
// A vehicle is free to report a different level and the fleet control must
// honour what it was sent; this table is for the vehicle side, so that an
// implementation raising a predefined error cannot pair it with the wrong
// level, and for diagnostics when a vehicle does deviate.
var predefinedErrorLevels = map[string]ErrorLevel{
	ErrUnsupportedParameter:       ErrorLevelCritical,
	ErrNoOrderToCancel:            ErrorLevelWarning,
	ErrValidationFailure:          ErrorLevelWarning,
	ErrInvalidOrderAction:         ErrorLevelWarning,
	ErrInvalidOrder:               ErrorLevelWarning,
	ErrInvalidInstantAction:       ErrorLevelWarning,
	ErrOutdatedOrderUpdate:        ErrorLevelWarning,
	ErrSameOrderUpdateID:          ErrorLevelWarning,
	ErrOrderUpdateFollowingCancel: ErrorLevelWarning,
	ErrOutsideOfCorridor:          ErrorLevelCritical,
	ErrInsufficientMemory:         ErrorLevelUrgent,
	ErrDuplicateMap:               ErrorLevelWarning,
	ErrBlockedZoneViolation:       ErrorLevelCritical,
	ErrDuplicateZoneSet:           ErrorLevelWarning,
	ErrReleaseLost:                ErrorLevelCritical,
	ErrZoneActionConflict:         ErrorLevelCritical,
	ErrNodeUnreachable:            ErrorLevelCritical,
	ErrLocalizationError:          ErrorLevelFatal,
	ErrNoRouteToTarget:            ErrorLevelWarning,
	ErrOtherOrderActive:           ErrorLevelWarning,
	ErrStartNodeOutOfRange:        ErrorLevelWarning,
	ErrMobileRobotNotAvailable:    ErrorLevelWarning,
	ErrUnknownMapID:               ErrorLevelWarning,
	ErrInstantActionStatesFull:    ErrorLevelUrgent,
	ErrZoneActionStatesFull:       ErrorLevelUrgent,
}

// PredefinedErrorLevel returns the level Table 9 assigns to an error type, and
// whether the type is one of the predefined ones at all.
func PredefinedErrorLevel(errorType string) (ErrorLevel, bool) {
	lvl, ok := predefinedErrorLevels[errorType]
	return lvl, ok
}

// IsPredefinedError reports whether an error type is one the specification
// names. A false result means a manufacturer-specific error, which is legal
// and must still be surfaced to an operator.
func IsPredefinedError(errorType string) bool {
	_, ok := predefinedErrorLevels[errorType]
	return ok
}

// orderRejections are the error types a vehicle raises when it declines an
// order, all from §6.1.4 plus the two cancellation-related ones. Every one is
// reported "until a new order is accepted", so a fleet control that sees one
// should stop waiting for the order it sent and decide what to send instead.
var orderRejections = map[string]struct{}{
	ErrValidationFailure:          {},
	ErrUnsupportedParameter:       {},
	ErrInvalidOrderAction:         {},
	ErrInvalidOrder:               {},
	ErrOutdatedOrderUpdate:        {},
	ErrSameOrderUpdateID:          {},
	ErrOtherOrderActive:           {},
	ErrStartNodeOutOfRange:        {},
	ErrNoRouteToTarget:            {},
	ErrMobileRobotNotAvailable:    {},
	ErrUnknownMapID:               {},
	ErrOrderUpdateFollowingCancel: {},
	ErrInsufficientMemory:         {},
}

// IsOrderRejection reports whether an error type means "I did not take the
// order you sent". These need a fleet-control decision; they will not clear on
// their own, because the vehicle reports them until it accepts a new order.
func IsOrderRejection(errorType string) bool {
	_, ok := orderRejections[errorType]
	return ok
}

// IsRetryable reports whether re-sending a corrected order is a sensible
// response to this rejection. The distinction matters to a dispatcher: a
// START_NODE_OUT_OF_RANGE clears once the vehicle is repositioned or the order
// is rebuilt from its real pose, whereas an INVALID_ORDER_ACTION means the
// vehicle will never accept that action however many times it is sent.
func IsRetryable(errorType string) bool {
	switch errorType {
	case ErrValidationFailure,
		ErrOutdatedOrderUpdate,
		ErrSameOrderUpdateID,
		ErrOtherOrderActive,
		ErrStartNodeOutOfRange,
		ErrNoRouteToTarget,
		ErrMobileRobotNotAvailable,
		ErrUnknownMapID,
		ErrOrderUpdateFollowingCancel,
		ErrInsufficientMemory,
		ErrNodeUnreachable:
		return true
	}
	return false
}

// NewError builds an error object with the level the specification assigns to
// the type. Passing a manufacturer-specific type is allowed; it defaults to
// WARNING, which a caller can override by setting ErrorLevel afterwards.
func NewError(errorType, description string, refs ...ErrorReference) Error {
	level, ok := PredefinedErrorLevel(errorType)
	if !ok {
		level = ErrorLevelWarning
	}
	e := Error{ErrorType: errorType, ErrorLevel: level}
	if description != "" {
		e.ErrorDescription = Str(description)
	}
	if len(refs) > 0 {
		e.ErrorReferences = refs
	}
	return e
}

// Ref builds an error or info reference.
func Ref(key, value string) ErrorReference {
	return ErrorReference{ReferenceKey: key, ReferenceValue: value}
}

// FindError returns the first error of a given type in a state message, or nil.
func FindError(s *State, errorType string) *Error {
	if s == nil {
		return nil
	}
	for i := range s.Errors {
		if s.Errors[i].ErrorType == errorType {
			return &s.Errors[i]
		}
	}
	return nil
}

// RejectionIn returns the first order rejection reported in a state message,
// or nil. A dispatcher polls this after publishing an order: the rejection is
// the vehicle's answer, and it arrives on the state topic rather than as a
// failure of the publish.
func RejectionIn(s *State) *Error {
	if s == nil {
		return nil
	}
	for i := range s.Errors {
		if IsOrderRejection(s.Errors[i].ErrorType) {
			return &s.Errors[i]
		}
	}
	return nil
}

// ErrorReferenceValue looks up one reference on an error.
func ErrorReferenceValue(e *Error, key string) (string, bool) {
	if e == nil {
		return "", false
	}
	for _, r := range e.ErrorReferences {
		if r.ReferenceKey == key {
			return r.ReferenceValue, true
		}
	}
	return "", false
}

// Describe renders an error for a log line or an operator-facing message,
// preferring the vehicle's own description and falling back to the type.
func Describe(e *Error) string {
	if e == nil {
		return ""
	}
	out := string(e.ErrorLevel) + " " + e.ErrorType
	if d := StrOr(e.ErrorDescription, ""); d != "" {
		out += ": " + d
	}
	if h := StrOr(e.ErrorHint, ""); h != "" {
		out += " (" + h + ")"
	}
	return out
}
