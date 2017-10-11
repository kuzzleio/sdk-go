package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// ZinterStore computes the intersection of the provided sorted sets and stores the result in the destination key.
// If the destination key already exists, it is overwritten.
func (ms Ms) ZinterStore(destination string, keys []string, options types.QueryOptions) (int, error) {
	if destination == "" {
		return 0, types.NewError("Ms.ZinterStore: destination required")
	}
	if len(keys) == 0 {
		return 0, types.NewError("Ms.ZinterStore: please provide at least one key")
	}

	result := make(chan *types.KuzzleResponse)

	type body struct {
		Keys      []string `json:"keys"`
		Weights   []int    `json:"weight,omitempty"`
		Aggregate string   `json:"aggregate,omitempty"`
	}

	bodyContent := body{Keys: keys}

	if len(options.GetWeights()) > 0 {
		bodyContent.Weights = options.GetWeights()
	}

	if options.GetAggregate() != "" {
		bodyContent.Aggregate = options.GetAggregate()
	}

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "zinterstore",
		Id:         destination,
		Body:       &bodyContent,
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
