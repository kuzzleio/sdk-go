package kuzzle

import (
  "github.com/kuzzleio/sdk-go/types"
  "errors"
  "encoding/json"
)

/*
 * Get Kuzzle usage statistics
 */
func (k Kuzzle) GetStatistics(options *types.Options) (types.Statistics, error) {
  result := make(chan types.KuzzleResponse)

  query := types.KuzzleRequest{
    Controller: "server",
    Action:     "getLastStats",
  }

  go k.Query(query, nil, result)

  res := <-result

  if res.Error.Message != "" {
    return types.Statistics{}, errors.New(res.Error.Message)
  }

  s := types.Statistics{}
  json.Unmarshal(res.Result, &s)

  return s, nil
}
