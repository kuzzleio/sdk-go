package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Hstrlen returns the string length of a field’s value in a hash.
func (ms Ms) Hstrlen(key string, field string, options types.QueryOptions) (int, error) {
	if key == "" {
		return -1, types.NewError("Ms.Hstrlen: key required", 400)
	}
	if field == "" {
		return -1, types.NewError("Ms.Hstrlen: field required", 400)
	}

	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "hstrlen",
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
