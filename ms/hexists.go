package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Hexists check if a field exists in a hash
func (ms Ms) Hexists(key string, field string, options types.QueryOptions) (int, error) {
	if key == "" {
		return -1, types.NewError("Ms.Hexists: key required", 400)
	}
	if field == "" {
		return -1, types.NewError("Ms.Hexists: field required", 400)
	}

	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "hexists",
		Id:         key,
		Field:      field,
	}

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return -1, res.Error
	}

	var returnedResult int
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
