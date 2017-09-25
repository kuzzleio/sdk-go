package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

// SunionStore computes the union of the provided sets of unique values and stores the result in the destination key.
// If the destination key already exists, it is overwritten.
func (ms Ms) SunionStore(destination string, sets []string, options types.QueryOptions) (int, error) {
	if destination == "" {
		return 0, errors.New("Ms.SunionStore: destination required")
	}
	if len(sets) < 2 {
		return 0, errors.New("Ms.SunionStore: please provide at least 2 sets")
	}

	result := make(chan types.KuzzleResponse)

	type body struct {
		Destination string   `json:"destination"`
		Keys        []string `json:"keys"`
	}

	query := types.KuzzleRequest{
		Controller: "ms",
		Action:     "sunionstore",
		Body:       &body{Destination: destination, Keys: sets},
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
