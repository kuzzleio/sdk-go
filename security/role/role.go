package role

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
)

type SecurityRole struct {
	Kuzzle kuzzle.Kuzzle
}

/*
  Retrieves a Role using its provided unique id.
*/
func (sr SecurityRole) Fetch(id string, options types.QueryOptions) (types.Role, error) {
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
func (sr SecurityRole) Search(filters interface{}, options types.QueryOptions) (types.KuzzleSearchRolesResult, error) {
	ch := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Controller: "security",
		Action:     "searchRoles",
		Body:       filters,
	}

	if options != nil {
		query.From = options.GetFrom()
		query.Size = options.GetSize()

		scroll := options.GetScroll()
		if scroll != "" {
			query.Scroll = scroll
		}
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
