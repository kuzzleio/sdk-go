package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Sdiff returns the difference between the set of unique values stored at key and the other provided sets.
func (ms Ms) Sdiff(key string, sets []string, options types.QueryOptions) ([]string, error) {
	if key == "" {
		return []string{}, types.NewError("Ms.Sdiff: key required")
	}
	if len(sets) == 0 {
		return []string{}, types.NewError("Ms.Sdiff: please provide at least one set")
	}

	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "sdiff",
		Id:         key,
		Keys:       sets,
	}

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return []string{}, res.Error
	}
	var returnedResult []string
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
