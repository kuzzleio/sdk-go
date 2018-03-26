package security

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

// GetProfile fetch profile matching with given id
func (s *Security) GetProfile(id string, options types.QueryOptions) (*Profile, error) {
	if id == "" {
		return nil, types.NewError("Security.GetProfile: id is required", 400)
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "security",
		Action:     "getProfile",
		Id:         id,
	}
	go s.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return nil, res.Error
	}

	jsonProfile := &jsonProfile{}
	json.Unmarshal(res.Result, jsonProfile)

	return s.jsonProfileToProfile(jsonProfile), nil
}
