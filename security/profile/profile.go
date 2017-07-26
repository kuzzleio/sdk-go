package profile

import (
  "errors"
  "encoding/json"
  "github.com/kuzzleio/sdk-go/types"
  "github.com/kuzzleio/sdk-go/kuzzle"
)

type SecurityProfile struct {
  Kuzzle kuzzle.Kuzzle
}

/*
  Retrieves a Profile using its provided unique id.
*/
func (sp SecurityProfile) Fetch(id string, options *types.Options) (types.Profile, error) {
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
func (sp SecurityProfile) Search(filters interface{}, options *types.Options) (types.KuzzleSearchProfilesResult, error) {
	ch := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Controller: "security",
		Action:     "searchProfiles",
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
func (sp SecurityProfile) Scroll(scrollId string, options *types.Options) (types.KuzzleSearchProfilesResult, error) {
	if scrollId == "" {
		return types.KuzzleSearchProfilesResult{}, errors.New("Security.Profile.Scroll: scroll id required")
	}

	ch := make(chan types.KuzzleResponse)

	if options == nil {
		options = &types.Options{Scroll: "1m"}
	}

	query := types.KuzzleRequest{
		Controller: "security",
		Action:     "scrollProfiles",
		Scroll:     options.Scroll,
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
