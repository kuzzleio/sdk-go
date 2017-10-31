package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Pfadd Adds elements to an HyperLogLog data structure.
func (ms Ms) Pfadd(key string, elements []string, options types.QueryOptions) (int, error) {
	if len(elements) == 0 {
		return 0, types.NewError("Ms.Pfadd: please provide at least one element to add", 400)
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

	if res.Error != nil {
		return 0, res.Error
	}
	var returnedResult int
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
