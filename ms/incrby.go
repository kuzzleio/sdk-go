package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

/*
  Increments the number stored at key by the provided integer value.
  If the key does not exist, it is set to 0 before performing the operation.
*/
func (ms Ms) Incrby(key string, value int, options types.QueryOptions) (int, error) {
	if key == "" {
		return 0, errors.New("Ms.Incr: key required")
	}

	result := make(chan types.KuzzleResponse)

	type body struct {
		Value int `json:"value"`
	}

	query := types.KuzzleRequest{
		Controller: "ms",
		Action:     "incrby",
		Id:         key,
		Body:				&body{Value: value},
	}

	go ms.Kuzzle.Query(query, nil, result)

	res := <-result

	if res.Error.Message != "" {
		return 0, errors.New(res.Error.Message)
	}
	var returnedResult int
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
