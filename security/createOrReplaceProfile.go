package security

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

// CreateOrReplaceProfile creates or replaces (if _id matches an existing one) a profile with a list of policies.
func (s *Security) CreateOrReplaceProfile(id, body string, options types.QueryOptions) (*Profile, error) {
	if body == "" {
		return nil, types.NewError("Kuzzle.CreateOrReplaceProfile: body is required", 400)
	}
	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "security",
		Action:     "createOrReplaceProfile",
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
