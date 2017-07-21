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
  hits := make([]types.KuzzleResult, 1)
  hits[0] = types.KuzzleResult{Id: "doc42", Source: json.RawMessage(`{"foo":"bar"}`)}
  var results = types.KuzzleSearchResult{Hits: hits}

  c := &internal.MockedConnection{
    MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
      parsedQuery := &types.KuzzleRequest{}
      json.Unmarshal(query, parsedQuery)

      assert.Equal(t, "document", parsedQuery.Controller)
      assert.Equal(t, "search", parsedQuery.Action)
      assert.Equal(t, "index", parsedQuery.Index)
      assert.Equal(t, "collection", parsedQuery.Collection)

      res := types.KuzzleSearchResult{Total: 42, Hits: results.Hits}
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

func TestSearchWithScroll(t *testing.T) {
  hits := make([]types.KuzzleResult, 1)
  hits[0] = types.KuzzleResult{Id: "doc42", Source: json.RawMessage(`{"foo":"bar"}`)}
  var results = types.KuzzleSearchResult{Hits: hits}

  c := &internal.MockedConnection{
    MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
      parsedQuery := &types.KuzzleRequest{}
      json.Unmarshal(query, parsedQuery)

      assert.Equal(t, "document", parsedQuery.Controller)
      assert.Equal(t, "search", parsedQuery.Action)
      assert.Equal(t, "index", parsedQuery.Index)
      assert.Equal(t, "collection", parsedQuery.Collection)
      assert.Equal(t, 2, parsedQuery.From)
      assert.Equal(t, 4, parsedQuery.Size)
      assert.Equal(t, "1m", parsedQuery.Scroll)

      res := types.KuzzleSearchResult{Total: 42, Hits: results.Hits, ScrollId: "f00b4r"}
      r, _ := json.Marshal(res)
      return types.KuzzleResponse{Result: r}
    },
  }
  k, _ := kuzzle.NewKuzzle(c, nil)

  res, _ := collection.NewCollection(k,"collection", "index").Search(nil, &types.Options{From: 2, Size: 4, Scroll: "1m"})
  assert.Equal(t, 42, res.Total)
  assert.Equal(t, hits, res.Hits)
  assert.Equal(t, "f00b4r", res.ScrollId)
  assert.Equal(t, res.Hits[0].Id, "doc42")
  assert.Equal(t, res.Hits[0].Source, json.RawMessage(`{"foo":"bar"}`))
}