package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

// Hincrby increments the number stored in a hash field by the provided integer value.
func (ms Ms) Hincrby(key string, field string, value int, options types.QueryOptions) (int, error) {
	if key == "" {
		return 0, errors.New("Ms.Hincrby: key required")
	}
	if field == "" {
		return 0, errors.New("Ms.Hincrby: field required")
	}

	result := make(chan *types.KuzzleResponse)

	type body struct {
		Value int    `json:"value"`
		Field string `json:"field"`
	}

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "hincrby",
		Id:         key,
		Body:       &body{Value: value, Field: field},
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
