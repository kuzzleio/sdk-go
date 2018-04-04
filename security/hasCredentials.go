package security

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

// HasCredentials check the existence of the specified strategy credentials for the id
func (s *Security) HasCredentials(strategy, id string, options types.QueryOptions) (bool, error) {
	if strategy == "" || id == "" {
		return false, types.NewError("Security.HasCredentials: strategy and id are required", 400)
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "security",
		Action:     "hasCredentials",
		Strategy:   strategy,
		Id:         id,
	}
	go s.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return false, res.Error
	}

	var hasCreds bool
	json.Unmarshal(res.Result, &hasCreds)

	return hasCreds, nil
}
