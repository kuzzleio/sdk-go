package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

// Decrby decrements the value of a key by a given value
func (ms Ms) Decrby(key string, value int, options types.QueryOptions) (int, error) {
	if key == "" {
		return 0, errors.New("Ms.Decrby: key required")
	}

	result := make(chan *types.KuzzleResponse)

	type body struct {
		Value int `json:"value"`
	}

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "decrby",
		Id:         key,
		Body:       &body{Value: value},
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
