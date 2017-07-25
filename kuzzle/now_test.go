package kuzzle_test

import (
  "github.com/kuzzleio/sdk-go/internal"
  "testing"
  "github.com/stretchr/testify/assert"
  "github.com/kuzzleio/sdk-go/kuzzle"
  "github.com/kuzzleio/sdk-go/types"
  "encoding/json"
)

func TestNowQueryError(t *testing.T) {
  c := &internal.MockedConnection{
    MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
      return types.KuzzleResponse{Error: types.MessageError{Message: "error"}}
    },
  }
  k, _ := kuzzle.NewKuzzle(c, nil)
  k.Connect()
  _, err := k.Now(nil)
  assert.NotNil(t, err)
}

func TestNow(t *testing.T) {
  type now struct {
    Now int `json:"now"`
  }

  c := &internal.MockedConnection{
    MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
      n := now{Now: 1500646351073}

      marsh, _ := json.Marshal(n)
      return types.KuzzleResponse{Result: marsh}
    },
  }
  k, _ := kuzzle.NewKuzzle(c, nil)

  res, _ := k.Now(nil)
  assert.Equal(t, 1500646351073, res)
}
