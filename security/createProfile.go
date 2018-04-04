package security

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

// CreateProfile creates or replaces (if _id matches an existing one) a profile with a list of policies.
func (s *Security) CreateProfile(id string, body json.RawMessage, options types.QueryOptions) (*Profile, error) {
	if id == "" || body == nil {
		return nil, types.NewError("Kuzzle.CreateProfile: id and body are required", 400)
	}
	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "security",
		Action:     "createProfile",
		Id:         id,
		Body:       body,
	}

	go s.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return nil, res.Error
	}

	var profile *Profile

	json.Unmarshal(res.Result, &profile)

	return profile, nil
}
