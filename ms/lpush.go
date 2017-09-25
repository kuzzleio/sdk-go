package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

// Lpush prepends the specified values to a list.
// If the key does not exist, it is created holding
// an empty list before performing the operation.
func (ms Ms) Lpush(key string, values []string, options types.QueryOptions) (int, error) {
	if key == "" {
		return 0, errors.New("Ms.Lpush: key required")
	}

	result := make(chan types.KuzzleResponse)

	type body struct {
		Values []string `json:"values"`
	}

	query := types.KuzzleRequest{
		Controller: "ms",
		Action:     "lpush",
		Id:         key,
		Body:       &body{Values: values},
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
