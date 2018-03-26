package security

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

// GetProfileRights gets rights for given profile id
func (s *Security) GetProfileRights(id string, options types.QueryOptions) (json.RawMessage, error) {
	if id == "" {
		return nil, types.NewError("Security.GetProfileRights: id is required", 400)
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "security",
		Action:     "getProfileRights",
		Id:         id,
	}
	go s.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return nil, res.Error
	}

	return res.Result, nil
}
