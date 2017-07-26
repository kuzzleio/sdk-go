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
  Executes a search on Users according to filters.
*/
func (su SecurityUser) Search(filters interface{}, options *types.Options) (types.KuzzleSearchUsersResult, error) {
	ch := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Controller: "security",
		Action:     "searchUsers",
		Body:       filters,
	}

	if options != nil {
		query.From = options.From
		query.Size = options.Size
		if options.Scroll != "" {
			query.Scroll = options.Scroll
		}
	} else {
		query.Size = 10
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
func (su SecurityUser) Scroll(scrollId string, options *types.Options) (types.KuzzleSearchUsersResult, error) {
	if scrollId == "" {
		return types.KuzzleSearchUsersResult{}, errors.New("Security.User.Scroll: scroll id required")
	}

	ch := make(chan types.KuzzleResponse)

	if options == nil {
		options = &types.Options{Scroll: "1m"}
	}

	query := types.KuzzleRequest{
		Controller: "security",
		Action:     "scrollUsers",
		Scroll:     options.Scroll,
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
