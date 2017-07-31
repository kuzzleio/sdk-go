package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

/*
  Set an expiration timestamp to a key
*/
func (ms Ms) Expireat(key string, timestamp int, options types.QueryOptions) (int, error) {
	if key == "" {
		return 0, errors.New("Ms.Expireat: key required")
	}

	result := make(chan types.KuzzleResponse)

	type body struct {
		Timestamp int `json:"timestamp"`
	}

	query := types.KuzzleRequest{
		Controller: "ms",
		Action:     "expireat",
		Id:         key,
		Body:       &body{Timestamp: timestamp},
	}

	go ms.Kuzzle.Query(query, nil, result)

	res := <-result

	if res.Error.Message != "" {
		return 0, errors.New(res.Error.Message)
	}
	var returnedResult int
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
