package vda5050

import "testing"

// Predefined action names are wire values fixed by the standard: a vehicle
// matches the exact string and silently ignores anything else, so a wrong name
// produces a fleet manager that publishes happily to a vehicle that never
// responds -- no error on either side.
//
// These are the 3.0.0 spellings. 3.0.0 renamed initPosition to
// initializePosition, so a v2 vehicle will not match and a v2 reference
// implementation is not evidence about what we should send.
func TestPredefinedActionNames(t *testing.T) {
	for name, got := range map[string]string{
		"initializePosition":  ActionInitializePosition,
		"startPause":          ActionStartPause,
		"stopPause":           ActionStopPause,
		"startCharging":       ActionStartCharging,
		"stopCharging":        ActionStopCharging,
		"cancelOrder":         ActionCancelOrder,
		"factsheetRequest":    ActionFactsheetRequest,
		"stateRequest":        ActionStateRequest,
		"waitForTrigger":      ActionWaitForTrigger,
		"trigger":             ActionTrigger,
		"clearInstantActions": ActionClearInstantActions,
	} {
		if got != name {
			t.Errorf("action constant = %q, want %q", got, name)
		}
	}
}

func TestInitializePositionCarriesEveryParameter(t *testing.T) {
	action := InitializePosition(1.5, -2.5, 0.25, "office", "n1")
	if action.ActionType != "initializePosition" {
		t.Fatalf("actionType = %q, want initializePosition (3.0.0 name)", action.ActionType)
	}
	for _, key := range []string{"x", "y", "theta", "mapId", "lastNodeId", "lastNodeSequenceId"} {
		if _, ok := ParamValue(action.ActionParameters, key); !ok {
			t.Errorf("parameter %q missing", key)
		}
	}
}
