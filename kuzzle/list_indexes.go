package kuzzle

import (
  "github.com/kuzzleio/sdk-go/types"
  "errors"
  "encoding/json"
)

/*
 * List data indexes
 */
func (k Kuzzle) ListIndexes(options *types.Options) ([]string, error) {
  result := make(chan types.KuzzleResponse)

  query := types.KuzzleRequest{
    Controller: "index",
    Action:     "list",
  }

  go k.Query(query, nil, result)

  res := <-result

  type indexes struct {
    Indexes []string  `json:"indexes"`
  }

  if res.Error.Message != "" {
    return nil, errors.New(res.Error.Message)
  }

  var i indexes
  json.Unmarshal(res.Result, &i)

  return i.Indexes, nil
}
