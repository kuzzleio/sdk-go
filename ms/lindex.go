package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Lindex returns all keys matching the provided pattern.
func (ms Ms) Lindex(key string, index int, options types.QueryOptions) (string, error) {
	if key == "" {
		return "", types.NewError("Ms.Lindex: key required")
	}

	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "lindex",
		Id:         key,
		Idx:        index,
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
