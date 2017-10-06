package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

// Hstrlen returns the string length of a field’s value in a hash.
func (ms Ms) Hstrlen(key string, field string, options types.QueryOptions) (int, error) {
	if key == "" {
		return 0, errors.New("Ms.Hstrlen: key required")
	}
	if field == "" {
		return 0, errors.New("Ms.Hstrlen: field required")
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
		return 0, errors.New(res.Error.Message)
	}

	var returnedResult int
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
