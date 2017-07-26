package user_test

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
      return types.KuzzleResponse{Error: types.MessageError{Message: "Security.User.Fetch: user id required"}}
    },
  }
  k, _ := kuzzle.NewKuzzle(c, nil)

  _, err := security.NewSecurity(k).User.Fetch("", nil)
  assert.NotNil(t, err)
}

func TestFetchError(t *testing.T) {
  c := &internal.MockedConnection{
    MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
      return types.KuzzleResponse{Error: types.MessageError{Message: "Unit test error"}}
    },
  }
  k, _ := kuzzle.NewKuzzle(c, nil)

  _, err := security.NewSecurity(k).User.Fetch("userId", nil)
  assert.NotNil(t, err)
}

func TestFetch(t *testing.T) {
  id := "userId"

  c := &internal.MockedConnection{
    MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
      parsedQuery := &types.KuzzleRequest{}
      json.Unmarshal(query, parsedQuery)

      assert.Equal(t, "security", parsedQuery.Controller)
      assert.Equal(t, "getUser", parsedQuery.Action)
      assert.Equal(t, id, parsedQuery.Id)

      res := types.User{Id: id, Source: []byte(`{"profileIds":["admin","other"],"name":"Luke","function":"Jedi"}`)}
      r, _ := json.Marshal(res)
      return types.KuzzleResponse{Result: r}
    },
  }
  k, _ := kuzzle.NewKuzzle(c, nil)

  res, _ := security.NewSecurity(k).User.Fetch(id, nil)

  assert.Equal(t, id, res.Id)

  assert.Equal(t, []string{"admin", "other"}, res.ProfileIDs())

  assert.Equal(t, "Luke", res.Content("name"))
  assert.Equal(t, "Jedi", res.Content("function"))

  contentAsMap := make(map[string]interface{})
  contentAsMap["name"] = "Luke"
  contentAsMap["function"] = "Jedi"

  assert.Equal(t, contentAsMap, res.ContentMap("name", "function"))
}

func TestSearchError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := security.NewSecurity(k).User.Search(nil, nil)
	assert.NotNil(t, err)
}

func TestSearch(t *testing.T) {
	hits := make([]types.User, 1)
	hits[0] = types.User{Id: "user42", Source: json.RawMessage(`{"profileIds":["admin","other"],"foo":"bar"}`)}
	var results = types.KuzzleSearchUsersResult{Total: 42, Hits: hits}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "searchUsers", parsedQuery.Action)

			res := types.KuzzleSearchUsersResult{Total: results.Total, Hits: results.Hits}
			r, _ := json.Marshal(res)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := security.NewSecurity(k).User.Search(nil, nil)
	assert.Equal(t, results.Total, res.Total)
	assert.Equal(t, hits, res.Hits)
	assert.Equal(t, res.Hits[0].Id, "user42")
	assert.Equal(t, res.Hits[0].Source, json.RawMessage(`{"profileIds":["admin","other"],"foo":"bar"}`))
	assert.Equal(t, res.Hits[0].ProfileIDs(), []string{"admin", "other"})
	assert.Equal(t, res.Hits[0].Content("foo"), "bar")
}

func TestSearchWithScroll(t *testing.T) {
	hits := make([]types.User, 1)
	hits[0] = types.User{Id: "user42", Source: json.RawMessage(`{"profileIds":["admin","other"],"foo":"bar"}`)}
	var results = types.KuzzleSearchUsersResult{Total: 42, Hits: hits}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "searchUsers", parsedQuery.Action)

			res := types.KuzzleSearchUsersResult{Total: results.Total, Hits: results.Hits, ScrollId: "f00b4r"}
			r, _ := json.Marshal(res)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := security.NewSecurity(k).User.Search(nil, &types.Options{Size: 2, From: 4, Scroll: "1m"})
	assert.Equal(t, results.Total, res.Total)
	assert.Equal(t, hits, res.Hits)
	assert.Equal(t, "f00b4r", res.ScrollId)
	assert.Equal(t, res.Hits[0].Id, "user42")
	assert.Equal(t, res.Hits[0].Source, json.RawMessage(`{"profileIds":["admin","other"],"foo":"bar"}`))
	assert.Equal(t, res.Hits[0].ProfileIDs(), []string{"admin", "other"})
	assert.Equal(t, res.Hits[0].Content("foo"), "bar")
}

func TestScrollEmptyScrollId(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Security.User.Scroll: scroll id required"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := security.NewSecurity(k).User.Scroll("", nil)
	assert.NotNil(t, err)
}

func TestScrollError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := security.NewSecurity(k).User.Scroll("f00b4r", nil)
	assert.NotNil(t, err)
}

func TestScroll(t *testing.T) {
	type response struct {
		Total int          `json:"total"`
		Hits  []types.User `json:"hits"`
	}

	hits := make([]types.User, 1)
	hits[0] = types.User{Id: "user42", Source: json.RawMessage(`{"profileIds":["admin","other"],"foo":"bar"}`)}
	var results = types.KuzzleSearchUsersResult{Total: 42, Hits: hits}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "scrollUsers", parsedQuery.Action)

			res := response{Total: results.Total, Hits: results.Hits}
			r, _ := json.Marshal(res)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := security.NewSecurity(k).User.Scroll("f00b4r", nil)
	assert.Equal(t, results.Total, res.Total)
	assert.Equal(t, hits, res.Hits)
	assert.Equal(t, res.Hits[0].Id, "user42")
	assert.Equal(t, res.Hits[0].Source, json.RawMessage(`{"profileIds":["admin","other"],"foo":"bar"}`))
	assert.Equal(t, res.Hits[0].ProfileIDs(), []string{"admin", "other"})
	assert.Equal(t, res.Hits[0].Content("foo"), "bar")
}
