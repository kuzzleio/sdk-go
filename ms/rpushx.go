package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

/*
  Appends the specified value at the end of a list, only if the key already exists and if it holds a list.
*/
func (ms Ms) Rpushx(key string, value string, options types.QueryOptions) (int, error) {
	if key == "" {
		return 0, errors.New("Ms.Rpushx: key required")
	}
	if value == "" {
		return 0, errors.New("Ms.Rpushx: value required")
	}

	result := make(chan types.KuzzleResponse)

	type body struct {
		Value string `json:"value"`
	}

	query := types.KuzzleRequest{
		Controller: "ms",
		Action:     "rpushx",
		Id:         key,
		Body:       &body{Value: value},
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
