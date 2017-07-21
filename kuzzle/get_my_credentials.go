package kuzzle

import (
  "github.com/kuzzleio/sdk-go/types"
  "errors"
  "encoding/json"
)

/*
 * Get all Kuzzle usage statistics frames
 */
func (k Kuzzle) GetMyCredentials(strategy string, options *types.Options) (json.RawMessage, error) {
  result := make(chan types.KuzzleResponse)

  query := types.KuzzleRequest{
    Controller: "auth",
    Action:     "getMyCredentials",
    Strategy:   strategy,
  }

  go k.Query(query, nil, result)

  res := <-result

  if res.Error.Message != "" {
    return nil, errors.New(res.Error.Message)
  }

  return res.Result, nil
}
