package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Incr increments the number stored at key by 1.
// If the key does not exist, it is set to 0 before performing the operation.
func (ms Ms) Incr(key string, options types.QueryOptions) (int, error) {
	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "incr",
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
