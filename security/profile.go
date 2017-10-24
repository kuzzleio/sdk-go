package security

import (
	"encoding/json"
	"fmt"
	"github.com/kuzzleio/sdk-go/types"
)

type Profile struct {
	Id       string `json:"_id"`
	Policies []*types.Policy
	Security *Security
}

type ProfileSearchResult struct {
	Hits     []*Profile
	Total    int    `json:"total"`
	ScrollId string `json:"scrollId"`
}

func (p *Profile) AddPolicy(policy *types.Policy) *Profile {
	p.Policies = append(p.Policies, policy)

	return p
}

// Delete deletes the profile form Kuzzle.
func (p *Profile) Delete(options types.QueryOptions) (string, error) {
	return p.Security.rawDelete("deleteProfile", p.Id, options)
}

// Save creates or replaces the profile in Kuzzle.
func (p Profile) Save(options types.QueryOptions) (*Profile, error) {
	action := "createOrReplaceProfile"

	if options == nil && p.Id == "" {
		action = "createProfile"
	}

	if options != nil {
		if options.GetIfExist() == "error" {
			action = "createProfile"
		} else if options.GetIfExist() != "replace" {
			return nil, types.NewError(fmt.Sprintf("Invalid value for 'ifExist' option: '%s'", options.GetIfExist()), 400)
		}
	}

	return p.persist(action, options)
}

// Update performs a partial content update on this object.
func (p *Profile) Update(policies []*types.Policy, options types.QueryOptions) (*Profile, error) {
	p.Policies = policies
	return p.persist("updateProfile", options)
}

func (p *Profile) persist(action string, options types.QueryOptions) (*Profile, error) {
	if options == nil {
		options = types.NewQueryOptions()
	}

	if action != "createProfile" && p.Id == "" {
		return nil, types.NewError("Profile."+action+": id is required", 400)
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "security",
		Action:     action,
		Body: types.Policies{
			Policies: p.Policies,
		},
		Id: p.Id,
	}
	go p.Security.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return nil, res.Error
	}

	jsonProfile := &jsonProfile{}
	json.Unmarshal(res.Result, jsonProfile)

	p.Id = jsonProfile.Id
	p.Policies = jsonProfile.Source.Policies

	return p, nil
}
