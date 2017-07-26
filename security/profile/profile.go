package profile

import (
	"errors"
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"fmt"
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
  Create a new Profile in Kuzzle.
*/
func (sp SecurityProfile) Create(id string, policies types.Policies, options types.QueryOptions) (types.Profile, error) {
	if id == "" {
		return types.Profile{}, errors.New("Security.Profile.Create: profile id required")
	}

	action := "createProfile"

	if options != nil {
		if options.GetIfExist() == "replace" {
			action = "createOrReplaceProfile"
		} else if options.GetIfExist() != "error" {
			return types.Profile{}, errors.New(fmt.Sprintf("Invalid value for the 'ifExist' option: '%s'", options.GetIfExist()))
		}
	}

	ch := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Controller: "security",
		Action:     action,
		Body:       policies,
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
  Update a Profile in Kuzzle.
*/
func (sp SecurityProfile) Update(id string, policies types.Policies, options types.QueryOptions) (types.Profile, error) {
	if id == "" {
		return types.Profile{}, errors.New("Security.Profile.Update: profile id required")
	}

	ch := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Controller: "security",
		Action:     "updateProfile",
		Body:       policies,
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
 * Delete a Profile in Kuzzle.
 *
 * There is a small delay between profile deletion and their deletion in our advanced search layer, usually a couple of seconds.
 * This means that a profile that has just been deleted will still be returned by this function.
*/
func (sp SecurityProfile) Delete(id string, options types.QueryOptions) (string, error) {
	if id == "" {
		return "", errors.New("Security.Profile.Delete: profile id required")
	}

	ch := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Controller: "security",
		Action:     "deleteProfile",
		Id:         id,
	}
	go sp.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Message != "" {
		return "", errors.New(res.Error.Message)
	}

	shardResponse := types.ShardResponse{}
	json.Unmarshal(res.Result, &shardResponse)

	return shardResponse.Id, nil
}
