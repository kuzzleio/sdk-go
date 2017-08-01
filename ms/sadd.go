package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

/*
  Creates a key holding the provided value, or overwrites it if it already exists.
*/
func (ms Ms) Sadd(key string, values []string, options types.QueryOptions) (int, error) {
	if key == "" {
		return 0, errors.New("Ms.Sadd: key required")
	}
	if len(values) == 0 {
		return 0, errors.New("Ms.Sadd: please provide at least one value")
	}

	result := make(chan types.KuzzleResponse)

	type body struct {
		Members []string `json:"members"`
	}

	query := types.KuzzleRequest{
		Controller: "ms",
		Action:     "sadd",
		Id:         key,
		Body:    		&body{Members: values},
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
