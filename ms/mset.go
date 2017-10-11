package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

// Mset sets the provided keys to their respective values.
// If a key does not exist, it is created. Otherwise, the key’s value is overwritten.
func (ms Ms) Mset(entries []*types.MSKeyValue, options types.QueryOptions) (string, error) {
	if len(entries) == 0 {
		return "", errors.New("Ms.Mset: please provide at least one key/value entry")
	}

	result := make(chan *types.KuzzleResponse)

	type body struct {
		Entries []*types.MSKeyValue `json:"entries"`
	}

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "mset",
		Body:       &body{Entries: entries},
	}
	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return "", errors.New(res.Error.Message)
	}
	var returnedResult string
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
