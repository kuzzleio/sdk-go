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

/*
  Executes a search on Roles according to filters.
*/
func (sr SecurityRole) Search(filters interface{}, options *types.Options) (types.KuzzleSearchRolesResult, error) {
	ch := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Controller: "security",
		Action:     "searchRoles",
		Body:       filters,
	}

	if options != nil {
		query.From = options.From
		query.Size = options.Size
	} else {
		query.Size = 10
	}

	go sr.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Message != "" {
		return types.KuzzleSearchRolesResult{}, errors.New(res.Error.Message)
	}

	searchResult := types.KuzzleSearchRolesResult{}
	json.Unmarshal(res.Result, &searchResult)

	return searchResult, nil
}
