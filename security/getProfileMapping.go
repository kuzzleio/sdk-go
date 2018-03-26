package security

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

// GetProfileMapping gets an array of strategy's fieldnames for each strategies
func (s *Security) GetProfileMapping(options types.QueryOptions) (json.RawMessage, error) {
	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "security",
		Action:     "getProfileMapping",
	}
	go s.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return nil, res.Error
	}

	return res.Result, nil
}
