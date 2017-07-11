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
  _, err := kuzzle.CreateIndex(nil, "", nil)
  assert.NotNil(t, err)
}

func TestCreateIndexQueryError(t *testing.T) {
  k := &internal.MockedKuzzle{
    MockQuery: func() types.KuzzleResponse {
      return types.KuzzleResponse{Error: types.MessageError{Message: "error"}}
    },
  }

  _, err := kuzzle.CreateIndex(k, "index", nil)
  assert.NotNil(t, err)
}

func TestCreateIndex(t *testing.T) {
  type ackResult struct {
    Acknowledged bool
    ShardsAcknowledged bool
  }

  k := &internal.MockedKuzzle{
    MockQuery: func() types.KuzzleResponse {
      ack := ackResult{Acknowledged: true, ShardsAcknowledged: true}
      r, _ := json.Marshal(ack)
      return types.KuzzleResponse{Result: r}
    },
  }

  res, _ := kuzzle.CreateIndex(k, "index", nil)
  assert.Equal(t, true, res.Acknowledged)
  assert.Equal(t, true, res.ShardsAcknowledged)
}