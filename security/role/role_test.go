package role_test

import (
	"encoding/json"
	"fmt"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/security"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
	"github.com/kuzzleio/sdk-go/security/role"
)

func TestRoleSetContent(t *testing.T) {
	id := "roleId"

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "getRole", parsedQuery.Action)
			assert.Equal(t, id, parsedQuery.Id)

			res := role.Role{Id: id, Source: []byte(`{"controllers":{"*":{"actions":{"*":true}}}}`)}
			r, _ := json.Marshal(res)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	r, _ := security.NewSecurity(k).Role.Fetch(id, nil)

	newContent := types.Controllers{Controllers: map[string]types.Controller{"document": {Actions: map[string]bool{"get": true}}}}

	r.SetContent(newContent)

	assert.Equal(t, newContent.Controllers, r.Controllers())
}

func TestRoleSave(t *testing.T) {
	id := "roleId"
	callCount := 0

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			if callCount == 0 {
				callCount++
				assert.Equal(t, "security", parsedQuery.Controller)
				assert.Equal(t, "getRole", parsedQuery.Action)
				assert.Equal(t, id, parsedQuery.Id)

				res := role.Role{Id: id, Source: []byte(`{"controllers":{"*":{"actions":{"*":true}}}}`)}
				r, _ := json.Marshal(res)
				return types.KuzzleResponse{Result: r}
			}
			if callCount == 1 {
				callCount++
				assert.Equal(t, "security", parsedQuery.Controller)
				assert.Equal(t, "createOrReplaceRole", parsedQuery.Action)
				assert.Equal(t, "replace", options.GetIfExist())
				assert.Equal(t, id, parsedQuery.Id)

				res := role.Role{Id: id, Source: []byte(`{"controllers":{"document":{"actions":{"get":true}}}}`)}
				r, _ := json.Marshal(res)
				return types.KuzzleResponse{Result: r}
			}

			return types.KuzzleResponse{Result: nil}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	r, _ := security.NewSecurity(k).Role.Fetch(id, nil)

	newContent := types.Controllers{Controllers: map[string]types.Controller{"document": {Actions: map[string]bool{"get": true}}}}

	r.SetContent(newContent)
	r.Save(nil)

	assert.Equal(t, newContent.Controllers, r.Controllers())
}

func TestRoleUpdate(t *testing.T) {
	id := "roleId"
	callCount := 0

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			if callCount == 0 {
				callCount++
				assert.Equal(t, "security", parsedQuery.Controller)
				assert.Equal(t, "getRole", parsedQuery.Action)
				assert.Equal(t, id, parsedQuery.Id)

				res := role.Role{Id: id, Source: []byte(`{"controllers":{"*":{"actions":{"*":true}}}}`)}
				r, _ := json.Marshal(res)
				return types.KuzzleResponse{Result: r}
			}
			if callCount == 1 {
				callCount++
				assert.Equal(t, "security", parsedQuery.Controller)
				assert.Equal(t, "updateRole", parsedQuery.Action)
				assert.Equal(t, id, parsedQuery.Id)

				res := role.Role{Id: id, Source: []byte(`{"controllers":{"document":{"actions":{"get":true}}}}`)}
				r, _ := json.Marshal(res)
				return types.KuzzleResponse{Result: r}
			}

			return types.KuzzleResponse{Result: nil}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	r, _ := security.NewSecurity(k).Role.Fetch(id, nil)

	newContent := types.Controllers{Controllers: map[string]types.Controller{"document": {Actions: map[string]bool{"get": true}}}}

	updatedRole, _ := r.Update(newContent,nil)

	assert.Equal(t, newContent.Controllers, updatedRole.Controllers())
}

func TestRoleDelete(t *testing.T) {
	id := "SomeMenJustWantToWatchTheWorldBurn"
	callCount := 0

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			if callCount == 0 {
				callCount++
				assert.Equal(t, "security", parsedQuery.Controller)
				assert.Equal(t, "getRole", parsedQuery.Action)
				assert.Equal(t, id, parsedQuery.Id)

				res := role.Role{Id: id, Source: []byte(`{"controllers":{"*":{"actions":{"*":true}}}}`)}
				r, _ := json.Marshal(res)
				return types.KuzzleResponse{Result: r}
			}
			if callCount == 1 {
				callCount++
				assert.Equal(t, "security", parsedQuery.Controller)
				assert.Equal(t, "deleteRole", parsedQuery.Action)
				assert.Equal(t, id, parsedQuery.Id)

				res := types.ShardResponse{Id: id}
				r, _ := json.Marshal(res)
				return types.KuzzleResponse{Result: r}
			}

			return types.KuzzleResponse{Result: nil}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	r, _ := security.NewSecurity(k).Role.Fetch(id, nil)

	inTheEnd, _ := r.Delete(nil)

	assert.Equal(t, id, inTheEnd)
}

func TestFetchEmptyId(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Security.Role.Fetch: role id required"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := security.NewSecurity(k).Role.Fetch("", nil)
	assert.NotNil(t, err)
}

func TestFetchError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := security.NewSecurity(k).Role.Fetch("roleId", nil)
	assert.NotNil(t, err)
}

func TestFetch(t *testing.T) {
	id := "roleId"

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "getRole", parsedQuery.Action)
			assert.Equal(t, id, parsedQuery.Id)

			res := role.Role{Id: id, Source: []byte(`{"controllers":{"*":{"actions":{"*":true}}}}`)}
			r, _ := json.Marshal(res)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := security.NewSecurity(k).Role.Fetch(id, nil)

	assert.Equal(t, id, res.Id)
	assert.Equal(t, true, res.Controllers()["*"].Actions["*"])
}

func TestSearchError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := security.NewSecurity(k).Role.Search(nil, nil)
	assert.NotNil(t, err)
}

func TestSearch(t *testing.T) {
	hits := make([]role.Role, 1)
	hits[0] = role.Role{Id: "role42", Source: json.RawMessage(`{"controllers":{"*":{"actions":{"*":true}}}}`)}
	var results = role.RoleSearchResult{Total: 42, Hits: hits}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "searchRoles", parsedQuery.Action)

			res := role.RoleSearchResult{Total: results.Total, Hits: results.Hits}
			r, _ := json.Marshal(res)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := security.NewSecurity(k).Role.Search(nil, nil)
	assert.Equal(t, results.Total, res.Total)
	assert.Equal(t, hits, res.Hits)
	assert.Equal(t, res.Hits[0].Id, "role42")
	assert.Equal(t, res.Hits[0].Source, json.RawMessage(`{"controllers":{"*":{"actions":{"*":true}}}}`))
	assert.Equal(t, res.Hits[0].Controllers()["*"].Actions["*"], true)
}

func TestSearchWithOptions(t *testing.T) {
	hits := make([]role.Role, 1)
	hits[0] = role.Role{Id: "role42", Source: json.RawMessage(`{"controllers":{"*":{"actions":{"*":true}}}}`)}
	var results = role.RoleSearchResult{Total: 42, Hits: hits}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "searchRoles", parsedQuery.Action)

			res := role.RoleSearchResult{Total: results.Total, Hits: results.Hits}
			r, _ := json.Marshal(res)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	opts := types.NewQueryOptions()
	opts.SetFrom(2)
	opts.SetSize(4)
	opts.SetScroll("1m")

	res, _ := security.NewSecurity(k).Role.Search(nil, opts)
	assert.Equal(t, results.Total, res.Total)
	assert.Equal(t, hits, res.Hits)
	assert.Equal(t, res.Hits[0].Id, "role42")
	assert.Equal(t, res.Hits[0].Source, json.RawMessage(`{"controllers":{"*":{"actions":{"*":true}}}}`))
	assert.Equal(t, res.Hits[0].Controllers()["*"].Actions["*"], true)
}

func TestCreateEmptyId(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Security.Role.Create: role id required"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := security.NewSecurity(k).Role.Create("", types.Controllers{}, nil)
	assert.NotNil(t, err)
}

func TestCreateError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := security.NewSecurity(k).Role.Create("roleId", types.Controllers{}, nil)
	assert.NotNil(t, err)
}

func TestCreate(t *testing.T) {
	id := "roleId"

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
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
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := security.NewSecurity(k).Role.Create(id, types.Controllers{Controllers: map[string]types.Controller{"*": {Actions: map[string]bool{"*": true}}}}, nil)

	assert.Equal(t, id, res.Id)
	assert.Equal(t, true, res.Controllers()["*"].Actions["*"])
}

func TestCreateIfExists(t *testing.T) {
	id := "roleId"

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
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
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	opts := types.NewQueryOptions()
	opts.SetIfExist("replace")

	res, _ := security.NewSecurity(k).Role.Create(id, types.Controllers{Controllers: map[string]types.Controller{"*": {Actions: map[string]bool{"*": true}}}}, opts)

	assert.Equal(t, id, res.Id)
	assert.Equal(t, true, res.Controllers()["*"].Actions["*"])
}

func TestCreateWithStrictOption(t *testing.T) {
	id := "roleId"

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
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
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	opts := types.NewQueryOptions()
	opts.SetIfExist("error")

	res, _ := security.NewSecurity(k).Role.Create(id, types.Controllers{Controllers: map[string]types.Controller{"*": {Actions: map[string]bool{"*": true}}}}, opts)

	assert.Equal(t, id, res.Id)
	assert.Equal(t, true, res.Controllers()["*"].Actions["*"])
}

func TestCreateWithWrongOption(t *testing.T) {
	id := "roleId"

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	opts := types.NewQueryOptions()
	opts.SetIfExist("unknown")

	_, err := security.NewSecurity(k).Role.Create(id, types.Controllers{Controllers: map[string]types.Controller{"*": {Actions: map[string]bool{"*": true}}}}, opts)

	assert.Equal(t, "Invalid value for the 'ifExist' option: 'unknown'", fmt.Sprint(err))
}

func TestUpdateEmptyId(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Security.Role.Update: role id required"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := security.NewSecurity(k).Role.Update("", types.Controllers{}, nil)
	assert.NotNil(t, err)
}

func TestUpdateError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := security.NewSecurity(k).Role.Update("roleId", types.Controllers{}, nil)
	assert.NotNil(t, err)
}

func TestUpdate(t *testing.T) {
	id := "roleId"

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
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
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := security.NewSecurity(k).Role.Update(id, types.Controllers{Controllers: map[string]types.Controller{"*": {Actions: map[string]bool{"*": true}}}}, nil)

	assert.Equal(t, id, res.Id)
	assert.Equal(t, true, res.Controllers()["*"].Actions["*"])
}

func TestDeleteEmptyId(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Security.Role.Delete: role id required"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := security.NewSecurity(k).Role.Delete("", nil)
	assert.NotNil(t, err)
}

func TestDeleteError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := security.NewSecurity(k).Role.Delete("roleId", nil)
	assert.NotNil(t, err)
}

func TestDelete(t *testing.T) {
	id := "roleId"

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "deleteRole", parsedQuery.Action)
			assert.Equal(t, id, parsedQuery.Id)

			res := types.ShardResponse{Id: id}
			r, _ := json.Marshal(res)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := security.NewSecurity(k).Role.Delete(id, nil)

	assert.Equal(t, id, res)
}
