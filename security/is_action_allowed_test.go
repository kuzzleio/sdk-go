package security_test

import (
	"fmt"
	"testing"

	"github.com/kuzzleio/sdk-go/connection/websocket"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/security"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func TestIsActionAllowedNilRights(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)

	_, err := k.Security.IsActionAllowed(nil, "wow-controller", "such-action", "", "")
	assert.NotNil(t, err)
}

func TestIsActionAllowedEmptyController(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)

	_, err := k.Security.IsActionAllowed([]*types.UserRights{}, "", "such-action", "", "")
	assert.NotNil(t, err)
}

func TestIsActionAllowedEmptyAction(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)

	_, err := k.Security.IsActionAllowed([]*types.UserRights{}, "wow-controller", "", "", "")
	assert.NotNil(t, err)
}

func TestIsActionAllowedEmptyRights(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)

	res, _ := k.Security.IsActionAllowed([]*types.UserRights{}, "wow-controller", "such-action", "much-index", "very-collection")

	assert.Equal(t, security.ActionIsDenied, res)
}

func TestIsActionAllowedResultAllowed(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)

	userRights := []*types.UserRights{
		{Controller: "wow-controller", Action: "*", Index: "much-index", Collection: "very-collection", Value: "allowed"},
	}

	res, _ := k.Security.IsActionAllowed(userRights, "wow-controller", "such-action", "much-index", "very-collection")

	assert.Equal(t, security.ActionIsAllowed, res)
}

func TestIsActionAllowedResultConditional(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)

	userRights := []*types.UserRights{
		{Controller: "wow-controller", Action: "*", Index: "much-index", Collection: "very-collection", Value: "conditional"},
	}

	res, _ := k.Security.IsActionAllowed(userRights, "wow-controller", "action", "much-index", "very-collection")

	assert.Equal(t, security.ActionIsConditional, res)
}

func TestIsActionAllowedResultDenied(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)

	userRights := []*types.UserRights{
		{Controller: "wow-controller.", Action: "action-such", Index: "much-index", Collection: "very-collection", Value: "allowed"},
	}

	res, _ := k.Security.IsActionAllowed(userRights, "wow-controller", "action", "much-index", "very-collection")

	assert.Equal(t, security.ActionIsDenied, res)
}

func ExampleSecurityUser_IsActionAllowed() {
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)

	userRights := []*types.UserRights{
		{Controller: "wow-controller", Action: "*", Index: "much-index", Collection: "very-collection", Value: "allowed"},
	}

	res, err := k.Security.IsActionAllowed(userRights, "wow-controller", "such-action", "much-index", "very-collection")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res)
}
