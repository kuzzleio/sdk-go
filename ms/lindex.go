package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Returns the element at the provided index in a list.
func (ms *Ms) Lindex(key string, index int, options types.QueryOptions) (*string, error) {
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
		return nil, res.Error
	}
	var returnedResult *string
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
