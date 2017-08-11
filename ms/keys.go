package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

/*
  Returns all keys matching the provided pattern.
*/
func (ms Ms) Keys(pattern string, options types.QueryOptions) ([]string, error) {
	if pattern == "" {
		return nil, errors.New("Ms.Keys: pattern required")
	}

	result := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Controller: "ms",
		Action:     "keys",
		Pattern:    pattern,
	}

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error.Message != "" {
		return nil, errors.New(res.Error.Message)
	}
	var returnedResult []string
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
