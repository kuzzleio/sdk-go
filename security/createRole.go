package security

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

// CreateRole creates or replaces (if _id matches an existing one) a Role with a list of policies.
func (s *Security) CreateRole(id, body string, options types.QueryOptions) (*Role, error) {
	if id == "" || body == "" {
		return nil, types.NewError("Kuzzle.CreateRole: id and body are required", 400)
	}
	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "security",
		Action:     "createRole",
		Id:         id,
		Body:       body,
	}

	go s.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return nil, res.Error
	}

	var role *Role

	json.Unmarshal(res.Result, &role)

	return role, nil
}
