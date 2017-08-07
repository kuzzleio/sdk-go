package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

/*
  Sets an expiration timestamp on a key. After the timestamp has been reached, the key will automatically be deleted.

  The timestamp parameter accepts an Epoch time value, in milliseconds.
*/
func (ms Ms) PexpireAt(key string, timestamp int, options types.QueryOptions) (int, error) {
	if key == "" {
		return 0, errors.New("Ms.PexpireAt: key required")
	}

	result := make(chan types.KuzzleResponse)

	type body struct {
		Timestamp int `json:"timestamp"`
	}

	query := types.KuzzleRequest{
		Controller: "ms",
		Action:     "pexpireat",
		Id:         key,
		Body:       &body{Timestamp: timestamp},
	}
	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error.Message != "" {
		return 0, errors.New(res.Error.Message)
	}
	var returnedResult int
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
