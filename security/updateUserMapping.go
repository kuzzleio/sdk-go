package security

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

func (s *Security) UpdateUserMapping(body json.RawMessage, options types.QueryOptions) error {
	if body == nil {
		return types.NewError("Security.UpdateUserMapping: body is required", 400)
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "security",
		Action:     "updateUserMapping",
		Body:       body,
	}
	go s.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return res.Error
	}

	return nil
}
