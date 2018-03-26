package security

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

// DeleteRole creates or replaces (if _id matches an existing one) a Profile with a list of policies.
func (s *Security) DeleteRole(id string, options types.QueryOptions) (string, error) {
	if id == "" {
		return "", types.NewError("Kuzzle.DeleteRole: id is required", 400)
	}
	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "security",
		Action:     "deleteRole",
		Id:         id,
	}

	go s.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return "", res.Error
	}

	var deletedRole struct {
		ID string `json:"_id"`
	}

	json.Unmarshal(res.Result, &deletedRole)

	return deletedRole.ID, nil
}
