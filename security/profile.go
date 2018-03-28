package security

import (
	"github.com/kuzzleio/sdk-go/types"
)

type Profile struct {
	Id       string `json:"_id"`
	Policies []*types.Policy
}

type ProfileSearchResult struct {
	Hits     []*Profile
	Total    int    `json:"total"`
	ScrollId string `json:"scrollId"`
}

func NewProfile(id string, policies []*types.Policy) *Profile {
	return &Profile{
		Id:       id,
		Policies: policies,
	}
}
