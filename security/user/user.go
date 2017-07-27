package user

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
)

type SecurityUser struct {
	Kuzzle kuzzle.Kuzzle
}

/*
  Retrieves an User using its provided unique id.
*/
func (su SecurityUser) Fetch(id string, options types.QueryOptions) (types.User, error) {
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
  Executes a search on Users according to filters.
*/
func (su SecurityUser) Search(filters interface{}, options types.QueryOptions) (types.KuzzleSearchUsersResult, error) {
	ch := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Controller: "security",
		Action:     "searchUsers",
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

	go su.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Message != "" {
		return types.KuzzleSearchUsersResult{}, errors.New(res.Error.Message)
	}

	searchResult := types.KuzzleSearchUsersResult{}
	json.Unmarshal(res.Result, &searchResult)

	return searchResult, nil
}

/*
  Executes a scroll search on Users.
*/
func (su SecurityUser) Scroll(scrollId string, options types.QueryOptions) (types.KuzzleSearchUsersResult, error) {
	if scrollId == "" {
		return types.KuzzleSearchUsersResult{}, errors.New("Security.User.Scroll: scroll id required")
	}

	ch := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Controller: "security",
		Action:     "scrollUsers",
		ScrollId:   scrollId,
	}

	go su.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Message != "" {
		return types.KuzzleSearchUsersResult{}, errors.New(res.Error.Message)
	}

	searchResult := types.KuzzleSearchUsersResult{}
	json.Unmarshal(res.Result, &searchResult)

	return searchResult, nil
}

/*
  Gets the rights of an User using its provided unique id.
*/
func (su SecurityUser) GetRights(kuid string, options types.QueryOptions) ([]types.UserRights, error) {
	if kuid == "" {
		return []types.UserRights{}, errors.New("Security.User.GetRights: user id required")
	}

	ch := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Controller: "security",
		Action:     "getUserRights",
		Id:         kuid,
	}
	go su.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Message != "" {
		return []types.UserRights{}, errors.New(res.Error.Message)
	}

	type response struct {
		UserRights []types.UserRights `json:"hits"`
	}
	userRights := response{}
	json.Unmarshal(res.Result, &userRights)

	return userRights.UserRights, nil
}
