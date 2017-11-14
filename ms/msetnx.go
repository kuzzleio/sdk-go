package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// MsetNx sets the provided keys to their respective values, only if they do not exist.
// If a key exists, then the whole operation is aborted and no key is set.
func (ms Ms) Msetnx(entries []*types.MSKeyValue, options types.QueryOptions) (bool, error) {
	if len(entries) == 0 {
		return false, types.NewError("Ms.Msetnx: please provide at least one key/value entry", 400)
	}

	result := make(chan *types.KuzzleResponse)

	type body struct {
		Entries []*types.MSKeyValue `json:"entries"`
	}

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "msetnx",
		Body:       &body{Entries: entries},
	}
	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return false, res.Error
	}
	var returnedResult int
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult == 1, nil
}
