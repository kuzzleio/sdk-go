package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Hset sets a field and its value in a hash.
// if the key does not exist, a new key holding a hash is created.
func (ms Ms) Hset(key string, field string, value string, options types.QueryOptions) (int, error) {
	if key == "" {
		return 0, types.NewError("Ms.Hset: key required", 400)
	}
	if field == "" {
		return 0, types.NewError("Ms.Hset: field required", 400)
	}

	result := make(chan *types.KuzzleResponse)

	type body struct {
		Field string `json:"field"`
		Value string `json:"value"`
	}

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "hset",
		Id:         key,
		Body:       &body{Field: field, Value: value},
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
