package security

import (
	"github.com/kuzzleio/sdk-go/types"
)

// DeleteCredentials delete credentials of the specified strategy with given body infos.
func (s *Security) DeleteCredentials(strategy, id string, options types.QueryOptions) error {
	if strategy == "" || id == "" {
		return types.NewError("Kuzzle.DeleteCredentials: strategy and id are required", 400)
	}
	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "security",
		Action:     "deleteCredentials",
		Id:         id,
		Strategy:   strategy,
	}

	go s.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return res.Error
	}

	return nil
}
