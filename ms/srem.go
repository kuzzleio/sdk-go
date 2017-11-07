package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Srem removes members from a set of unique values.
func (ms Ms) Srem(key string, valuesToRemove []string, options types.QueryOptions) (int, error) {
	if len(valuesToRemove) == 0 {
		return 0, types.NewError("Ms.Srem: please provide at least one value to remove", 400)
	}

	result := make(chan *types.KuzzleResponse)

	type body struct {
		Members interface{} `json:"members"`
	}

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "srem",
		Id:         key,
		Body:       &body{Members: valuesToRemove},
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
