package security

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

// ReplaceUser replaces the matching user with the given one
func (s *Security) ReplaceUser(id string, content json.RawMessage, options types.QueryOptions) (*User, error) {
	if id == "" || content == nil {
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
	var rawReplaced *jsonUser
	json.Unmarshal(res.Result, &rawReplaced)
	replaced = rawReplaced.jsonUserToUser()

	return replaced, nil

}
