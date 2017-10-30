package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Get returns the value of a key, or null if the key doesn’t exist.
func (ms Ms) Get(key string, options types.QueryOptions) (string, error) {
	if key == "" {
		return "", types.NewError("Ms.Get: key required", 400)
	}

	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "get",
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
