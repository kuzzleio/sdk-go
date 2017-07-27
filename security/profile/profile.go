package profile

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
)

type SecurityProfile struct {
	Kuzzle kuzzle.Kuzzle
}

/*
  Retrieves a Profile using its provided unique id.
*/
func (sp SecurityProfile) Fetch(id string, options types.QueryOptions) (types.Profile, error) {
	if id == "" {
		return types.Profile{}, errors.New("Security.Profile.Fetch: profile id required")
	}

	ch := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Controller: "security",
		Action:     "getProfile",
		Id:         id,
	}
	go sp.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Message != "" {
		return types.Profile{}, errors.New(res.Error.Message)
	}

	profile := types.Profile{}
	json.Unmarshal(res.Result, &profile)

	return profile, nil
}

/*
  Executes a search on Profiles according to filters.
*/
func (sp SecurityProfile) Search(filters interface{}, options types.QueryOptions) (types.KuzzleSearchProfilesResult, error) {
	ch := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Controller: "security",
		Action:     "searchProfiles",
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

	go sp.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Message != "" {
		return types.KuzzleSearchProfilesResult{}, errors.New(res.Error.Message)
	}

	searchResult := types.KuzzleSearchProfilesResult{}
	json.Unmarshal(res.Result, &searchResult)

	return searchResult, nil
}

/*
  Executes a scroll search on Profiles.
*/
func (sp SecurityProfile) Scroll(scrollId string, options types.QueryOptions) (types.KuzzleSearchProfilesResult, error) {
	if scrollId == "" {
		return types.KuzzleSearchProfilesResult{}, errors.New("Security.Profile.Scroll: scroll id required")
	}

	ch := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Controller: "security",
		Action:     "scrollProfiles",
		ScrollId:   scrollId,
	}

	go sp.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Message != "" {
		return types.KuzzleSearchProfilesResult{}, errors.New(res.Error.Message)
	}

	searchResult := types.KuzzleSearchProfilesResult{}
	json.Unmarshal(res.Result, &searchResult)

	return searchResult, nil
}
