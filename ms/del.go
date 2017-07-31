package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

/*
  Delete keys
*/
func (ms Ms) Del(keys []string, options types.QueryOptions) (int, error) {
	result := make(chan types.KuzzleResponse)

	type body struct {
		Keys []string `json:"value"`
	}

	query := types.KuzzleRequest{
		Controller: "ms",
		Action:     "del",
		Body:       &body{Keys: keys},
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
