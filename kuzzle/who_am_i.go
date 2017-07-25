package kuzzle

import (
  "github.com/kuzzleio/sdk-go/types"
  "errors"
  "encoding/json"
)

/*
 * Gets the rights array for the currently logged user.
 */
func (k Kuzzle) WhoAmI() (*types.User, error) {
  result := make(chan types.KuzzleResponse)

  query := types.KuzzleRequest{
    Controller: "auth",
    Action:     "getCurrentUser",
  }

  go k.Query(query, nil, result)

  res := <-result

  if res.Error.Message != "" {
    return nil, errors.New(res.Error.Message)
  }

  u := types.User{}
  json.Unmarshal(res.Result, &u)

  return &u, nil
}
