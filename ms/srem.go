package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

/*
  Removes members from a set of unique values.
*/
func (ms Ms) Srem(key string, valuesToRemove []string, options types.QueryOptions) (int, error) {
	if key == "" {
		return 0, errors.New("Ms.Srem: key required")
	}
	if len(valuesToRemove) == 0 {
		return 0, errors.New("Ms.Srem: please provide at least one value to remove")
	}

	result := make(chan types.KuzzleResponse)

	type body struct {
		Members interface{} `json:"members"`
	}

	query := types.KuzzleRequest{
		Controller: "ms",
		Action:     "srem",
		Id:         key,
		Body:       &body{Members: valuesToRemove},
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
