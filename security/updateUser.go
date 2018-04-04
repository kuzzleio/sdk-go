package security

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

func (s *Security) UpdateUser(id string, body json.RawMessage, options types.QueryOptions) (*User, error) {
	if id == "" || body == nil {
		return nil, types.NewError("Security.UpdateUser: id and body are required", 400)
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "security",
		Action:     "updateUser",
		Id:         id,
		Body:       body,
	}
	go s.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return nil, res.Error
	}

	var rawUpdated *jsonUser
	var updated *User
	json.Unmarshal(res.Result, &rawUpdated)
	updated = rawUpdated.jsonUserToUser()

	return updated, nil
}
