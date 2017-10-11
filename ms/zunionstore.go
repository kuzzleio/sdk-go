package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

// ZunionStore computes the union of the provided sorted sets and stores the result in the destination key.
// If the destination key already exists, it is overwritten.
func (ms Ms) ZunionStore(destination string, keys []string, options types.QueryOptions) (int, error) {
	if destination == "" {
		return 0, errors.New("Ms.ZunionStore: destination required")
	}
	if len(keys) == 0 {
		return 0, errors.New("Ms.ZunionStore: please provide at least one key")
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
		Action:     "zunionstore",
		Id:         destination,
		Body:       &bodyContent,
	}

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return 0, errors.New(res.Error.Message)
	}

	var scanResponse int
	json.Unmarshal(res.Result, &scanResponse)

	return scanResponse, nil
}
