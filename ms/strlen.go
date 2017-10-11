package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Strlen returns the length of a value stored at key.
func (ms Ms) Strlen(key string, options types.QueryOptions) (int, error) {
	if key == "" {
		return 0, types.NewError("Ms.Strlen: key required")
	}

	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "strlen",
		Id:         key,
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
