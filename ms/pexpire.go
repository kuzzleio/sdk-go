package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

/*
  Sets a timeout (in milliseconds) on a key. After the timeout has expired, the key will automatically be deleted.
*/
func (ms Ms) Pexpire(key string, ttl int, options types.QueryOptions) (int, error) {
	if key == "" {
		return 0, errors.New("Ms.Pexpire: key required")
	}

	result := make(chan types.KuzzleResponse)

	type body struct {
		Milliseconds int `json:"milliseconds"`
	}

	query := types.KuzzleRequest{
		Controller: "ms",
		Action:     "pexpire",
		Id:         key,
		Body:       &body{Milliseconds: ttl},
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
