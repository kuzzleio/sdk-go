package security

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

func (s *Security) GetUser(id string, options types.QueryOptions) (*User, error) {
	if id == "" {
		return nil, types.NewError("Security.GetUser: id is required", 400)
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "security",
		Action:     "getUser",
		Id:         id,
	}
	go s.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return nil, res.Error
	}

	jsonUser := &jsonUser{}
	json.Unmarshal(res.Result, jsonUser)

	return s.jsonUserToUser(jsonUser), nil
}
