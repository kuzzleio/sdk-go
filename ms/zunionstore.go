package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Zunionstore computes the union of the provided sorted sets and stores the result in the destination key.
// If the destination key already exists, it is overwritten.
func (ms *Ms) Zunionstore(destination string, keys []string, options types.QueryOptions) (int, error) {
	if len(keys) == 0 {
		return 0, types.NewError("Ms.Zunionstore: please provide at least one key", 400)
	}

	result := make(chan *types.KuzzleResponse)

	type body struct {
		Keys      []string `json:"keys"`
		Weights   []int    `json:"weight,omitempty"`
		Aggregate string   `json:"aggregate,omitempty"`
	}

	bodyContent := body{Keys: keys}

	if options != nil {
		if len(options.Weights()) > 0 {
			bodyContent.Weights = options.Weights()
		}

		if options.Aggregate() != "" {
			bodyContent.Aggregate = options.Aggregate()
		}
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
		return 0, res.Error
	}

	var scanResponse int
	json.Unmarshal(res.Result, &scanResponse)

	return scanResponse, nil
}
