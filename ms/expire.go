package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Expire sets an expiration timeout on a key
func (ms Ms) Expire(key string, seconds int, options types.QueryOptions) (int, error) {
	if key == "" {
		return 0, types.NewError("Ms.Expire: key required")
	}

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
		return 0, res.Error
	}
	var returnedResult int
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
