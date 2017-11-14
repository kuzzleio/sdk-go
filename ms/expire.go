package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Expire sets an expiration timeout on a key
func (ms Ms) Expire(key string, seconds int, options types.QueryOptions) (bool, error) {
	result := make(chan *types.KuzzleResponse)

	type body struct {
		Seconds int `json:"seconds"`
	}

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "expire",
		Id:         key,
		Body:       &body{Seconds: seconds},
	}

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return false, res.Error
	}

	var returnedResult int
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult == 1, nil
}
