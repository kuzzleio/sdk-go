package kuzzle

import (
  "github.com/kuzzleio/sdk-go/types"
  "encoding/json"
  "errors"
)

type Collection struct {
  kuzzle            *Kuzzle
  index, collection string
  subscribeCallback interface{}
}

func NewCollection(kuzzle *Kuzzle, collection, index string) *Collection {
  return &Collection{
    index:      index,
    collection: collection,
    kuzzle:     kuzzle,
  }
}

/*
  Returns the number of documents matching the provided set of filters.

  There is a small delay between documents creation and their existence in our advanced search layer,
  usually a couple of seconds.
  That means that a document that was just been created won’t be returned by this function
*/
func (dc Collection) Count(filters interface{}, options *types.Options) (*int, error) {
  type countResult struct {
    Count int `json:"count"`
  }

  ch := make(chan types.KuzzleResponse)

  go dc.kuzzle.Query(buildQuery(dc.collection, dc.index, "document", "count", filters), options, ch)

  res := <-ch

  if res.Error.Message != "" {
    return nil, errors.New(res.Error.Message)
  }
  result := &countResult{}
  json.Unmarshal(res.Result, result)

  return &result.Count, nil
}

/*
  Create a new empty data collection, with no associated mapping.
*/
func (dc Collection) Create(options *types.Options) (*types.AckResponse, error) {
  ch := make(chan types.KuzzleResponse)

  go dc.kuzzle.Query(buildQuery(dc.collection, dc.index, "collection", "create", nil), options, ch)

  res := <-ch

  if res.Error.Message != "" {
    return nil, errors.New(res.Error.Message)
  }

  ack := &types.AckResponse{}
  json.Unmarshal(res.Result, &ack)

  return ack, nil
}

/*
  Searches documents in the given Collection, using provided filters and options.
*/
func (dc Collection) Search(filters interface{}, options *types.Options) (*types.KuzzleSearchResult, error) {
  ch := make(chan types.KuzzleResponse)

  go dc.kuzzle.Query(buildQuery(dc.collection, dc.index, "document", "search", filters), options, ch)

  res := <-ch

  if res.Error.Message != "" {
    return nil, errors.New(res.Error.Message)
  }

  searchResult := &types.KuzzleSearchResult{}
  json.Unmarshal(res.Result, &searchResult)

  return searchResult, nil
}
