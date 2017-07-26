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

func TestWhoAmIQueryError(t *testing.T) {
  c := &internal.MockedConnection{
    MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
      request := types.KuzzleRequest{}
      json.Unmarshal(query, &request)
      assert.Equal(t, "auth", request.Controller)
      assert.Equal(t, "getCurrentUser", request.Action)
      return types.KuzzleResponse{Error: types.MessageError{Message: "error"}}
    },
  }
  k, _ := kuzzle.NewKuzzle(c, nil)
  _, err := k.WhoAmI()
  assert.NotNil(t, err)
}

func TestWhoAmI(t *testing.T) {
  c := &internal.MockedConnection{
    MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
      request := types.KuzzleRequest{}
      json.Unmarshal(query, &request)
      assert.Equal(t, "auth", request.Controller)
      assert.Equal(t, "getCurrentUser", request.Action)

      toMarshal := types.User{Id: "id"}

      h, err := json.Marshal(toMarshal)
      if err != nil {
        log.Fatal(err)
      }

      return types.KuzzleResponse{Result: h}
    },
  }
  k, _ := kuzzle.NewKuzzle(c, nil)

  res, _ := k.WhoAmI()

  assert.Equal(t, "id", res.Id)
}
