package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

/*
  Returns the value of a key, or null if the key doesn’t exist.
*/
func (ms Ms) Get(key string, options types.QueryOptions) (interface{}, error) {
	if key == "" {
		return "", errors.New("Ms.Get: key required")
	}

	result := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Controller: "ms",
		Action:     "get",
		Id:         key,
	}

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error.Message != "" {
		return "", errors.New(res.Error.Message)
	}

	var returnedResult interface{}
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
