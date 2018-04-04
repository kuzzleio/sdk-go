package security

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

// CreateUser creates or replaces (if _id matches an existing one) a User with a list of policies.
func (s *Security) CreateUser(body json.RawMessage, options types.QueryOptions) (json.RawMessage, error) {
	if body == nil {
		return nil, types.NewError("Kuzzle.CreateUser: body is required", 400)
	}
	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "security",
		Action:     "createUser",
		Body:       body,
	}

	go s.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return nil, res.Error
	}

	return res.Result, nil
}
