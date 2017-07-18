package kuzzle_test

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "github.com/kuzzleio/sdk-go/internal"
  "github.com/kuzzleio/sdk-go/types"
  "encoding/json"
  "github.com/kuzzleio/sdk-go/kuzzle"
  "fmt"
)

func TestDeleteMyCredentialsQueryError(t *testing.T) {
  c := &internal.MockedConnection{
    MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
      request := types.KuzzleRequest{}
      json.Unmarshal(query, &request)
      assert.Equal(t, "auth", request.Controller)
      assert.Equal(t, "deleteMyCredentials", request.Action)
      return types.KuzzleResponse{Error: types.MessageError{Message: "error"}}
    },
  }
  k, _ := kuzzle.NewKuzzle(c, nil)
  _, err := k.DeleteMyCredentials("local", nil, nil)
  assert.NotNil(t, err)
}

func TestDeleteMyCredentials(t *testing.T) {
  type myCredentials struct {
    Username string `json:"username"`
    Password string `json:"password"`
  }

  c := &internal.MockedConnection{
    MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
      type ackResult struct {
        Acknowledged       bool
        ShardsAcknowledged bool
      }

      ack := ackResult{Acknowledged: true, ShardsAcknowledged: true}
      r, _ := json.Marshal(ack)

      request := types.KuzzleRequest{}
      json.Unmarshal(query, &request)
      assert.Equal(t, "auth", request.Controller)
      assert.Equal(t, "deleteMyCredentials", request.Action)
      return types.KuzzleResponse{Result: r}
    },
  }
  k, _ := kuzzle.NewKuzzle(c, nil)

  res, _ := k.DeleteMyCredentials("local", myCredentials{"foo", "bar"}, nil)

  fmt.Printf("%s\n", res)
  assert.Equal(t, true, res.Acknowledged)
  assert.Equal(t, true, res.ShardsAcknowledged)
}
