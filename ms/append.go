package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

/*
  Append a value to a key
*/
func (ms Ms) Append(key string, value string, options types.QueryOptions) (int, error) {
	if key == "" {
		return 0, errors.New("Ms.Append: key required")
	}

	result := make(chan types.KuzzleResponse)

	type body struct {
		Value string `json:"value"`
	}

	query := types.KuzzleRequest{
		Controller: "ms",
		Action:     "append",
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
