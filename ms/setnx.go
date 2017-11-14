package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// SetNx sets a value on a key, only if it does not already exist.
func (ms Ms) Setnx(key string, value interface{}, options types.QueryOptions) (bool, error) {
	result := make(chan *types.KuzzleResponse)

	type body struct {
		Value interface{} `json:"value"`
	}

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "setnx",
		Id:         key,
		Body:       &body{Value: value},
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
