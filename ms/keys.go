package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Keys returns all keys matching the provided pattern.
func (ms Ms) Keys(pattern string, options types.QueryOptions) ([]string, error) {
	if pattern == "" {
		return nil, types.NewError("Ms.Keys: pattern required", 400)
	}

	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "keys",
		Pattern:    pattern,
	}

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return nil, res.Error
	}
	var returnedResult []string
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
