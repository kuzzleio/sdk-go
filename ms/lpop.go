package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Lpop removes and returns the first element of a list.
func (ms Ms) Lpop(key string, options types.QueryOptions) (string, error) {
	if key == "" {
		return "", types.NewError("Ms.Lpop: key required")
	}

	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "lpop",
		Id:         key,
	}

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return "", res.Error
	}
	var returnedResult string
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
