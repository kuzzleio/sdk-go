package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Decrby decrements the value of a key by a given value
func (ms Ms) Decrby(key string, value int, options types.QueryOptions) (int, error) {
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

	if res.Error != nil {
		return 0, res.Error
	}
	var returnedResult int
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
