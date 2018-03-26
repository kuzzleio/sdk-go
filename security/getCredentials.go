package security

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

func (s *Security) GetCredentials(strategy, id string, options types.QueryOptions) (json.RawMessage, error) {
	if strategy == "" || id == "" {
		return nil, types.NewError("Security.GetCredentials: strategy and id are required", 400)
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "security",
		Action:     "getCredentials",
		Strategy:   strategy,
		Id:         id,
	}
	go s.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return nil, res.Error
	}

	return res.Result, nil
}
