package security

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

// GetUserRights gets rights for given profile id
func (s *Security) GetUserRights(id string, options types.QueryOptions) (json.RawMessage, error) {
	if id == "" {
		return nil, types.NewError("Security.GetUserRights: id is required", 400)
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "security",
		Action:     "getUserRights",
		Id:         id,
	}
	go s.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return nil, res.Error
	}

	return res.Result, nil
}
