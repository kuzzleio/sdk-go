package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Hexists check if a field exists in a hash
func (ms Ms) Hexists(key string, field string, options types.QueryOptions) (int, error) {
	if key == "" {
		return 0, types.NewError("Ms.Hexists: key required")
	}
	if field == "" {
		return 0, types.NewError("Ms.Hexists: field required")
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
		return 0, res.Error
	}

	var returnedResult int
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
