package security

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

// CreateOrReplaceRole creates or replaces (if _id matches an existing one) a Role with a list of policies.
func (s *Security) CreateOrReplaceRole(id string, body json.RawMessage, options types.QueryOptions) (*Role, error) {
	if body == nil {
		return nil, types.NewError("Kuzzle.CreateOrReplaceRole: body is required", 400)
	}
	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "security",
		Action:     "createOrReplaceRole",
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
