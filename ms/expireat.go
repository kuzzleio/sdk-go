package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Expireat sets an expiration timestamp to a key
func (ms *Ms) Expireat(key string, timestamp int, options types.QueryOptions) (bool, error) {
	result := make(chan *types.KuzzleResponse)

	type body struct {
		Timestamp int `json:"timestamp"`
	}

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "expireat",
		Id:         key,
		Body:       &body{Timestamp: timestamp},
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
