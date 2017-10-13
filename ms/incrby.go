package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Incrby increments the number stored at key by the provided integer value.
// If the key does not exist, it is set to 0 before performing the operation.
func (ms Ms) Incrby(key string, value int, options types.QueryOptions) (int, error) {
	if key == "" {
		return 0, types.NewError("Ms.Incrby: key required", 400)
	}

	result := make(chan *types.KuzzleResponse)

	type body struct {
		Value int `json:"value"`
	}

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "incrby",
		Id:         key,
		Body:       &body{Value: value},
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
