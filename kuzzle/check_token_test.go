package kuzzle_test

import (
  "github.com/kuzzleio/sdk-go/internal"
  "testing"
  "github.com/stretchr/testify/assert"
  "github.com/kuzzleio/sdk-go/kuzzle"
  "encoding/json"
  "github.com/kuzzleio/sdk-go/types"
)

func TestCheckTokenTokenNull(t *testing.T) {
  k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
  _, err := k.CheckToken("")
  assert.NotNil(t, err)
}

func TestCheckTokenQueryError(t *testing.T) {
  c := &internal.MockedConnection{
    MockSend: func(query []byte) types.KuzzleResponse {
      return types.KuzzleResponse{Error: types.MessageError{Message: "error"}}
    },
  }
  k, _ := kuzzle.NewKuzzle(c, nil)
  _, err := k.CheckToken("token")
  assert.NotNil(t, err)
}

func TestCheckToken(t *testing.T) {
  c := &internal.MockedConnection{
    MockSend: func(query []byte) types.KuzzleResponse {
      tokenValidity := kuzzle.TokenValidity{Valid: true}
      r, _ := json.Marshal(tokenValidity)
      return types.KuzzleResponse{Result: r}
    },
  }
  k, _ := kuzzle.NewKuzzle(c, nil)


  type ackResult struct {
    Acknowledged bool
    ShardsAcknowledged bool
  }
  res, _ := k.CheckToken("token")
  assert.Equal(t, true, res.Valid)
}

