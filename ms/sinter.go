package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Sinter returns the intersection of the provided sets of unique values.
func (ms Ms) Sinter(keys []string, options types.QueryOptions) ([]string, error) {
	if len(keys) == 0 {
		return nil, types.NewError("Ms.Sinter: please provide at least one key", 400)
	}

	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "sinter",
		Keys:       keys,
	}

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return nil, res.Error
	}
	var returnedResult []string
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
