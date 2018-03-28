package security

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

// MGetRoles gets all roles matching with given ids
func (s *Security) MGetRoles(ids []string, options types.QueryOptions) ([]*Role, error) {
	if len(ids) == 0 {
		return nil, types.NewError("Security.MGetRoles: ids array can't be nil", 400)
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "security",
		Action:     "mGetRoles",
		Body: struct {
			Ids []string `json:"ids"`
		}{ids},
	}
	go s.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return nil, res.Error
	}

	var fetchedRaw jsonRoleSearchResult
	var fetchedRoles []*Role
	json.Unmarshal(res.Result, &fetchedRaw)

	for _, jsonRoleRaw := range fetchedRaw.Hits {
		fetchedRoles = append(fetchedRoles, jsonRoleRaw.jsonRoleToRole())
	}

	return fetchedRoles, nil

}
