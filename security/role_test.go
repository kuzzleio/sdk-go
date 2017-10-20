package security_test

import (
	"testing"
	"encoding/json"
	"fmt"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
	"github.com/kuzzleio/sdk-go/connection/websocket"
	"github.com/kuzzleio/sdk-go/security"
)

func TestRoleSetContent(t *testing.T) {
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
				Controllers: map[string]*types.Controller {
					"*": {
						Actions: map[string]bool {
							"*": true,
						},
					},
				},
			}
			r, _ := security.RoleToJson(&res)
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	r, _ := k.Security.FetchRole(id, nil)

	newContent := types.Controllers{Controllers: map[string]*types.Controller{"document": {Actions: map[string]bool{"get": true}}}}

	r.Controllers = newContent.Controllers
	assert.Equal(t, newContent.Controllers, r.Controllers)
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
					Controllers: map[string]*types.Controller {
						"*": {
							Actions: map[string]bool {
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

				res := security.Role{
					Id: id,
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
	r.Save(nil)

	assert.Equal(t, newContent.Controllers, r.Controllers)
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

func TestRoleUpdate(t *testing.T) {
	id := "roleId"
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
					Controllers: map[string]*types.Controller {
						"*": {
							Actions: map[string]bool {
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
				assert.Equal(t, "updateRole", parsedQuery.Action)
				assert.Equal(t, id, parsedQuery.Id)

				res := security.Role{
					Id: id,
					Controllers: map[string]*types.Controller {
						"document": {
							Actions: map[string]bool {
								"get": true,
							},
						},
					},
				}
				r, _ := security.RoleToJson(&res)
				return &types.KuzzleResponse{Result: r}
			}

			return &types.KuzzleResponse{Result: nil}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	r, _ := k.Security.FetchRole(id, nil)

	newContent := types.Controllers{Controllers: map[string]*types.Controller{"document": {Actions: map[string]bool{"get": true}}}}

	updatedRole, _ := r.Update(&newContent, nil)

	assert.Equal(t, newContent.Controllers, updatedRole.Controllers)
}

func ExampleRole_Update() {
	id := "roleId"
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	r, _ := k.Security.FetchRole(id, nil)

	newContent := types.Controllers{Controllers: map[string]*types.Controller{"document": {Actions: map[string]bool{"get": true}}}}

	res, err := r.Update(&newContent, nil)

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
					Controllers: map[string]*types.Controller {
						"*": {
							Actions: map[string]bool {
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

func TestFetchRoleEmptyId(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.MessageError{Message: "Security.Role.Fetch: role id required"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	_, err := k.Security.FetchRole("", nil)

	assert.NotNil(t, err)
}

func TestFetchRoleError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	_, err := k.Security.FetchRole("roleId", nil)

	assert.NotNil(t, err)
}

func TestFetchRole(t *testing.T) {
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
				Controllers: map[string]*types.Controller {
					"*": {
						Actions: map[string]bool {
							"*": true,
						},
					},
				},
			}
			r, _ := security.RoleToJson(&res)
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	res, _ := k.Security.FetchRole(id, nil)

	assert.Equal(t, id, res.Id)
	assert.Equal(t, true, res.Controllers["*"].Actions["*"])
}

func ExampleSecurityRole_Fetch() {
	id := "roleId"
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	res, err := k.Security.FetchRole(id, nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res.Id, res.Controllers)
}

func TestSearchRolesError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	_, err := k.Security.FetchRole("", nil)

	assert.NotNil(t, err)
}

func TestSearchRoles(t *testing.T) {
	jsonResult := []byte(`{"total":42,"hits":[{"_id":"role42","_source":{"controllers":{"*":{"actions":{"*":true}}}}}]}`)

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "searchRoles", parsedQuery.Action)

			return &types.KuzzleResponse{Result: jsonResult}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	res, _ := k.Security.SearchRoles(nil, nil)

	assert.Equal(t, 42, res.Total)
	assert.Equal(t, res.Hits[0].Controllers["*"].Actions["*"], true)
}

func ExampleSecurityRole_Search() {
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	res, err := k.Security.SearchRoles(nil, nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res.Hits[0].Id, res.Hits[0].Controllers)
}

func TestSearchWithOptions(t *testing.T) {
	jsonResult := []byte(`{
		"total": 42,
		"hits": [
			{
				"_id": "role42",
				"_source": {
					"controllers": {
						"*": {
							"actions": {
								"*": true
							}
						}
					}
				}
			}
		]
	}`)

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "searchRoles", parsedQuery.Action)

			return &types.KuzzleResponse{Result: jsonResult}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	opts := types.NewQueryOptions()
	opts.SetFrom(2)
	opts.SetSize(4)
	opts.SetScroll("1m")

	res, _ := k.Security.SearchRoles(nil, opts)
	assert.Equal(t, 42, res.Total)
	assert.Equal(t, res.Hits[0].Id, "role42")
	assert.Equal(t, res.Hits[0].Controllers["*"].Actions["*"], true)
}

/*
func TestCreateEmptyId(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.MessageError{Message: "Security.Role.Create: role id required"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	r, err := role.Fetch(k.Security, id, nil)
	_, err := security.New(k).Role.Create("", &types.Controllers{}, nil)
	assert.NotNil(t, err)
}
*/

/*
func TestCreateError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := security.New(k).Role.Create("roleId", &types.Controllers{}, nil)
	assert.NotNil(t, err)
}
*/

/*
func TestCreate(t *testing.T) {
	id := "roleId"

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "createRole", parsedQuery.Action)
			assert.Equal(t, id, parsedQuery.Id)

			res := role.Role{
				Id: id,
				Controllers: map[string]types.Controller {
					"*": {
						Actions: map[string]bool {
							"*": true,
						},
					},
				},
			}
			r, _ := res.ToJson()
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := security.New(k).Role.Create(id, &types.Controllers{Controllers: map[string]types.Controller{"*": {Actions: map[string]bool{"*": true}}}}, nil)

	assert.Equal(t, id, res.Id)
	assert.Equal(t, true, res.Controllers()["*"].Actions["*"])
}
*/

/*
func ExampleSecurityRole_Create() {
	id := "roleId"
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, err := security.New(k).Role.Create(id, &types.Controllers{Controllers: map[string]types.Controller{"*": {Actions: map[string]bool{"*": true}}}}, nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res.Id, res.Controllers())
}
*/

/*
func TestCreateIfExists(t *testing.T) {
	id := "roleId"

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "createOrReplaceRole", parsedQuery.Action)
			assert.Equal(t, id, parsedQuery.Id)

			res := role.Role{
				Id:     id,
				Source: []byte(`{"controllers":{"*":{"actions":{"*":true}}}}`),
			}
			r, _ := json.Marshal(res)
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	opts := types.NewQueryOptions()
	opts.SetIfExist("replace")

	res, _ := security.New(k).Role.Create(id, &types.Controllers{Controllers: map[string]types.Controller{"*": {Actions: map[string]bool{"*": true}}}}, opts)

	assert.Equal(t, id, res.Id)
	assert.Equal(t, true, res.Controllers()["*"].Actions["*"])
}
*/

/*
func TestCreateWithStrictOption(t *testing.T) {
	id := "roleId"

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "createRole", parsedQuery.Action)
			assert.Equal(t, id, parsedQuery.Id)

			res := role.Role{
				Id:     id,
				Source: []byte(`{"controllers":{"*":{"actions":{"*":true}}}}`),
			}
			r, _ := json.Marshal(res)
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	opts := types.NewQueryOptions()
	opts.SetIfExist("error")

	res, _ := security.New(k).Role.Create(id, &types.Controllers{Controllers: map[string]types.Controller{"*": {Actions: map[string]bool{"*": true}}}}, opts)

	assert.Equal(t, id, res.Id)
	assert.Equal(t, true, res.Controllers()["*"].Actions["*"])
}
*/

/*
func TestCreateWithWrongOption(t *testing.T) {
	id := "roleId"

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	opts := types.NewQueryOptions()
	opts.SetIfExist("unknown")

	_, err := security.New(k).Role.Create(id, &types.Controllers{Controllers: map[string]types.Controller{"*": {Actions: map[string]bool{"*": true}}}}, opts)

	assert.Equal(t, "Invalid value for the 'ifExist' option: 'unknown'", fmt.Sprint(err))
}
*/

/*
func TestUpdateEmptyId(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.MessageError{Message: "Security.Role.Update: role id required"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := security.New(k).Role.Update("", &types.Controllers{}, nil)
	assert.NotNil(t, err)
}
*/

/*
func TestUpdateError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := security.New(k).Role.Update("roleId", &types.Controllers{}, nil)
	assert.NotNil(t, err)
}
*/

/*
func TestUpdate(t *testing.T) {
	id := "roleId"

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "updateRole", parsedQuery.Action)
			assert.Equal(t, id, parsedQuery.Id)

			res := role.Role{
				Id:     id,
				Source: []byte(`{"controllers":{"*":{"actions":{"*":true}}}}`),
			}
			r, _ := json.Marshal(res)
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := security.New(k).Role.Update(id, &types.Controllers{Controllers: map[string]types.Controller{"*": {Actions: map[string]bool{"*": true}}}}, nil)

	assert.Equal(t, id, res.Id)
	assert.Equal(t, true, res.Controllers()["*"].Actions["*"])
}
*/


/*
func ExampleSecurityRole_Update() {
	id := "roleId"
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, err := security.New(k).Role.Update(id, &types.Controllers{Controllers: map[string]types.Controller{"*": {Actions: map[string]bool{"*": true}}}}, nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res.Id, res.Controllers())
}
*/

/*
func TestDeleteEmptyId(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.MessageError{Message: "Security.Role.Delete: role id required"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := security.New(k).Role.Delete("", nil)
	assert.NotNil(t, err)
}
*/

/*
func TestDeleteError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := security.New(k).Role.Delete("roleId", nil)
	assert.NotNil(t, err)
}
*/

/*
func TestDelete(t *testing.T) {
	id := "roleId"

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "deleteRole", parsedQuery.Action)
			assert.Equal(t, id, parsedQuery.Id)

			res := types.ShardResponse{Id: id}
			r, _ := json.Marshal(res)
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := security.New(k).Role.Delete(id, nil)

	assert.Equal(t, id, res)
}
*/

/*
func ExampleSecurityRole_Delete() {
	id := "roleId"
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, err := security.New(k).Role.Delete(id, nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res)
}
*/
