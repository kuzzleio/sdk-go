package security

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

// GetRoleMapping gets mapping for Roles
func (s *Security) GetRoleMapping(options types.QueryOptions) (json.RawMessage, error) {
	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "security",
		Action:     "getRoleMapping",
	}
	go s.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return nil, res.Error
	}

	return res.Result, nil
}
