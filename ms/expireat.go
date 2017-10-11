package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Expireat sets an expiration timestamp to a key
func (ms Ms) Expireat(key string, timestamp int, options types.QueryOptions) (int, error) {
	if key == "" {
		return 0, types.NewError("Ms.Expireat: key required")
	}

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
		return 0, res.Error
	}
	var returnedResult int
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
