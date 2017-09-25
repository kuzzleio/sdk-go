package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

// SdiffStore computes the difference between the set of unique values stored at key
// and the other provided sets, and stores the result in the key stored at destination.
// If the destination key already exists, it is overwritten.
func (ms Ms) SdiffStore(key string, sets []string, destination string, options types.QueryOptions) (int, error) {
	if key == "" {
		return 0, errors.New("Ms.SdiffStore: key required")
	}
	if len(sets) == 0 {
		return 0, errors.New("Ms.SdiffStore: please provide at least one set")
	}
	if destination == "" {
		return 0, errors.New("Ms.SdiffStore: destination required")
	}

	result := make(chan types.KuzzleResponse)

	type body struct {
		Keys        []string `json:"keys"`
		Destination string   `json:"destination"`
	}

	query := types.KuzzleRequest{
		Controller: "ms",
		Action:     "sdiffstore",
		Id:         key,
		Body:       &body{Keys: sets, Destination: destination},
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
