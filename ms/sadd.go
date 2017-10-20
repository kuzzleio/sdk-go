package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Sadd creates a key holding the provided value, or overwrites it if it already exists.
func (ms Ms) Sadd(key string, values []string, options types.QueryOptions) (int, error) {
	if key == "" {
		return 0, types.NewError("Ms.Sadd: key required", 400)
	}
	if len(values) == 0 {
		return 0, types.NewError("Ms.Sadd: please provide at least one value", 400)
	}

	result := make(chan *types.KuzzleResponse)

	type body struct {
		Members []string `json:"members"`
	}

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "sadd",
		Id:         key,
		Body:       &body{Members: values},
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
