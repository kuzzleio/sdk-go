// Copyright 2015-2017 Kuzzle
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

package security_test

import (
	"fmt"
	"testing"

	"github.com/kuzzleio/sdk-go/security"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func TestIsActionAllowedNilRights(t *testing.T) {
	res := security.IsActionAllowed(nil, "wow-controller", "such-action", "", "")
	assert.Equal(t, security.ActionIsDenied, res)
}

func TestIsActionAllowedEmptyController(t *testing.T) {
	res := security.IsActionAllowed([]*types.UserRights{}, "", "such-action", "", "")
	assert.Equal(t, security.ActionIsDenied, res)
}

func TestIsActionAllowedEmptyAction(t *testing.T) {
	res := security.IsActionAllowed([]*types.UserRights{}, "wow-controller", "", "", "")
	assert.Equal(t, security.ActionIsDenied, res)
}

func TestIsActionAllowedEmptyRights(t *testing.T) {
	res := security.IsActionAllowed([]*types.UserRights{}, "wow-controller", "such-action", "much-index", "very-collection")
	assert.Equal(t, security.ActionIsDenied, res)
}

func TestIsActionAllowedResultAllowed(t *testing.T) {
	userRights := []*types.UserRights{
		{Controller: "wow-controller", Action: "*", Index: "much-index", Collection: "very-collection", Value: "allowed"},
	}

	res := security.IsActionAllowed(userRights, "wow-controller", "such-action", "much-index", "very-collection")
	assert.Equal(t, security.ActionIsAllowed, res)
}

func TestIsActionAllowedResultConditional(t *testing.T) {
	userRights := []*types.UserRights{
		{Controller: "wow-controller", Action: "*", Index: "much-index", Collection: "very-collection", Value: "conditional"},
	}

	res := security.IsActionAllowed(userRights, "wow-controller", "action", "much-index", "very-collection")

	assert.Equal(t, security.ActionIsConditional, res)
}

func TestIsActionAllowedResultDenied(t *testing.T) {
	userRights := []*types.UserRights{
		{Controller: "wow-controller.", Action: "action-such", Index: "much-index", Collection: "very-collection", Value: "allowed"},
	}

	res := security.IsActionAllowed(userRights, "wow-controller", "action", "much-index", "very-collection")
	assert.Equal(t, security.ActionIsDenied, res)
}

func ExampleIsActionAllowed() {
	userRights := []*types.UserRights{
		{Controller: "wow-controller", Action: "*", Index: "much-index", Collection: "very-collection", Value: "allowed"},
	}

	res := security.IsActionAllowed(userRights, "wow-controller", "such-action", "much-index", "very-collection")
	fmt.Println(res)
}
