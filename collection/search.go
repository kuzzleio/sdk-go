package collection

import (
  "errors"
  "encoding/json"
  "github.com/kuzzleio/sdk-go/types"
)

/*
  Searches documents in the given Collection, using provided filters and options.
*/
func (dc Collection) Search(filters interface{}, options *types.Options) (*types.KuzzleSearchResult, error) {
  ch := make(chan types.KuzzleResponse)

  query := types.KuzzleRequest{
    Collection: dc.collection,
    Index:      dc.index,
    Controller: "document",
    Action:     "search",
    Body:       filters,
  }
  go dc.kuzzle.Query(query, options, ch)

  res := <-ch

  if res.Error.Message != "" {
    return nil, errors.New(res.Error.Message)
  }

  searchResult := &types.KuzzleSearchResult{}
  json.Unmarshal(res.Result, &searchResult)

  return searchResult, nil
}
