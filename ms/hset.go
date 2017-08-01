package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

/*
  Sets a field and its value in a hash.
  If the key does not exist, a new key holding a hash is created.
*/
func (ms Ms) Hset(key string, field string, value string, options types.QueryOptions) (int, error) {
	if key == "" {
		return 0, errors.New("Ms.Hset: key required")
	}

	result := make(chan types.KuzzleResponse)

	type body struct {
		Field string `json:"field"`
		Value string `json:"value"`
	}

	query := types.KuzzleRequest{
		Controller: "ms",
		Action:     "hset",
		Id:         key,
		Body:       &body{Value: value, Field: field},
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
