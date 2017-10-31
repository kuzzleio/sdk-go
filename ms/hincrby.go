package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Hincrby increments the number stored in a hash field by the provided integer value.
func (ms Ms) Hincrby(key string, field string, value int, options types.QueryOptions) (int, error) {
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

	if res.Error != nil {
		return 0, res.Error
	}
	var returnedResult int
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
