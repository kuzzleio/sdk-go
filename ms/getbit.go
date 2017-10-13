package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Getbit returns the bit value at offset, in the string value stored in a key.
func (ms Ms) Getbit(key string, offset int, options types.QueryOptions) (int, error) {
	if key == "" {
		return -1, types.NewError("Ms.Getbit: key required", 400)
	}

	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "getbit",
		Id:         key,
		Offset:     offset,
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
