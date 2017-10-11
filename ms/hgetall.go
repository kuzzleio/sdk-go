package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Hgetall returns all fields and values of a hash
func (ms Ms) Hgetall(key string, options types.QueryOptions) (map[string]string, error) {
	if key == "" {
		return nil, types.NewError("Ms.Hgetall: key required")
	}

	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "hgetall",
		Id:         key,
	}

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return nil, res.Error
	}
	returnedResult := make(map[string]string)
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
