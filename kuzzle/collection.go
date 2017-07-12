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
func (dc *Collection) Count(filters interface{}) (*int, error) {
  type countResult struct {
    Count int `json:"count"`
  }

  ch := make(chan types.KuzzleResponse)

  go dc.kuzzle.Query(buildQueryArgs(dc.collection, dc.index, "document", "count", filters), ch, nil)

  res := <-ch

  if res.Error.Message != "" {
    return nil, errors.New(res.Error.Message)
  }
  result := &countResult{}
  json.Unmarshal(res.Result, result)

  return &result.Count, nil
}
