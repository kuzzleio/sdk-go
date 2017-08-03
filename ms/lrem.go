package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

/*
  Removes the first count occurences of elements equal to value from a list.
*/
func (ms Ms) Lrem(key string, count int, value string, options types.QueryOptions) (int, error) {
	if key == "" {
		return 0, errors.New("Ms.Lrem: key required")
	}

	result := make(chan types.KuzzleResponse)

	type body struct {
		Value string `json:"value"`
		Count int    `json:"count"`
	}

	query := types.KuzzleRequest{
		Controller: "ms",
		Action:     "lrem",
		Id:         key,
		Body:       &body{Value: value, Count: count},
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
