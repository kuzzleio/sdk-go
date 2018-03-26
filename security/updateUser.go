package security

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

func (s *Security) UpdateUser(id, body string, options types.QueryOptions) (*User, error) {
	if id == "" || body == "" {
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

	var updated *User
	json.Unmarshal(res.Result, &updated)

	return updated, nil
}
