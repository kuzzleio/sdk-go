package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

/*
  Check if the specified keys exist
*/
func (ms Ms) Exists(keys []string, options types.QueryOptions) (int, error) {
	result := make(chan types.KuzzleResponse)

	type body struct {
		Keys []string `json:"value"`
	}

	query := types.KuzzleRequest{
		Controller: "ms",
		Action:     "exists",
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
