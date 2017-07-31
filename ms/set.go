package ms

import (
"encoding/json"
"errors"
"github.com/kuzzleio/sdk-go/types"
)

/*
  Creates a key holding the provided value, or overwrites it if it already exists.
*/
func (ms Ms) Set(key string, value interface{}, options types.QueryOptions) (string, error) {
	if key == "" {
		return "", errors.New("Ms.Set: key required")
	}

	result := make(chan types.KuzzleResponse)

	type body struct {
		Value interface{} `json:"value"`
	}

	query := types.KuzzleRequest{
		Controller: "ms",
		Action:     "set",
		Id:         key,
		Body:       &body{Value: value},
	}

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error.Message != "" {
		return "", errors.New(res.Error.Message)
	}
	var returnedResult string
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
