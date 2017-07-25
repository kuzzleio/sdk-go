package user

import (
  "errors"
  "encoding/json"
  "github.com/kuzzleio/sdk-go/types"
  "github.com/kuzzleio/sdk-go/kuzzle"
)

type SecurityUser struct {
  Kuzzle kuzzle.Kuzzle
}

/*
  Retrieves an User using its provided unique id.
*/
func (su SecurityUser) Fetch(id string, options *types.Options) (types.User, error) {
  if id == "" {
    return types.User{}, errors.New("Security.User.Fetch: user id required")
  }

  ch := make(chan types.KuzzleResponse)

  query := types.KuzzleRequest{
    Controller: "security",
    Action:     "getUser",
    Id:         id,
  }
  go su.Kuzzle.Query(query, options, ch)

  res := <-ch

  if res.Error.Message != "" {
    return types.User{}, errors.New(res.Error.Message)
  }

  user := types.User{}
  json.Unmarshal(res.Result, &user)

  return user, nil
}
