package collection_test

import (
  "testing"
  "encoding/json"
  "github.com/kuzzleio/sdk-go/internal"
  "github.com/kuzzleio/sdk-go/kuzzle"
  "github.com/stretchr/testify/assert"
  "github.com/kuzzleio/sdk-go/types"
  "github.com/kuzzleio/sdk-go/collection"
)

func TestSearchError(t *testing.T) {
  c := &internal.MockedConnection{
    MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
      return types.KuzzleResponse{Error: types.MessageError{Message: "Unit test error"}}
    },
  }
  k, _ := kuzzle.NewKuzzle(c, nil)

  _, err := collection.NewCollection(k, "collection", "index").Search(nil, nil)
  assert.NotNil(t, err)
}

func TestSearch(t *testing.T) {
  type response struct {
    Total int `json:"total"`
    Hits  []types.KuzzleResult `json:"hits"`
  }

  hits := make([]types.KuzzleResult, 1)
  hits[0] = types.KuzzleResult{Id: "doc42", Source: json.RawMessage(`{"foo":"bar"}`)}
  var results = types.KuzzleSearchResult{Hits: hits}

  c := &internal.MockedConnection{
    MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
      res := response{Total: 42, Hits: results.Hits}
      r, _ := json.Marshal(res)
      return types.KuzzleResponse{Result: r}
    },
  }
  k, _ := kuzzle.NewKuzzle(c, nil)

  res, _ := collection.NewCollection(k, "collection", "index").Search(nil, nil)
  assert.Equal(t, 42, res.Total)
  assert.Equal(t, hits, res.Hits)
  assert.Equal(t, res.Hits[0].Id, "doc42")
  assert.Equal(t, res.Hits[0].Source, json.RawMessage(`{"foo":"bar"}`))
}
