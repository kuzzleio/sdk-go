package security

import (
	"github.com/kuzzleio/sdk-go/types"
	"encoding/json"
)

type jsonRole struct {
	Id string `json:"_id"`
	Source types.Controllers `json:"_source"`
}

type jsonSearchResult struct {
	Hits []*jsonRole `json:"hits"`
	Total int        `json:"total"`
}

func (s *Security) jsonRoleToRole(j *jsonRole) *Role {
	r := &Role{}
	r.Id = j.Id
	r.Controllers = j.Source.Controllers
	r.Kuzzle = s.Kuzzle

	return r
}

func RoleToJson(r *Role) ([]byte, error) {
	j := &jsonRole{
		Id: r.Id,
		Source: types.Controllers{
			Controllers: r.Controllers,
		},
	}
	return json.Marshal(j)
}
