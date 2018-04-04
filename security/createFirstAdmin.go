package security

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

// CreateFirstAdmin create credentials of the specified strategy with given body infos.
func (s *Security) CreateFirstAdmin(body json.RawMessage, options types.QueryOptions) (json.RawMessage, error) {
	if body == nil {
		return nil, types.NewError("Kuzzle.CreateFirstAdmin: strategy, id and body are required", 400)
	}
	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "security",
		Action:     "createFirstAdmin",
		Body:       body,
	}

	if options != nil {
		query.Id = options.ID()
		query.Reset = options.Reset()
	}

	go s.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return nil, res.Error
	}

	return res.Result, nil
}
