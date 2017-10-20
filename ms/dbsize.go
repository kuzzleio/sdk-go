package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Dbsize returns the number of keys in the application database.
func (ms Ms) Dbsize(options types.QueryOptions) (int, error) {
	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "dbsize",
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
