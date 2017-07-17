package kuzzle_test

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "github.com/kuzzleio/sdk-go/internal"
  "github.com/kuzzleio/sdk-go/types"
  "encoding/json"
  "github.com/kuzzleio/sdk-go/kuzzle"
)

func TestCreateIndexNull(t *testing.T) {
  k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
  _, err := k.CreateIndex("", nil)
  assert.NotNil(t, err)
}

func TestCreateIndexQueryError(t *testing.T) {
  c := &internal.MockedConnection{
    MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
      return types.KuzzleResponse{Error: types.MessageError{Message: "error"}}
    },
  }
  k, _ := kuzzle.NewKuzzle(c, nil)
  _, err := k.CreateIndex("index", nil)
  assert.NotNil(t, err)
}

func TestCreateIndex(t *testing.T) {
  type ackResult struct {
    Acknowledged bool
    ShardsAcknowledged bool
  }

  c := &internal.MockedConnection{
    MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
      ack := ackResult{Acknowledged: true, ShardsAcknowledged: true}
      r, _ := json.Marshal(ack)
      return types.KuzzleResponse{Result: r}
    },
  }
  k, _ := kuzzle.NewKuzzle(c, nil)

  res, _ := k.CreateIndex("index", nil)
  assert.Equal(t, true, res.Acknowledged)
  assert.Equal(t, true, res.ShardsAcknowledged)
}