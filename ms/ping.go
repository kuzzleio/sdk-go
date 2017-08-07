package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

/*
  Pings the memory storage database.
*/
func (ms Ms) Ping(options types.QueryOptions) (string, error) {
	result := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Controller: "ms",
		Action:     "ping",
	}
	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error.Message != "" {
		return "", errors.New(res.Error.Message)
	}
	var returnedResult string
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
