package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Lpush prepends the specified values to a list.
// If the key does not exist, it is created holding
// an empty list before performing the operation.
func (ms Ms) Lpush(key string, values []string, options types.QueryOptions) (int, error) {
	if len(values) == 0 {
		return -1, types.NewError("Ms.Lpush: at least one value to push is required", 400)
	}

	result := make(chan *types.KuzzleResponse)

	type body struct {
		Values []string `json:"values"`
	}

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "lpush",
		Id:         key,
		Body:       &body{Values: values},
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
