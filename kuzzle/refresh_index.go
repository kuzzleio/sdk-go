package kuzzle

import (
  "github.com/kuzzleio/sdk-go/types"
  "errors"
  "encoding/json"
)

/*
 * Forces the provided data index to refresh on each modification
 */
func (k Kuzzle) RefreshIndex(index string, options *types.Options) (types.Shards, error) {
  result := make(chan types.KuzzleResponse)

  query := types.KuzzleRequest{
    Controller: "index",
    Action:     "refresh",
    Index:      index,
  }
  go k.Query(query, options, result)

  res := <-result

  if res.Error.Message != "" {
    return types.Shards{}, errors.New(res.Error.Message)
  }

  type s struct {
    Shards types.Shards `json:"_shards"`
  }

  shards := s{}

  json.Unmarshal(res.Result, &shards)

  return shards.Shards, nil
}
