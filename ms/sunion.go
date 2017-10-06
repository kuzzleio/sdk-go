package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

// Sunion returns the union of the provided sets of unique values.
func (ms Ms) Sunion(sets []string, options types.QueryOptions) ([]string, error) {
	if len(sets) < 2 {
		return []string{}, errors.New("Ms.Sunion: please provide at least 2 sets")
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
		return []string{}, errors.New(res.Error.Message)
	}

	var returnedResult []string
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
