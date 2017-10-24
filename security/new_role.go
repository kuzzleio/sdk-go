package security

import "github.com/kuzzleio/sdk-go/types"

func (s *Security) NewRole(id string, controllers map[string]*types.Controller) *Role {
	r := &Role{
		Id:          id,
		Controllers: controllers,
		Security:    s,
	}

	return r
}
