package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

// Rpop removes and returns the last element of a list.
func (ms Ms) Rpop(key string, options types.QueryOptions) (interface{}, error) {
	if key == "" {
		return "", errors.New("Ms.Rpop: key required")
	}

	result := make(chan *types.KuzzleResponse)

	type body struct {
		NewKey string `json:"newkey"`
	}

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "rpop",
		Id:         key,
	}
	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return "", errors.New(res.Error.Message)
	}
	var returnedResult interface{}
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
