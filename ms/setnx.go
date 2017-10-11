package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

// SetNx sets a value on a key, only if it does not already exist.
func (ms Ms) SetNx(key string, value interface{}, options types.QueryOptions) (int, error) {
	if key == "" {
		return 0, errors.New("Ms.SetNx: key required")
	}

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
		return 0, errors.New(res.Error.Message)
	}
	var returnedResult int
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
