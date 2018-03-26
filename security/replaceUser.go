package security

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

// ReplaceUser replaces the matching user with the given one
func (s *Security) ReplaceUser(id, content string, options types.QueryOptions) (*User, error) {
	if id == "" || content == "" {
		return nil, types.NewError("Security.ReplaceUser: id and content are required", 400)
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "security",
		Action:     "replaceUser",
		Id:         id,
		Body:       content,
	}
	go s.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return nil, res.Error
	}

	var replaced *User
	json.Unmarshal(res.Result, &replaced)

	return replaced, nil

}
