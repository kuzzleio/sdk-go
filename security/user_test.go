package security_test

import (
  "testing"
  "github.com/kuzzleio/sdk-go/internal"
  "github.com/kuzzleio/sdk-go/kuzzle"
  "github.com/stretchr/testify/assert"
  "encoding/json"
  "github.com/kuzzleio/sdk-go/types"
  "github.com/kuzzleio/sdk-go/security"
)

func TestFetchUserEmptyId(t *testing.T) {
  c := &internal.MockedConnection{
    MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
      return types.KuzzleResponse{Error: types.MessageError{Message: "Security.FetchUser: user id required"}}
    },
  }
  k, _ := kuzzle.NewKuzzle(c, nil)

  _, err := security.NewSecurity(k).FetchUser("", nil)
  assert.NotNil(t, err)
}

func TestFetchUserError(t *testing.T) {
  c := &internal.MockedConnection{
    MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
      return types.KuzzleResponse{Error: types.MessageError{Message: "Unit test error"}}
    },
  }
  k, _ := kuzzle.NewKuzzle(c, nil)

  _, err := security.NewSecurity(k).FetchUser("userId", nil)
  assert.NotNil(t, err)
}

func TestFetchUser(t *testing.T) {
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

  res, _ := security.NewSecurity(k).FetchUser(id, nil)

  assert.Equal(t, id, res.Id)

  assert.Equal(t, []string{"admin", "other"}, res.ProfileIDs())

  assert.Equal(t, "Luke", res.Content("name"))
  assert.Equal(t, "Jedi", res.Content("function"))

  contentAsMap := make(map[string]interface{})
  contentAsMap["name"] = "Luke"
  contentAsMap["function"] = "Jedi"

  assert.Equal(t, contentAsMap, res.ContentMap("name", "function"))
}
