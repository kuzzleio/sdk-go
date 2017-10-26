package security

import "github.com/kuzzleio/sdk-go/types"

func (s *Security) NewProfile(id string, policies []*types.Policy) *Profile {
	return &Profile{
		Id:       id,
		Policies: policies,
		Security: s,
	}
}
