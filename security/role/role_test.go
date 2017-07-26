package role_test

import (
  "testing"
  "github.com/kuzzleio/sdk-go/internal"
  "github.com/kuzzleio/sdk-go/kuzzle"
  "github.com/stretchr/testify/assert"
  "encoding/json"
  "github.com/kuzzleio/sdk-go/types"
  "github.com/kuzzleio/sdk-go/security"
)

func TestFetchEmptyId(t *testing.T) {
  c := &internal.MockedConnection{
    MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
      return types.KuzzleResponse{Error: types.MessageError{Message: "Security.Role.Fetch: role id required"}}
    },
  }
  k, _ := kuzzle.NewKuzzle(c, nil)

  _, err := security.NewSecurity(k).Role.Fetch("", nil)
  assert.NotNil(t, err)
}

func TestFetchError(t *testing.T) {
  c := &internal.MockedConnection{
    MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
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
    MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
      parsedQuery := &types.KuzzleRequest{}
      json.Unmarshal(query, parsedQuery)

      assert.Equal(t, "security", parsedQuery.Controller)
      assert.Equal(t, "getRole", parsedQuery.Action)
      assert.Equal(t, id, parsedQuery.Id)

      res := types.Role{Id: id, Source: []byte(`{"controllers":{"*":{"actions":{"*":true}}}}`)}
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
		MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := security.NewSecurity(k).Role.Search(nil, nil)
	assert.NotNil(t, err)
}

func TestSearch(t *testing.T) {
	hits := make([]types.Role, 1)
	hits[0] = types.Role{Id: "role42", Source: json.RawMessage(`{"controllers":{"*":{"actions":{"*":true}}}}`)}
	var results = types.KuzzleSearchRolesResult{Total: 42, Hits: hits}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "searchRoles", parsedQuery.Action)

			res := types.KuzzleSearchRolesResult{Total: results.Total, Hits: results.Hits}
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
	hits := make([]types.Role, 1)
	hits[0] = types.Role{Id: "role42", Source: json.RawMessage(`{"controllers":{"*":{"actions":{"*":true}}}}`)}
	var results = types.KuzzleSearchRolesResult{Total: 42, Hits: hits}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "searchRoles", parsedQuery.Action)

			res := types.KuzzleSearchRolesResult{Total: results.Total, Hits: results.Hits}
			r, _ := json.Marshal(res)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := security.NewSecurity(k).Role.Search(nil, &types.Options{From: 2, Size: 4})
	assert.Equal(t, results.Total, res.Total)
	assert.Equal(t, hits, res.Hits)
	assert.Equal(t, res.Hits[0].Id, "role42")
	assert.Equal(t, res.Hits[0].Source, json.RawMessage(`{"controllers":{"*":{"actions":{"*":true}}}}`))
	assert.Equal(t, res.Hits[0].Controllers()["*"].Actions["*"], true)
}
