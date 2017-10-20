package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// PexipreAt sets an expiration timestamp on a key.
// After the timestamp has been reached, the key will automatically be deleted.
// The timestamp parameter accepts an Epoch time value, in milliseconds.
func (ms Ms) PexpireAt(key string, timestamp int, options types.QueryOptions) (int, error) {
	if key == "" {
		return 0, types.NewError("Ms.PexpireAt: key required", 400)
	}

	result := make(chan *types.KuzzleResponse)

	type body struct {
		Timestamp int `json:"timestamp"`
	}

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "pexpireat",
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
