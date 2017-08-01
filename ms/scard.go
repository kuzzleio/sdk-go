package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

/*
  Returns the number of members stored in a set of unique values.
*/
func (ms Ms) Scard(key string, options types.QueryOptions) (int, error) {
	if key == "" {
		return 0, errors.New("Ms.Scard: key required")
	}

	result := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Controller: "ms",
		Action:     "scard",
		Id:         key,
	}

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error.Message != "" {
		return 0, errors.New(res.Error.Message)
	}

	var scanResponse int
	json.Unmarshal(res.Result, &scanResponse)

	return scanResponse, nil
}
