package security

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

// CreateCredentials create credentials of the specified strategy with given body infos.
func (s *Security) CreateCredentials(strategy, id, body string, options types.QueryOptions) (json.RawMessage, error) {
	if strategy == "" || id == "" || body == "" {
		return nil, types.NewError("Kuzzle.CreateCredentials: strategy, id and body are required", 400)
	}
	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "security",
		Action:     "createCredentials",
		Id:         id,
		Strategy:   strategy,
		Body:       body,
	}

	go s.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return nil, res.Error
	}

	return res.Result, nil
}
