package kuzzle_test

import (
  "testing"
  "github.com/kuzzleio/sdk-go/internal"
  "github.com/kuzzleio/sdk-go/types"
  "encoding/json"
  "github.com/kuzzleio/sdk-go/kuzzle"
  "github.com/stretchr/testify/assert"
)

func TestCount(t *testing.T) {
  type result struct {
    Count int `json:"count"`
  }

  c := &internal.MockedConnection{
    MockSend: func() types.KuzzleResponse {
      res := result{Count: 10}
      r, _ := json.Marshal(res)
      return types.KuzzleResponse{Result: r}
    },
  }
  k, _ := kuzzle.NewKuzzle(c, nil)

  res, _ := k.Collection("collection", "index").Count(nil, nil)
  assert.Equal(t, 10, *res)
}
