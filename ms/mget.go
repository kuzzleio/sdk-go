package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

/*
  Returns the values of the provided keys.
*/
func (ms Ms) Mget(keys []string, options types.QueryOptions) ([]string, error) {
	if len(keys) == 0 {
		return []string{}, errors.New("Ms.Mget: please provide at least one key")
	}

	result := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Controller: "ms",
		Action:     "mget",
		Keys:       keys,
	}
	go ms.Kuzzle.Query(query, nil, result)

	res := <-result

	if res.Error.Message != "" {
		return []string{}, errors.New(res.Error.Message)
	}
	var returnedResult []string
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}