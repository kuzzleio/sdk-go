package kuzzle_test

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "github.com/kuzzleio/sdk-go/internal"
  "github.com/kuzzleio/sdk-go/types"
  "encoding/json"
  "github.com/kuzzleio/sdk-go/kuzzle"
)

func TestGetMyCredentialsQueryError(t *testing.T) {
  c := &internal.MockedConnection{
    MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
      request := types.KuzzleRequest{}
      json.Unmarshal(query, &request)
      assert.Equal(t, "auth", request.Controller)
      assert.Equal(t, "getMyCredentials", request.Action)
      assert.Equal(t, "local", request.Strategy)
      return types.KuzzleResponse{Error: types.MessageError{Message: "error"}}
    },
  }
  k, _ := kuzzle.NewKuzzle(c, nil)
  _, err := k.GetMyCredentials("local", nil)
  assert.NotNil(t, err)
}

func TestGetMyCredentials(t *testing.T) {
  c := &internal.MockedConnection{
    MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
      request := types.KuzzleRequest{}
      json.Unmarshal(query, &request)
      assert.Equal(t, "auth", request.Controller)
      assert.Equal(t, "getMyCredentials", request.Action)
      assert.Equal(t, "local", request.Strategy)

      type myCredentials struct {
        Username string `json:"username"`
        Password string `json:"password"`
      }

      myCred := myCredentials{"admin", "test"}
      marsh, _ := json.Marshal(myCred)

      return types.KuzzleResponse{Result: marsh}
    },
  }

  k, _ := kuzzle.NewKuzzle(c, nil)
  res, err := k.GetMyCredentials("local", nil)
  assert.Nil(t, err)

  type myCredentials struct {
    Username string `json:"username"`
    Password string `json:"password"`
  }

  cred := myCredentials{}
  json.Unmarshal(res, &cred)

  assert.Equal(t, "admin", cred.Username)
  assert.Equal(t, "test", cred.Password)
}