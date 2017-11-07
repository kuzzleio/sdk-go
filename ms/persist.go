package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Persist removes the expiration delay or timestamp from a key, making it persistent.
func (ms Ms) Persist(key string, options types.QueryOptions) (int, error) {
	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "persist",
		Id:         key,
	}
	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return 0, res.Error
	}
	var returnedResult int
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
