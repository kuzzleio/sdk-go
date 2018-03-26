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
