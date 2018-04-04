package security

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

type jsonRole struct {
	Id     string            `json:"_id"`
	Source types.Controllers `json:"_source"`
}

type jsonRoleSearchResult struct {
	Hits     []*jsonRole `json:"hits"`
	Total    int         `json:"total"`
	ScrollId string      `json:"scrollId"`
}

func (j *jsonRole) jsonRoleToRole() *Role {
	r := &Role{}
	r.Id = j.Id
	r.Controllers = j.Source.Controllers

	return r
}

func (r *Role) RoleToJson() ([]byte, error) {
	j := &jsonRole{
		Id: r.Id,
		Source: types.Controllers{
			Controllers: r.Controllers,
		},
	}
	return json.Marshal(j)
}
