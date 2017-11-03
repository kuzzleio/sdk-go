package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Time returns the current server time.
func (ms Ms) Time(options types.QueryOptions) ([]int, error) {
	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "time",
	}
	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return nil, res.Error
	}
	var returnedResult []int
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
