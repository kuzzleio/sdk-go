package security

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

// DeleteProfile deletes (if _id matches an existing one) a Profile with a list of policies.
func (s *Security) DeleteProfile(id string, options types.QueryOptions) (string, error) {
	if id == "" {
		return "", types.NewError("Kuzzle.DeleteProfile: id is required", 400)
	}
	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "security",
		Action:     "deleteProfile",
		Id:         id,
	}

	go s.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return "", res.Error
	}

	var deletedProfileID struct {
		ID string `json:"_id"`
	}

	json.Unmarshal(res.Result, &deletedProfileID)

	return deletedProfileID.ID, nil
}
