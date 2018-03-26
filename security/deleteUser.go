package security

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

// DeleteUser deletes (if _id matches an existing one) a Profile with a list of policies.
func (s *Security) DeleteUser(id string, options types.QueryOptions) (string, error) {
	if id == "" {
		return "", types.NewError("Kuzzle.DeleteUser: id is required", 400)
	}
	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "security",
		Action:     "deleteUser",
		Id:         id,
	}

	go s.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return "", res.Error
	}

	var deletedUser struct {
		ID string `json:"_id"`
	}

	json.Unmarshal(res.Result, &deletedUser)

	return deletedUser.ID, nil
}
