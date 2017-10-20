package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// RandomKey returns a random key from the memory storage.
func (ms Ms) RandomKey(options types.QueryOptions) (string, error) {
	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "randomkey",
	}
	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return "", res.Error
	}
	var returnedResult string
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
