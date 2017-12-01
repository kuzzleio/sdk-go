package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Sunion returns the union of the provided sets of unique values.
func (ms *Ms) Sunion(sets []string, options types.QueryOptions) ([]string, error) {
	if len(sets) == 0 {
		return nil, types.NewError("Ms.Sunion: please provide at least 1 set", 400)
	}

	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "sunion",
		Keys:       sets,
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
