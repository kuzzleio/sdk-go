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

func TestRoleSaveEmptyId(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "createRole", parsedQuery.Action)

			return &types.KuzzleResponse{Result: []byte{}}
		},
	}

	k, _ := kuzzle.NewKuzzle(c, nil)
	role := k.Security.NewRole("", nil)

	_, err := role.Save(nil)
	assert.Nil(t, err)
}

func TestRoleSaveEmptyIdWithReplaceOption(t *testing.T) {
	c := &internal.MockedConnection{}

	k, _ := kuzzle.NewKuzzle(c, nil)
	role := k.Security.NewRole("", nil)

	options := types.NewQueryOptions()
	options.SetIfExist("replace")
	_, err := role.Save(options)

	assert.NotNil(t, err)
	assert.Equal(t, 400, err.(*types.KuzzleError).Status)
	assert.Equal(t, "Role.Save: role id is required", err.(*types.KuzzleError).Message)
}

func TestRoleSaveWithErrorOption(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parseQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parseQuery)

			assert.Equal(t, "security", parseQuery.Controller)
			assert.Equal(t, "createRole", parseQuery.Action)
			assert.Equal(t, "roleid", parseQuery.Id)

			return &types.KuzzleResponse{Result: []byte{}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	role := k.Security.NewRole("roleid", nil)
	options := types.NewQueryOptions()
	options.SetIfExist("error")

	_, err := role.Save(options)
	assert.Nil(t, err)
}

func TestRoleSaveInvalidIfExistOption(t *testing.T) {
	c := &internal.MockedConnection{}
	k, _ := kuzzle.NewKuzzle(c, nil)
	role := k.Security.NewRole("roleid", nil)

	options := types.NewQueryOptions()
	options.SetIfExist("invalid")

	_, err := role.Save(options)
	assert.NotNil(t, err)
	assert.Equal(t, 400, err.(*types.KuzzleError).Status)
	assert.Equal(t, "Invalid value for 'ifExist' option: 'invalid'", err.(*types.KuzzleError).Message)
}

func TestRoleSave(t *testing.T) {
	id := "roleId"
	callCount := 0
	newContent := types.Controllers{Controllers: map[string]*types.Controller{"document": {Actions: map[string]bool{"get": true}}}}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			if callCount == 0 {
				callCount++
				assert.Equal(t, "security", parsedQuery.Controller)
				assert.Equal(t, "getRole", parsedQuery.Action)
				assert.Equal(t, id, parsedQuery.Id)

				res := security.Role{
					Id: id,
					Controllers: map[string]*types.Controller{
						"*": {
							Actions: map[string]bool{
								"*": true,
							},
						},
					},
				}
				r, _ := security.RoleToJson(&res)
				return &types.KuzzleResponse{Result: r}
			}
			if callCount == 1 {
				callCount++
				assert.Equal(t, "security", parsedQuery.Controller)
				assert.Equal(t, "createOrReplaceRole", parsedQuery.Action)
				assert.Equal(t, id, parsedQuery.Id)
				assert.Equal(t, map[string]interface{}{
					"controllers": map[string]interface{}{
						"document": map[string]interface{}{
							"actions": map[string]interface{}{"get": true},
						},
					},
				}, parsedQuery.Body)

				res := security.Role{
					Id:          id,
					Controllers: newContent.Controllers,
				}
				r, _ := security.RoleToJson(&res)
				return &types.KuzzleResponse{Result: r}
			}

			return &types.KuzzleResponse{Result: nil}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	r, _ := k.Security.FetchRole(id, nil)

	r.Controllers = newContent.Controllers
	_, err := r.Save(nil)
	assert.Nil(t, err)
}

func ExampleRole_Save() {
	id := "roleId"
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	r, _ := k.Security.FetchRole(id, nil)
	newContent := types.Controllers{Controllers: map[string]*types.Controller{"document": {Actions: map[string]bool{"get": true}}}}

	r.Controllers = newContent.Controllers
	res, err := r.Save(nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res.Id, res.Controllers)
}

func TestRoleDelete(t *testing.T) {
	id := "SomeMenJustWantToWatchTheWorldBurn"
	callCount := 0

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			if callCount == 0 {
				callCount++
				assert.Equal(t, "security", parsedQuery.Controller)
				assert.Equal(t, "getRole", parsedQuery.Action)
				assert.Equal(t, id, parsedQuery.Id)

				res := security.Role{
					Id: id,
					Controllers: map[string]*types.Controller{
						"*": {
							Actions: map[string]bool{
								"*": true,
							},
						},
					},
				}
				r, _ := security.RoleToJson(&res)
				return &types.KuzzleResponse{Result: r}
			}
			if callCount == 1 {
				callCount++
				assert.Equal(t, "security", parsedQuery.Controller)
				assert.Equal(t, "deleteRole", parsedQuery.Action)
				assert.Equal(t, id, parsedQuery.Id)

				res := types.ShardResponse{Id: id}
				r, _ := json.Marshal(res)
				return &types.KuzzleResponse{Result: r}
			}

			return &types.KuzzleResponse{Result: nil}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	r, _ := k.Security.FetchRole(id, nil)

	inTheEnd, _ := r.Delete(nil)
	assert.Equal(t, id, inTheEnd)
}

func ExampleRole_Delete() {
	id := "SomeMenJustWantToWatchTheWorldBurn"
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	r, _ := k.Security.FetchRole(id, nil)

	res, err := r.Delete(nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res)
}
