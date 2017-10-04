package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

// SinterStore computes the intersection of the provided sets of unique values and stores
// the result in the destination key.
// If the destination key already exists, it is overwritten.
func (ms Ms) SinterStore(destination string, keys []string, options types.QueryOptions) (int, error) {
	if destination == "" {
		return 0, errors.New("Ms.SinterStore: destination required")
	}
	if len(keys) == 0 {
		return 0, errors.New("Ms.SinterStore: please provide at least one key")
	}

	result := make(chan *types.KuzzleResponse)

	type body struct {
		Destination string   `json:"destination"`
		Keys        []string `json:"keys"`
	}

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "sinterstore",
		Body:       &body{Destination: destination, Keys: keys},
	}

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error.Message != "" {
		return 0, errors.New(res.Error.Message)
	}
	var returnedResult int
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
