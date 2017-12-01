package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Pfadd Adds elements to an HyperLogLog data structure.
func (ms *Ms) Pfadd(key string, elements []string, options types.QueryOptions) (bool, error) {
	if len(elements) == 0 {
		return false, types.NewError("Ms.Pfadd: please provide at least one element to add", 400)
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
		return false, res.Error
	}
	var returnedResult int
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult == 1, nil
}
