package security

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

func (s *Security) WhoAmI() (*User, error) {
	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "auth",
		Action:     "getCurrentUser",
	}

	go s.Kuzzle.Query(query, nil, result)

	res := <-result

	if res.Error != nil {
		return nil, res.Error
	}

	jsonUser := &jsonUser{}
	json.Unmarshal(res.Result, jsonUser)

	return s.jsonUserToUser(jsonUser), nil
}
