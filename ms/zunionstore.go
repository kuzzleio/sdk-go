package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

/*
  Computes the union of the provided sorted sets and stores the result in the destination key.

  If the destination key already exists, it is overwritten.
*/
func (ms Ms) ZunionStore(destination string, keys []string, options types.QueryOptions) (int, error) {
	if destination == "" {
		return 0, errors.New("Ms.ZunionStore: destination required")
	}
	if len(keys) == 0 {
		return 0, errors.New("Ms.ZunionStore: please provide at least one key")
	}

	result := make(chan types.KuzzleResponse)

	type body struct {
		Keys []string `json:"keys"`
	}

	query := types.KuzzleRequest{
		Controller: "ms",
		Action:     "zunionstore",
		Id:         destination,
		Body:       &body{Keys: keys},
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
