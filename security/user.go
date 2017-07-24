package security

import (
  "errors"
  "encoding/json"
  "github.com/kuzzleio/sdk-go/types"
)

/*
  Retrieves an User using its provided unique id.
*/
func (security Security) FetchUser(id string, options *types.Options) (types.User, error) {
  if id == "" {
    return types.User{}, errors.New("Security.FetchUser: user id required")
  }

  ch := make(chan types.KuzzleResponse)

  query := types.KuzzleRequest{
    Controller: "security",
    Action:     "getUser",
    Id:         id,
  }
  go security.kuzzle.Query(query, options, ch)

  res := <-ch

  if res.Error.Message != "" {
    return types.User{}, errors.New(res.Error.Message)
  }

  user := types.User{}
  json.Unmarshal(res.Result, &user)

  return user, nil
}
