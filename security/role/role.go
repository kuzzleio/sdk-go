package role

import (
  "errors"
  "encoding/json"
  "github.com/kuzzleio/sdk-go/types"
  "github.com/kuzzleio/sdk-go/kuzzle"
)

type SecurityRole struct {
  Kuzzle kuzzle.Kuzzle
}

/*
  Retrieves a Role using its provided unique id.
*/
func (sr SecurityRole) Fetch(id string, options *types.Options) (types.Role, error) {
  if id == "" {
    return types.Role{}, errors.New("Security.Role.Fetch: role id required")
  }

  ch := make(chan types.KuzzleResponse)

  query := types.KuzzleRequest{
    Controller: "security",
    Action:     "getRole",
    Id:         id,
  }
  go sr.Kuzzle.Query(query, options, ch)

  res := <-ch

  if res.Error.Message != "" {
    return types.Role{}, errors.New(res.Error.Message)
  }

  role := types.Role{}
  json.Unmarshal(res.Result, &role)

  return role, nil
}
