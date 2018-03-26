package security

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

func (s *Security) GetRole(id string, options types.QueryOptions) (*Role, error) {
	if id == "" {
		return nil, types.NewError("Security.GetRole: id is required", 400)
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "security",
		Action:     "getRole",
		Id:         id,
	}
	go s.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return nil, res.Error
	}

	jsonRole := &jsonRole{}
	json.Unmarshal(res.Result, jsonRole)

	return s.jsonRoleToRole(jsonRole), nil
}
