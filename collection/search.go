package collection

import (
  "errors"
  "encoding/json"
  "github.com/kuzzleio/sdk-go/internal"
  "github.com/kuzzleio/sdk-go/types"
)

/*
  Searches documents in the given Collection, using provided filters and options.
*/
func (dc Collection) Search(filters interface{}, options *types.Options) (*types.KuzzleSearchResult, error) {
  ch := make(chan types.KuzzleResponse)

  go dc.kuzzle.Query(internal.BuildQuery(dc.collection, dc.index, "document", "search", filters), options, ch)

  res := <-ch

  if res.Error.Message != "" {
    return nil, errors.New(res.Error.Message)
  }

  searchResult := &types.KuzzleSearchResult{}
  json.Unmarshal(res.Result, &searchResult)

  return searchResult, nil
}
