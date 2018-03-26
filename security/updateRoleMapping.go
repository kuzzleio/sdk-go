package security

import (
	"github.com/kuzzleio/sdk-go/types"
)

func (s *Security) UpdateRoleMapping(body string, options types.QueryOptions) error {
	if body == "" {
		return types.NewError("Security.UpdateRoleMapping: body is required", 400)
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "security",
		Action:     "updateRoleMapping",
		Body:       body,
	}
	go s.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return res.Error
	}

	return nil
}
