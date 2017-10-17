package security

import (
	"github.com/kuzzleio/sdk-go/types"
	"errors"
	"encoding/json"
)

func (s *Security) FetchRole(id string, options types.QueryOptions) (*Role, error) {
	if id == "" {
		return nil, errors.New("Security.Role.Fetch: role id is required")
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "security",
		Action: "getRole",
		Id: id,
	}
	go s.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return nil, errors.New(res.Error.Message)
	}

	jsonRole := &jsonRole{}
	json.Unmarshal(res.Result, jsonRole)

	role := s.jsonRoleToRole(jsonRole)
	role.Kuzzle = s.Kuzzle

	return role, nil
}
