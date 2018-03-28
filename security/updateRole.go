package security

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

func (s *Security) UpdateRole(id string, body json.RawMessage, options types.QueryOptions) (*Role, error) {
	if id == "" || body == nil {
		return nil, types.NewError("Security.UpdateRole: id and body are required", 400)
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "security",
		Action:     "updateRole",
		Id:         id,
		Body:       body,
	}
	go s.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return nil, res.Error
	}

	var updated *Role
	json.Unmarshal(res.Result, &updated)

	return updated, nil
}
