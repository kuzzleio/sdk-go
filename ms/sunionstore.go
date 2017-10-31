package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// SunionStore computes the union of the provided sets of unique values and stores the result in the destination key.
// If the destination key already exists, it is overwritten.
func (ms Ms) SunionStore(destination string, sets []string, options types.QueryOptions) (int, error) {
	if len(sets) == 0 {
		return 0, types.NewError("Ms.SunionStore: please provide at least 1 set", 400)
	}

	result := make(chan *types.KuzzleResponse)

	type body struct {
		Destination string   `json:"destination"`
		Keys        []string `json:"keys"`
	}

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "sunionstore",
		Body:       &body{Destination: destination, Keys: sets},
	}

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return 0, res.Error
	}

	var returnedResult int
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
