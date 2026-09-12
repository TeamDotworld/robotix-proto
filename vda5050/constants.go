package vda5050

// ProtocolVersion is the VDA5050 version this package implements. It is
// written into the `version` header field of every outgoing message.
const ProtocolVersion = "3.0.0"

// MajorVersion is the topic-level version prefix ("v" + major), used as the
// second level of the MQTT topic: interfaceName/majorVersion/manufacturer/...
const MajorVersion = "v3"

// DefaultInterfaceName is the conventional first topic level. The
// specification does not mandate it (cloud brokers often impose their own
// prefix), so it is configurable; only the final `topic` level is mandatory.
const DefaultInterfaceName = "vda5050"

// Topic is the final level of an MQTT topic. These names are mandatory per
// VDA5050 3.0.0 §4.3, regardless of how the preceding levels are structured.
type Topic string

const (
	// TopicOrder carries orders. Published by fleet control. Mandatory.
	TopicOrder Topic = "order"
	// TopicInstantActions carries actions to execute immediately.
	// Published by fleet control. Mandatory.
	TopicInstantActions Topic = "instantActions"
	// TopicState carries the vehicle state. Published by the robot. Mandatory.
	TopicState Topic = "state"
	// TopicVisualization carries high-frequency pose and planned path.
	// Published by the robot. Optional.
	TopicVisualization Topic = "visualization"
	// TopicConnection carries connection state, including the broker-delivered
	// last will. Published by robot and broker. Mandatory.
	TopicConnection Topic = "connection"
	// TopicFactsheet carries vehicle capabilities. Published by the robot.
	// Mandatory.
	TopicFactsheet Topic = "factsheet"
	// TopicZoneSet transfers zone sets to the robot. Published by fleet
	// control. Optional, new in 3.0.
	TopicZoneSet Topic = "zoneSet"
	// TopicResponses answers requests raised in the robot's state.
	// Published by fleet control. Optional, new in 3.0.
	TopicResponses Topic = "responses"
)

// AllTopics lists every topic defined by the specification.
var AllTopics = []Topic{
	TopicOrder, TopicInstantActions, TopicState, TopicVisualization,
	TopicConnection, TopicFactsheet, TopicZoneSet, TopicResponses,
}

// FleetControlTopics are published by the fleet control (master control).
var FleetControlTopics = []Topic{
	TopicOrder, TopicInstantActions, TopicZoneSet, TopicResponses,
}

// RobotTopics are published by the mobile robot (or, for connection, by the
// broker on its behalf). These are the topics a fleet control subscribes to.
var RobotTopics = []Topic{
	TopicState, TopicVisualization, TopicConnection, TopicFactsheet,
}

// QoS returns the MQTT quality of service mandated for a topic by §4.1:
// QoS 0 everywhere except `connection`, which uses QoS 1.
func (t Topic) QoS() byte {
	if t == TopicConnection {
		return 1
	}
	return 0
}

// Retained reports whether messages on this topic must be published with the
// MQTT retained flag. §6.5 requires it for `connection` so that a fleet
// control learns a robot's connectivity as soon as it subscribes, and §6.10
// requires it for `factsheet` -- "all messages on this topic shall be sent
// with a retained flag" -- so that a fleet control joining later learns the
// vehicle's capabilities without having to ask for them.
func (t Topic) Retained() bool {
	return t == TopicConnection || t == TopicFactsheet
}

// Predefined action types (§6.2.3). Vehicles must support cancelOrder,
// startPause and stopPause; the rest are used when the vehicle's factsheet
// advertises them under protocolFeatures.mobileRobotActions.
const (
	ActionStartPause       = "startPause"
	ActionStopPause        = "stopPause"
	ActionStartHibernation = "startHibernation"
	ActionStopHibernation  = "stopHibernation"
	ActionShutdown         = "shutdown"
	ActionStartCharging    = "startCharging"
	ActionStopCharging     = "stopCharging"
	// 3.0.0 renamed this from 2.x's initPosition. We implement 3.0.0 only, so
	// the 3.0.0 spelling is correct here -- a 2.x vehicle would not match it,
	// and that is intentional.
	ActionInitializePosition  = "initializePosition"
	ActionEnableMap           = "enableMap"
	ActionDownloadMap         = "downloadMap"
	ActionDeleteMap           = "deleteMap"
	ActionDownloadZoneSet     = "downloadZoneSet"
	ActionEnableZoneSet       = "enableZoneSet"
	ActionDeleteZoneSet       = "deleteZoneSet"
	ActionClearInstantActions = "clearInstantActions"
	ActionClearZoneActions    = "clearZoneActions"
	ActionStateRequest        = "stateRequest"
	ActionLogReport           = "logReport"
	ActionPick                = "pick"
	ActionDrop                = "drop"
	ActionDetectObject        = "detectObject"
	ActionFinePositioning     = "finePositioning"
	ActionWaitForTrigger      = "waitForTrigger"
	ActionTrigger             = "trigger"
	ActionRetry               = "retry"
	ActionSkipRetry           = "skipRetry"
	ActionCancelOrder         = "cancelOrder"
	ActionFactsheetRequest    = "factsheetRequest"
	ActionUpdateCertificate   = "updateCertificate"
)

// MandatoryActions are the action types every VDA5050 vehicle must support
// (§6.2.3). herdIQ can rely on these without consulting the factsheet.
var MandatoryActions = []string{
	ActionCancelOrder, ActionStartPause, ActionStopPause,
}

// StateIntervalSeconds is the maximum interval at which a compliant vehicle
// publishes its state even when nothing changes (§6.6). A fleet control that
// has not heard from a vehicle in appreciably longer than this should treat
// the vehicle as stale.
const StateIntervalSeconds = 30

// IsTerminal reports whether an action has reached a state from which it will
// not progress on its own. RETRIABLE is deliberately not terminal: it waits
// for a `retry` or `skipRetry` instant action from the fleet control.
func (v ActionStatus) IsTerminal() bool {
	return v == ActionStatusFinished || v == ActionStatusFailed
}

// NeedsIntervention reports whether an action is stalled awaiting a decision
// from the fleet control or an operator.
func (v ActionStatus) NeedsIntervention() bool {
	return v == ActionStatusRetriable
}

// IsBlocking reports whether an error level should take the vehicle out of
// service. 3.0.0 added URGENT and CRITICAL between WARNING and FATAL.
func (v ErrorLevel) IsBlocking() bool {
	return v == ErrorLevelCritical || v == ErrorLevelFatal
}

// Severity maps an error level onto an ordinal so that levels can be compared
// and the worst error in a state message selected.
func (v ErrorLevel) Severity() int {
	switch v {
	case ErrorLevelWarning:
		return 1
	case ErrorLevelUrgent:
		return 2
	case ErrorLevelCritical:
		return 3
	case ErrorLevelFatal:
		return 4
	}
	return 0
}

// IsConnected reports whether the fleet control may send orders to a vehicle
// in this connection state. HIBERNATING vehicles remain attached to the
// broker but accept only the stopHibernation instant action.
func (v ConnectionState) IsConnected() bool {
	return v == ConnectionStateOnline
}

// CanDrive reports whether the vehicle's operating mode puts fleet control in
// charge of motion (Table 11, "Fleet Control in control"). In every other mode
// the vehicle is steered locally, even where it still accepts orders.
func (v OperatingMode) CanDrive() bool {
	return v == OperatingModeAutomatic || v == OperatingModeSemiautomatic
}

// AcceptsOrders reports whether the vehicle will take an order in this mode
// (Table 11, "Sending orders allowed"). INTERVENED is included: an intervened
// vehicle is being steered by hand, but the specification explicitly allows
// fleet control to send orders and order updates to be executed once it
// returns to AUTOMATIC or SEMIAUTOMATIC, and §6.1.4.9 raises
// MOBILE_ROBOT_NOT_AVAILABLE only outside these three modes.
func (v OperatingMode) AcceptsOrders() bool {
	switch v {
	case OperatingModeAutomatic, OperatingModeSemiautomatic, OperatingModeIntervened:
		return true
	}
	return false
}

// AcceptsInstantAction reports whether an instant action of the given type may
// be sent in this mode (Table 11, "Sending instant actions allowed"). In
// INTERVENED only cancelOrder is permitted; in MANUAL, STARTUP, SERVICE and
// TEACH_IN nothing is.
func (v OperatingMode) AcceptsInstantAction(actionType string) bool {
	switch v {
	case OperatingModeAutomatic, OperatingModeSemiautomatic:
		return true
	case OperatingModeIntervened:
		return actionType == ActionCancelOrder
	}
	return false
}

// ClearsOrderOnEntry reports whether entering this mode makes the vehicle
// abandon its current order (Table 11, "Clear order when entering", and
// §6.6.7).
func (v OperatingMode) ClearsOrderOnEntry() bool {
	switch v {
	case OperatingModeManual, OperatingModeStartup, OperatingModeService, OperatingModeTeachIn:
		return true
	}
	return false
}

// ClearsZoneRequestsOnEntry reports whether entering this mode makes the
// vehicle drop every zone request from its state (Table 11, "Clear zone
// requests when entering"). INTERVENED is included even though it keeps the
// order: a hand-steered vehicle must not hold a reservation the fleet control
// is still honouring on its behalf.
func (v OperatingMode) ClearsZoneRequestsOnEntry() bool {
	switch v {
	case OperatingModeIntervened, OperatingModeManual, OperatingModeStartup,
		OperatingModeService, OperatingModeTeachIn:
		return true
	}
	return false
}

// HasValidStateContent reports whether the fields of a state message published
// in this mode can be trusted (Table 11, "Valid state message content"). Only
// STARTUP is exempt: a vehicle that has not finished booting may publish
// incomplete or invalid values.
func (v OperatingMode) HasValidStateContent() bool {
	return v != OperatingModeStartup
}
