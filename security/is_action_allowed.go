// Copyright 2015-2018 Kuzzle
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
