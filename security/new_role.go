package security

import "github.com/kuzzleio/sdk-go/types"

func (s *Security) NewRole(id string, controllers *types.Controllers) *Role {
	r := &Role{
		Id:          id,
		Security:    s,
	}
	if controllers != nil {
		r.Controllers = controllers.Controllers
	}

	return r
}
