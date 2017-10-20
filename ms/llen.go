package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Llen counts the number of items in a list.
func (ms Ms) Llen(key string, options types.QueryOptions) (int, error) {
	if key == "" {
		return 0, types.NewError("Ms.Llen: key required", 400)
	}

	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "llen",
		Id:         key,
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
