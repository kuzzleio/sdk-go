package security_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/kuzzleio/sdk-go/connection/websocket"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/security"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func TestGetRoleEmptyId(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Result: []byte{}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	_, err := k.Security.GetRole("", nil)

	assert.NotNil(t, err)
}

func TestGetRoleError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	_, err := k.Security.GetRole("roleId", nil)

	assert.NotNil(t, err)
}

func TestGetRole(t *testing.T) {
	id := "roleId"

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "getRole", parsedQuery.Action)
			assert.Equal(t, id, parsedQuery.Id)

			res := security.Role{
				Id: id,
				Controllers: map[string]*types.Controller{
					"*": {
						Actions: map[string]bool{
							"*": false,
						},
					},
					"document": {
						Actions: map[string]bool{
							"get":    true,
							"search": true,
						},
					},
				},
			}
			r, _ := (&res).RoleToJson()
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	res, _ := k.Security.GetRole(id, nil)

	assert.Equal(t, id, res.Id)
	assert.Equal(t, map[string]*types.Controller{
		"*": {
			Actions: map[string]bool{"*": false},
		},
		"document": {
			Actions: map[string]bool{
				"get":    true,
				"search": true,
			},
		},
	}, res.Controllers)
}

func ExampleSecurityRole_Fetch() {
	id := "roleId"
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	res, err := k.Security.GetRole(id, nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res.Id, res.Controllers)
}
