package kuzzle_test

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "github.com/kuzzleio/sdk-go/internal"
  "github.com/kuzzleio/sdk-go/types"
  "encoding/json"
  "github.com/kuzzleio/sdk-go/kuzzle"
  "log"
)

func TestUpdateSelfQueryError(t *testing.T) {
  c := &internal.MockedConnection{
    MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
      return types.KuzzleResponse{Error: types.MessageError{Message: "error"}}
    },
  }
  k, _ := kuzzle.NewKuzzle(c, nil)
  _, err := k.UpdateSelf("index", nil)
  assert.NotNil(t, err)
}

func TestUpdateSelf(t *testing.T) {
  q:= struct{Username string `json:"username"`}{"foo"}

  c := &internal.MockedConnection{
    MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
      request := types.KuzzleRequest{}
      json.Unmarshal(query, &request)
      assert.Equal(t, "auth", request.Controller)
      assert.Equal(t, "updateSelf", request.Action)

      assert.Equal(t, "foo", request.Body.(map[string]interface{})["username"])

      u := &types.User{Id: "login"}

      h, err := json.Marshal(u)
      if err != nil {
        log.Fatal(err)
      }

      return types.KuzzleResponse{Result: h}
    },
  }
  k, _ := kuzzle.NewKuzzle(c, nil)

  res, _ := k.UpdateSelf(q, nil)

  assert.Equal(t, "login", res.Id)
}
