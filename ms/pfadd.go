package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

// Pfadd Adds elements to an HyperLogLog data structure.
func (ms Ms) Pfadd(key string, elements []string, options types.QueryOptions) (int, error) {
	if key == "" {
		return 0, errors.New("Ms.Pfadd: key required")
	}
	if len(elements) == 0 {
		return 0, errors.New("Ms.Pfadd: please provide at least one element")
	}

	result := make(chan *types.KuzzleResponse)

	type body struct {
		Elements []string `json:"elements"`
	}

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "pfadd",
		Id:         key,
		Body:       &body{Elements: elements},
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
