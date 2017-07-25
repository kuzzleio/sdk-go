package kuzzle_test

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "github.com/kuzzleio/sdk-go/internal"
  "github.com/kuzzleio/sdk-go/types"
  "encoding/json"
  "github.com/kuzzleio/sdk-go/kuzzle"
  "log"
)

func TestGetStatisticsQueryError(t *testing.T) {
  c := &internal.MockedConnection{
    MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
      request := types.KuzzleRequest{}
      json.Unmarshal(query, &request)
      assert.Equal(t, "server", request.Controller)
      assert.Equal(t, "getLastStats", request.Action)
      return types.KuzzleResponse{Error: types.MessageError{Message: "error"}}
    },
  }
  k, _ := kuzzle.NewKuzzle(c, nil)
  _, err := k.GetStatistics(nil)
  assert.NotNil(t, err)
}

func TestGetStatistics(t *testing.T) {
  c := &internal.MockedConnection{
    MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
      request := types.KuzzleRequest{}
      json.Unmarshal(query, &request)
      assert.Equal(t, "server", request.Controller)
      assert.Equal(t, "getLastStats", request.Action)

      m := make(map[string]int)
      m["websocket"] = 42

      stats := types.Statistics{
        CompletedRequests: m,
      }

      h, err := json.Marshal(stats)
      if err != nil {
        log.Fatal(err)
      }

      return types.KuzzleResponse{Result: h}
    },
  }
  k, _ := kuzzle.NewKuzzle(c, nil)

  res, _ := k.GetStatistics(nil)

  assert.Equal(t, 42, res.CompletedRequests["websocket"])
}
