package kuzzle

import (
  "github.com/kuzzleio/sdk-go/types"
  "errors"
  "encoding/json"
)

/*
 * Update the currently authenticated user informations
 */
func (k Kuzzle) UpdateSelf(credentials interface{}, options *types.Options) (types.User, error) {
  result := make(chan types.KuzzleResponse)

  query := types.KuzzleRequest{
    Controller: "auth",
    Action:     "updateSelf",
    Body:       credentials,
  }

  go k.Query(query, nil, result)

  res := <-result

  if res.Error.Message != "" {
    return types.User{}, errors.New(res.Error.Message)
  }

  u := types.User{}
  json.Unmarshal(res.Result, &u)

  return u, nil
}
