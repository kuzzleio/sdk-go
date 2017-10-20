package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Flushdb delete all keys from the database
func (ms Ms) Flushdb(options types.QueryOptions) (string, error) {
	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "flushdb",
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
