package profile

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
)

type SecurityProfile struct {
	Kuzzle kuzzle.Kuzzle
}

type Profile struct {
	Id     string           `json:"_id"`
	Source json.RawMessage  `json:"_source"`
	Meta   types.KuzzleMeta `json:"_meta"`
	SP     SecurityProfile
}

type ProfileSearchResult struct {
	Hits     []Profile `json:"hits"`
	Total    int       `json:"total"`
	ScrollId string    `json:"scrollId"`
}

/*
  Adds a role to the profile.
*/
func (p *Profile) AddPolicy(policy types.Policy) Profile {
	return p.SetPolicies(append(p.GetPolicies(), policy))
}

/*
  Returns this profile associated role policies.
 */
func (p Profile) GetPolicies() []types.Policy {
	policies := types.Policies{}
	json.Unmarshal(p.Source, &policies)

	return policies.Policies
}

/*
  Replaces the roles policies associated to the profile.
 */
func (p *Profile) SetPolicies(policies []types.Policy) Profile {
	content := map[string]interface{}{}
	json.Unmarshal(p.Source, content)

	content["policies"] = policies

	p.Source, _ = json.Marshal(content)

	return *p
}

/*
  Replaces the content of the Profile object.
 */
func (p *Profile) SetContent(data json.RawMessage) Profile {
	p.Source = data

	p.SetPolicies(p.GetPolicies())

	return *p
}

/*
  Creates or replaces the profile in Kuzzle.
 */
func (p Profile) Save(options types.QueryOptions) (Profile, error) {
	if options == nil {
		options = types.NewQueryOptions()
	}

	return p.SP.Create(p.Id, types.Policies{Policies: p.GetPolicies()}, options.SetIfExist("replace"))
}

/*
  Performs a partial content update on this object.
*/
func (p Profile) Update(policies []types.Policy, options types.QueryOptions) (Profile, error) {
	return p.SP.Update(p.Id, types.Policies{Policies: policies}, options)
}

/*
  Deletes this profile from Kuzzle.
 */
func (p Profile) Delete(options types.QueryOptions) (string, error) {
	return p.SP.Delete(p.Id, options)
}

/*
  Retrieves a Profile using its provided unique id.
*/
func (sp SecurityProfile) Fetch(id string, options types.QueryOptions) (Profile, error) {
	if id == "" {
		return Profile{}, errors.New("Security.Profile.Fetch: profile id required")
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
		return Profile{}, errors.New(res.Error.Message)
	}

	profile := Profile{SP: sp}
	json.Unmarshal(res.Result, &profile)

	return profile, nil
}

/*
  Executes a search on Profiles according to filters.
*/
func (sp SecurityProfile) Search(filters interface{}, options types.QueryOptions) (ProfileSearchResult, error) {
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
		return ProfileSearchResult{}, errors.New(res.Error.Message)
	}

	searchResult := ProfileSearchResult{}
	json.Unmarshal(res.Result, &searchResult)

	return searchResult, nil
}

/*
  Executes a scroll search on Profiles.
*/
func (sp SecurityProfile) Scroll(scrollId string, options types.QueryOptions) (ProfileSearchResult, error) {
	if scrollId == "" {
		return ProfileSearchResult{}, errors.New("Security.Profile.Scroll: scroll id required")
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
		return ProfileSearchResult{}, errors.New(res.Error.Message)
	}

	searchResult := ProfileSearchResult{}
	json.Unmarshal(res.Result, &searchResult)

	return searchResult, nil
}

/*
  Create a new Profile in Kuzzle.
*/
func (sp SecurityProfile) Create(id string, policies types.Policies, options types.QueryOptions) (Profile, error) {
	if id == "" {
		return Profile{}, errors.New("Security.Profile.Create: profile id required")
	}

	action := "createProfile"

	if options != nil {
		if options.GetIfExist() == "replace" {
			action = "createOrReplaceProfile"
		} else if options.GetIfExist() != "error" {
			return Profile{}, errors.New(fmt.Sprintf("Invalid value for the 'ifExist' option: '%s'", options.GetIfExist()))
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
		return Profile{}, errors.New(res.Error.Message)
	}

	profile := Profile{SP: sp}
	json.Unmarshal(res.Result, &profile)

	return profile, nil
}

/*
  Update a Profile in Kuzzle.
*/
func (sp SecurityProfile) Update(id string, policies types.Policies, options types.QueryOptions) (Profile, error) {
	if id == "" {
		return Profile{}, errors.New("Security.Profile.Update: profile id required")
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
		return Profile{}, errors.New(res.Error.Message)
	}

	profile := Profile{SP: sp}
	json.Unmarshal(res.Result, &profile)

	return profile, nil
}

/*
  Delete a Profile in Kuzzle.
  There is a small delay between profile deletion and their deletion in our advanced search layer, usually a couple of seconds.
  This means that a profile that has just been deleted will still be returned by this function.
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
