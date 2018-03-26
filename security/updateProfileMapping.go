package security

import (
	"github.com/kuzzleio/sdk-go/types"
)

func (s *Security) UpdateProfileMapping(body string, options types.QueryOptions) error {
	if body == "" {
		return types.NewError("Security.UpdateProfileMapping: body is required", 400)
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "security",
		Action:     "updateProfileMapping",
		Body:       body,
	}
	go s.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return res.Error
	}

	return nil
}
