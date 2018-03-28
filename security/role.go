package security

import (
	"github.com/kuzzleio/sdk-go/types"
)

type Role struct {
	Id          string `json:"_id"`
	Controllers map[string]*types.Controller
}

type RoleSearchResult struct {
	Hits  []*Role
	Total int
}

func NewRole(id string, controllers *types.Controllers) *Role {
	r := &Role{
		Id: id,
	}
	if controllers != nil {
		r.Controllers = controllers.Controllers
	}

	return r
}
