package security

import (
	"github.com/kuzzleio/sdk-go/types"
)

const (
	ActionIsAllowed = iota
	ActionIsConditional
	ActionIsDenied
)

// IsActionAllowed indicates whether an action is allowed, denied or conditional based on user rights provided as the first argument.
// An action is defined as a couple of action and controller (mandatory), plus an index and a collection(optional).
func IsActionAllowed(rights []*types.UserRights, controller string, action string, index string, collection string) int {
	if rights == nil {
		return ActionIsDenied
	}
	if controller == "" {
		return ActionIsDenied
	}
	if action == "" {
		return ActionIsDenied
	}

	filteredUserRights := make([]*types.UserRights, 0, len(rights))

	for _, ur := range rights {
		if (ur.Controller == controller || ur.Controller == "*") && (ur.Action == action || ur.Action == "*") && (ur.Index == index || ur.Index == "*") && (ur.Collection == collection || ur.Collection == "*") {
			filteredUserRights = append(filteredUserRights, ur)
		}
	}

	for _, ur := range filteredUserRights {
		if ur.Value == "allowed" {
			return ActionIsAllowed
		}
		if ur.Value == "conditional" {
			return ActionIsConditional
		}
	}

	return ActionIsDenied
}
