package kuzzle_test

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "github.com/kuzzleio/sdk-go/internal"
  "github.com/kuzzleio/sdk-go/types"
  "encoding/json"
  "github.com/kuzzleio/sdk-go/kuzzle"
)

func TestCreateMyCredentialsQueryError(t *testing.T) {
  c := &internal.MockedConnection{
    MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
      request := types.KuzzleRequest{}
      json.Unmarshal(query, &request)
      assert.Equal(t, "auth", request.Controller)
      assert.Equal(t, "createMyCredentials", request.Action)
      return types.KuzzleResponse{Error: types.MessageError{Message: "error"}}
    },
  }
  k, _ := kuzzle.NewKuzzle(c, nil)
  _, err := k.CreateMyCredentials("local", nil, nil)
  assert.NotNil(t, err)
}

func TestCreateMyCredentials(t *testing.T) {
  type myCredentials struct {
    Username string `json:"username"`
    Password string `json:"password"`
  }

  c := &internal.MockedConnection{
    MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
      ack := myCredentials{Username: "foo", Password: "bar"}
      r, _ := json.Marshal(ack)

      request := types.KuzzleRequest{}
      json.Unmarshal(query, &request)
      assert.Equal(t, "auth", request.Controller)
      assert.Equal(t, "createMyCredentials", request.Action)
      return types.KuzzleResponse{Result: r}
    },
  }
  k, _ := kuzzle.NewKuzzle(c, nil)

  res, _ := k.CreateMyCredentials("local", myCredentials{"foo", "bar"}, nil)

  assert.Equal(t, "foo", res["username"])
  assert.Equal(t, "bar", res["password"])
}
