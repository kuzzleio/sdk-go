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

/*
  Create a new User in Kuzzle.
*/
func (su SecurityUser) Create(id string, content types.UserData, options *types.Options) (types.User, error) {
  if id == "" {
    return types.User{}, errors.New("Security.User.Create: user id required")
  }

  ch := make(chan types.KuzzleResponse)

  type userData map[string]interface {}
  ud := userData{}
  ud["profileIds"] = content.ProfileIds
  for key, value := range content.Content {
    ud[key] = value
  }
  type createBody struct {
    Content     userData              `json:"content"`
    Credentials types.UserCredentials `json:"credentials"`
  }

  body := createBody{Content: ud, Credentials: content.Credentials}

  query := types.KuzzleRequest{
    Controller: "security",
    Action:     "createUser",
    Body:       body,
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

/*
  Create a new restricted User in Kuzzle.
*/
func (su SecurityUser) CreateRestrictedUser(id string, content types.UserData, options *types.Options) (types.User, error) {
  if id == "" {
    return types.User{}, errors.New("Security.User.CreateRestrictedUser: user id required")
  }

  ch := make(chan types.KuzzleResponse)

  type userData map[string]interface {}
  ud := userData{}
  for key, value := range content.Content {
    ud[key] = value
  }
  type createBody struct {
    Content     userData              `json:"content"`
    Credentials types.UserCredentials `json:"credentials"`
  }

  body := createBody{Content: ud, Credentials: content.Credentials}

  query := types.KuzzleRequest{
    Controller: "security",
    Action:     "createRestrictedUser",
    Body:       body,
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

/*
  Replace an User in Kuzzle.
*/
func (su SecurityUser) Replace(id string, content types.UserData, options *types.Options) (types.User, error) {
  if id == "" {
    return types.User{}, errors.New("Security.User.Replace: user id required")
  }

  ch := make(chan types.KuzzleResponse)

  type userData map[string]interface {}
  ud := userData{}
  ud["profileIds"] = content.ProfileIds
  for key, value := range content.Content {
    ud[key] = value
  }

  query := types.KuzzleRequest{
    Controller: "security",
    Action:     "replaceUser",
    Body:       ud,
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

/*
  Update an User in Kuzzle.
*/
func (su SecurityUser) Update(id string, content types.UserData, options *types.Options) (types.User, error) {
  if id == "" {
    return types.User{}, errors.New("Security.User.Update: user id required")
  }

  ch := make(chan types.KuzzleResponse)

  type userData map[string]interface {}
  ud := userData{}
  ud["profileIds"] = content.ProfileIds
  for key, value := range content.Content {
    ud[key] = value
  }

  query := types.KuzzleRequest{
    Controller: "security",
    Action:     "updateUser",
    Body:       ud,
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

/*
  Delete an User in Kuzzle.

  There is a small delay between user deletion and their deletion in our advanced search layer, usually a couple of seconds.
  This means that a user that has just been deleted will still be returned by this function.
*/
func (su SecurityUser) Delete(id string, options *types.Options) (string, error) {
  if id == "" {
    return "", errors.New("Security.User.Delete: user id required")
  }

  ch := make(chan types.KuzzleResponse)

  query := types.KuzzleRequest{
    Controller: "security",
    Action:     "deleteUser",
    Id:         id,
  }
  go su.Kuzzle.Query(query, options, ch)

  res := <-ch

  if res.Error.Message != "" {
    return "", errors.New(res.Error.Message)
  }

  shardResponse := types.ShardResponse{}
  json.Unmarshal(res.Result, &shardResponse)

  return shardResponse.Id, nil
}
