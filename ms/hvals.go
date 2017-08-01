package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

/*
  Returns all values contained in a hash.
*/
func (ms Ms) Hvals(key string, options types.QueryOptions) ([]string, error) {
	if key == "" {
		return nil, errors.New("Ms.Hval: key required")
	}

	result := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Controller: "ms",
		Action:     "hvals",
		Id:         key,
	}

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error.Message != "" {
		return nil, errors.New(res.Error.Message)
	}

	var returnedResult []string
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
