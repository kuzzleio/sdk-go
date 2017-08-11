package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

/*
  Alters the last access time of one or multiple keys. A key is ignored if it does not exist.
*/
func (ms Ms) Touch(keys []string, options types.QueryOptions) (int, error) {
	if len(keys) == 0 {
		return 0, errors.New("Ms.Touch: please provide at least one key")
	}

	result := make(chan types.KuzzleResponse)

	type body struct {
		Keys []string `json:"keys"`
	}

	query := types.KuzzleRequest{
		Controller: "ms",
		Action:     "touch",
		Body:       &body{Keys: keys},
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
