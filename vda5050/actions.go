package vda5050

// Builders for the predefined actions of §6.2.3 (Table 4).
//
// Every one of these could be written by hand with NewInstantAction and
// Param, and for a manufacturer-specific action that is exactly what a caller
// should do. The value of having them typed is that the parameter *names* are
// wire values fixed by the standard in the same way the action names are: a
// vehicle looks up "mapId" and ignores anything it does not recognise, so a
// misspelled parameter produces an action that publishes cleanly, is accepted
// cleanly, and does nothing -- with no error on either side.
//
// Table 4 also fixes each action's scope (instant / node / edge / zone). The
// scope is encoded here by which constructor exists: actions valid only as
// instant actions return InstantAction, actions valid on a node or edge return
// Action, and the two that are valid in both scopes have a constructor each.

// ---------------------------------------------------------------------------
// Pause and hibernation
// ---------------------------------------------------------------------------

// StartHibernation puts the vehicle into hibernate mode: it stays attached to
// the broker, stops publishing state and reports connectionState HIBERNATING.
// An active order is cleared. wakeUpTime is optional -- pass "" for a vehicle
// that must be woken explicitly with StopHibernation; otherwise pass a
// VDA5050 timestamp, and the vehicle leaves hibernation on its own at that
// time.
//
// Note that after this action the vehicle answers nothing but stopHibernation,
// so a fleet control must keep an MQTT client subscribed to instantActions for
// it, and must not treat the absence of state messages as a fault.
func StartHibernation(wakeUpTime string) InstantAction {
	if wakeUpTime == "" {
		return NewInstantAction(ActionStartHibernation)
	}
	return NewInstantAction(ActionStartHibernation, Param("wakeUpTime", wakeUpTime))
}

// StopHibernation wakes a hibernating vehicle. On success it publishes
// connectionState ONLINE.
func StopHibernation() InstantAction { return NewInstantAction(ActionStopHibernation) }

// Shutdown disconnects the vehicle from the broker in a coordinated way. The
// vehicle must be idle. There is no VDA5050 way to restart it afterwards, and
// a hibernating vehicle must be woken with StopHibernation first.
func Shutdown() InstantAction { return NewInstantAction(ActionShutdown) }

// ---------------------------------------------------------------------------
// Charging
// ---------------------------------------------------------------------------

// StartCharging begins charging. Valid as an instant action and on a node.
func StartCharging() InstantAction { return NewInstantAction(ActionStartCharging) }

// StopCharging ends charging. Valid as an instant action and on a node.
func StopCharging() InstantAction { return NewInstantAction(ActionStopCharging) }

// StartChargingAt is the node-scoped form of StartCharging.
func StartChargingAt(blocking BlockingType) Action {
	return NewAction(ActionStartCharging, blocking)
}

// StopChargingAt is the node-scoped form of StopCharging.
func StopChargingAt(blocking BlockingType) Action {
	return NewAction(ActionStopCharging, blocking)
}

// ---------------------------------------------------------------------------
// Maps (§6.3)
// ---------------------------------------------------------------------------

// DownloadMap tells the vehicle to pull a map from the map server. The
// download is a vehicle-initiated pull: fleet control only supplies the link.
// mapHash is optional; pass "" to omit it.
//
// The action reaches FINISHED only once the vehicle has verified the download
// and added the map to its `maps` array with mapStatus DISABLED. Enabling it
// is a separate step -- see EnableMap.
func DownloadMap(mapID, mapVersion, mapDownloadLink, mapHash string) InstantAction {
	params := []ActionParameter{
		Param("mapId", mapID),
		Param("mapVersion", mapVersion),
		Param("mapDownloadLink", mapDownloadLink),
	}
	if mapHash != "" {
		params = append(params, Param("mapHash", mapHash))
	}
	return NewInstantAction(ActionDownloadMap, params...)
}

// EnableMap activates a map the vehicle already holds. Every other version of
// the same mapId is set to DISABLED as a result, so this is how a fleet
// control switches a vehicle onto a new map version without re-localising it.
func EnableMap(mapID, mapVersion string) InstantAction {
	return NewInstantAction(ActionEnableMap,
		Param("mapId", mapID), Param("mapVersion", mapVersion))
}

// EnableMapAt is the node-scoped form of EnableMap, used at an elevator or
// floor-transition node where the map changes mid-order.
func EnableMapAt(mapID, mapVersion string, blocking BlockingType) Action {
	return NewAction(ActionEnableMap, blocking,
		Param("mapId", mapID), Param("mapVersion", mapVersion))
}

// DeleteMap removes a map from the vehicle's memory. A map currently in use
// cannot be deleted; the vehicle reports the action FAILED instead.
func DeleteMap(mapID, mapVersion string) InstantAction {
	return NewInstantAction(ActionDeleteMap,
		Param("mapId", mapID), Param("mapVersion", mapVersion))
}

// ---------------------------------------------------------------------------
// Zone sets (§6.4.2)
// ---------------------------------------------------------------------------

// DownloadZoneSet tells the vehicle to pull a zone set from a server. This is
// the large-zone-set alternative to publishing on the zoneSet topic, and it
// follows the same download/enable/delete shape as maps. zoneSetHash is
// optional; pass "" to omit it.
func DownloadZoneSet(zoneSetID, zoneSetDownloadLink, zoneSetHash string) InstantAction {
	params := []ActionParameter{
		Param("zoneSetId", zoneSetID),
		Param("zoneSetDownloadLink", zoneSetDownloadLink),
	}
	if zoneSetHash != "" {
		params = append(params, Param("zoneSetHash", zoneSetHash))
	}
	return NewInstantAction(ActionDownloadZoneSet, params...)
}

// EnableZoneSet activates a zone set the vehicle already holds. Every other
// zone set for the same map is disabled as a result: at most one zone set per
// map may be ENABLED at a time.
func EnableZoneSet(zoneSetID string) InstantAction {
	return NewInstantAction(ActionEnableZoneSet, Param("zoneSetId", zoneSetID))
}

// EnableZoneSetAt is the node-scoped form of EnableZoneSet.
func EnableZoneSetAt(zoneSetID string, blocking BlockingType) Action {
	return NewAction(ActionEnableZoneSet, blocking, Param("zoneSetId", zoneSetID))
}

// DeleteZoneSet removes a zone set from the vehicle's memory.
func DeleteZoneSet(zoneSetID string) InstantAction {
	return NewInstantAction(ActionDeleteZoneSet, Param("zoneSetId", zoneSetID))
}

// ---------------------------------------------------------------------------
// State housekeeping
// ---------------------------------------------------------------------------

// ClearInstantActions drops every FINISHED or FAILED entry from the vehicle's
// instantActionStates. The vehicle keeps instant action states until told to
// clear them, so a fleet control that never sends this will eventually see the
// vehicle raise INSTANT_ACTION_STATES_FULL.
func ClearInstantActions() InstantAction { return NewInstantAction(ActionClearInstantActions) }

// ClearZoneActions drops every FINISHED or FAILED entry from the vehicle's
// zoneActionStates, with the same rationale as ClearInstantActions.
func ClearZoneActions() InstantAction { return NewInstantAction(ActionClearZoneActions) }

// LogReport asks the vehicle to generate and store a diagnostic log. The
// reason is free text and is echoed back in the action's actionResult
// together with the name of the stored log.
func LogReport(reason string) InstantAction {
	return NewInstantAction(ActionLogReport, Param("reason", reason))
}

// UpdateCertificate hands the vehicle a new certificate set for a named
// service. `service` is an extensible enum whose one predefined value is
// "MQTT". certificateAuthorityDownloadLink is optional; pass "" to omit it.
//
// §6.2.3.3 requires these downloads to be TLS-secured, because the sender of
// an instant action cannot itself be authenticated.
func UpdateCertificate(service, keyDownloadLink, certificateDownloadLink, caDownloadLink string) InstantAction {
	params := []ActionParameter{
		Param("service", service),
		Param("keyDownloadLink", keyDownloadLink),
		Param("certificateDownloadLink", certificateDownloadLink),
	}
	if caDownloadLink != "" {
		params = append(params, Param("certificateAuthorityDownloadLink", caDownloadLink))
	}
	return NewInstantAction(ActionUpdateCertificate, params...)
}

// ---------------------------------------------------------------------------
// Load handling and positioning (node/edge scope)
// ---------------------------------------------------------------------------

// LoadParams carries the optional parameters shared by pick and drop. Leave a
// field at its zero value to omit it: the specification distinguishes an
// absent height from a height of zero.
type LoadParams struct {
	// LHD names the load handling device, for a vehicle with more than one.
	LHD string
	// StationType describes how the operation is handled, e.g. "floor",
	// "rack", "conveyor".
	StationType string
	// StationName identifies the specific station.
	StationName string
	// LoadType is the load unit, e.g. "EPAL", "INDU".
	LoadType string
	// LoadID identifies the individual load, e.g. a barcode.
	LoadID string
	// Height is the bottom of the load relative to the floor, in metres.
	Height *float64
	// Depth is the depth for forklift-side operations, in metres.
	Depth *float64
	// Side is the approach side, e.g. "conveyor".
	Side string
}

func (p LoadParams) params() []ActionParameter {
	var out []ActionParameter
	if p.LHD != "" {
		out = append(out, Param("lhd", p.LHD))
	}
	if p.StationType != "" {
		out = append(out, Param("stationType", p.StationType))
	}
	if p.StationName != "" {
		out = append(out, Param("stationName", p.StationName))
	}
	if p.LoadType != "" {
		out = append(out, Param("loadType", p.LoadType))
	}
	if p.LoadID != "" {
		out = append(out, Param("loadId", p.LoadID))
	}
	if p.Height != nil {
		out = append(out, Param("height", *p.Height))
	}
	if p.Depth != nil {
		out = append(out, Param("depth", *p.Depth))
	}
	if p.Side != "" {
		out = append(out, Param("side", p.Side))
	}
	return out
}

// Pick builds a pick action for a node or an edge. It is not idempotent, so a
// fleet control must not re-send it blindly after a timeout: check the
// actionState first.
func Pick(blocking BlockingType, p LoadParams) Action {
	return NewAction(ActionPick, blocking, p.params()...)
}

// Drop builds a drop action for a node or an edge. Like Pick, it is not
// idempotent.
func Drop(blocking BlockingType, p LoadParams) Action {
	return NewAction(ActionDrop, blocking, p.params()...)
}

// DetectObject asks the vehicle to detect an object of a given type -- a load,
// a charging spot, a free parking position. objectType is optional; pass "".
func DetectObject(blocking BlockingType, objectType string) Action {
	if objectType == "" {
		return NewAction(ActionDetectObject, blocking)
	}
	return NewAction(ActionDetectObject, blocking, Param("objectType", objectType))
}

// FinePositioning asks the vehicle to position itself precisely on a target,
// overriding the node's allowed deviation. Both parameters are optional.
func FinePositioning(blocking BlockingType, stationType, stationName string) Action {
	var params []ActionParameter
	if stationType != "" {
		params = append(params, Param("stationType", stationType))
	}
	if stationName != "" {
		params = append(params, Param("stationName", stationName))
	}
	return NewAction(ActionFinePositioning, blocking, params...)
}

// Trigger types with a defined meaning (§6.2.3, waitForTrigger). The parameter
// is an array of strings and is extensible; these two are to be used whenever
// they fit, so that a fleet control can tell "I am waiting for you" from "I am
// waiting for a button on my own panel".
const (
	// TriggerFleetControl means the trigger will come from fleet control, as
	// the `trigger` instant action.
	TriggerFleetControl = "FLEET_CONTROL"
	// TriggerLocal means the trigger comes from an input on the vehicle -- a
	// button press, a manual load.
	TriggerLocal = "LOCAL"
)

// WaitForTrigger parks the vehicle at a node until one of the named triggers
// fires. Fleet control owns the timeout: if nothing fires, it must cancel the
// order.
func WaitForTrigger(blocking BlockingType, triggerTypes ...string) Action {
	if len(triggerTypes) == 0 {
		triggerTypes = []string{TriggerFleetControl}
	}
	// The parameter is an array of strings, not one string per parameter.
	types := make([]any, 0, len(triggerTypes))
	for _, t := range triggerTypes {
		types = append(types, t)
	}
	return NewAction(ActionWaitForTrigger, blocking, Param("triggerType", types))
}
