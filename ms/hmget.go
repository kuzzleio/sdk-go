package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Hmget returns the values of the specified hash’s fields.
func (ms Ms) Hmget(key string, fields []string, options types.QueryOptions) ([]string, error) {
	if key == "" {
		return nil, types.NewError("Ms.Hmget: key required")
	}

	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "hmget",
		Id:         key,
		Fields:     fields,
	}

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return nil, res.Error
	}

	var returnedResult []string
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
