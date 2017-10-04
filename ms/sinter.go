package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

// Sinter returns the intersection of the provided sets of unique values.
func (ms Ms) Sinter(keys []string, options types.QueryOptions) ([]string, error) {
	if len(keys) == 0 {
		return []string{}, errors.New("Ms.Sinter: please provide at least one key")
	}

	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "sinter",
		Keys:       keys,
	}

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error.Message != "" {
		return []string{}, errors.New(res.Error.Message)
	}
	var returnedResult []string
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
