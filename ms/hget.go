package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Hget returns the field's value of a hash
func (ms *Ms) Hget(key string, field string, options types.QueryOptions) (*string, error) {
	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "hget",
		Id:         key,
		Field:      field,
	}

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return nil, res.Error
	}

	var returnedResult *string
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
