package collection

import (
  "github.com/kuzzleio/sdk-go/internal"
  "errors"
  "encoding/json"
  "github.com/kuzzleio/sdk-go/types"
)

/*
  Returns the number of documents matching the provided set of filters.

  There is a small delay between documents creation and their existence in our advanced search layer,
  usually a couple of seconds.
  That means that a document that was just been created won’t be returned by this function
*/
func (dc *Collection) Count(filters interface{}, options *types.Options) (*int, error) {
  type countResult struct {
    Count int `json:"count"`
  }

  ch := make(chan types.KuzzleResponse)

  go dc.kuzzle.Query(internal.BuildQuery(dc.collection, dc.index, "document", "count", filters), options, ch)

  res := <-ch

  if res.Error.Message != "" {
    return nil, errors.New(res.Error.Message)
  }
  result := &countResult{}
  json.Unmarshal(res.Result, result)

  return &result.Count, nil
}