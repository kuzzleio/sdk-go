package auth

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

// GetStrategies retrieve a list of accepted fields per authentication strategy.
func (a *Auth) GetStrategies(options types.QueryOptions) ([]string, error) {
	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "auth",
		Action:     "getStrategies",
	}
	go a.kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return nil, res.Error
	}

	strategies := []string{}
	json.Unmarshal(res.Result, &strategies)

	return strategies, nil
}
