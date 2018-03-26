package security

import (
	"github.com/kuzzleio/sdk-go/types"
)

type Role struct {
	Id          string `json:"_id"`
	Controllers map[string]*types.Controller
	Security    *Security
}

type RoleSearchResult struct {
	Hits  []*Role
	Total int
}

func (s *Security) NewRole(id string, controllers *types.Controllers) *Role {
	r := &Role{
		Id:       id,
		Security: s,
	}
	if controllers != nil {
		r.Controllers = controllers.Controllers
	}

	return r
}
