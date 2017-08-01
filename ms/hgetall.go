package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

/*
  Return all fields and values of a hash
*/
func (ms Ms) Hgetall(key string, options types.QueryOptions) (map[string]string, error) {
	if key == "" {
		return nil, errors.New("Ms.Hgetall: key required")
	}

	result := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Controller: "ms",
		Action:     "hgetall",
		Id:         key,
	}

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error.Message != "" {
		return nil, errors.New(res.Error.Message)
	}
	returnedResult := make(map[string]string)
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
