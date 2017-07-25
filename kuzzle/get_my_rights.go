package kuzzle

import (
  "github.com/kuzzleio/sdk-go/types"
  "errors"
  "encoding/json"
)

/*
 * Gets the rights array for the currently logged user.
 */
func (k Kuzzle) GetMyRights(options *types.Options) ([]types.Rights, error) {
  result := make(chan types.KuzzleResponse)

  query := types.KuzzleRequest{
    Controller: "auth",
    Action:     "getMyRights",
  }

  type rights struct {
    Hits []types.Rights `json:"hits"`
  }

  go k.Query(query, nil, result)

  res := <-result

  if res.Error.Message != "" {
    return nil, errors.New(res.Error.Message)
  }

  r := rights{}
  json.Unmarshal(res.Result, &r)

  return r.Hits, nil
}
