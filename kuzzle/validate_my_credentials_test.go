package kuzzle_test

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "github.com/kuzzleio/sdk-go/internal"
  "github.com/kuzzleio/sdk-go/types"
  "encoding/json"
  "github.com/kuzzleio/sdk-go/kuzzle"
)

func TestValidateMyCredentialsQueryError(t *testing.T) {
  c := &internal.MockedConnection{
    MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
      request := types.KuzzleRequest{}
      json.Unmarshal(query, &request)
      assert.Equal(t, "auth", request.Controller)
      assert.Equal(t, "validateMyCredentials", request.Action)
      assert.Equal(t, "local", request.Strategy)
      return types.KuzzleResponse{Error: types.MessageError{Message: "error"}}
    },
  }
  k, _ := kuzzle.NewKuzzle(c, nil)
  _, err := k.ValidateMyCredentials("local", nil, nil)
  assert.NotNil(t, err)
}

func TestValidateMyCredentials(t *testing.T) {
  c := &internal.MockedConnection{
    MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
      request := types.KuzzleRequest{}
      json.Unmarshal(query, &request)
      assert.Equal(t, "auth", request.Controller)
      assert.Equal(t, "validateMyCredentials", request.Action)
      assert.Equal(t, "local", request.Strategy)
      assert.Equal(t, "foo", request.Body.(map[string]interface{})["username"])
      assert.Equal(t, "bar", request.Body.(map[string]interface{})["password"])

      ret, _ := json.Marshal(true)
      return types.KuzzleResponse{Result: ret}
    },
  }

  k, _ := kuzzle.NewKuzzle(c, nil)
  res, err := k.ValidateMyCredentials("local", struct{Username string `json:"username"`; Password string `json:"password"`}{"foo", "bar"}, nil)
  assert.Nil(t, err)

  assert.Equal(t, true, res)
}
