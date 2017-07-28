package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

func (ms Ms) Append(key string, value string, options types.QueryOptions) (int64, error) {
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
		Id:					key,
		Body: 		  &body{Value: value},
	}
	go ms.Kuzzle.Query(query, nil, result)

	res := <-result

	if res.Error.Message != "" {
		return 0, errors.New(res.Error.Message)
	}
	var appendResult int64
	json.Unmarshal(res.Result, &appendResult)

	return appendResult, nil
}
