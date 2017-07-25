package kuzzle

import (
  "testing"
  "github.com/kuzzleio/sdk-go/internal"
  "encoding/json"
  "github.com/kuzzleio/sdk-go/types"
  "github.com/stretchr/testify/assert"
)

func TestUnsetJwt(t *testing.T) {
  c := &internal.MockedConnection{
    MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
      request := types.KuzzleRequest{}
      json.Unmarshal(query, &request)

      assert.Equal(t, "auth", request.Controller)
      assert.Equal(t, "login", request.Action)
      assert.Equal(t, 0, request.ExpiresIn)

      type loginResult struct {
        Jwt string `json:"jwt"`
      }

      loginRes := loginResult{"token"}
      marsh, _ := json.Marshal(loginRes)

      return types.KuzzleResponse{Result: marsh}
    },
  }

  k, _ := NewKuzzle(c, nil)

  res, _ :=k.Login("local", nil, nil)
  assert.Equal(t, "token", res)
  assert.Equal(t, "token", k.jwt)
  k.UnsetJwt()
  assert.Equal(t, "", k.jwt)
}