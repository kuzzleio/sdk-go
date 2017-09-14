package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

// Strlen returns the length of a value stored at key.
func (ms Ms) Strlen(key string, options types.QueryOptions) (int, error) {
	if key == "" {
		return 0, errors.New("Ms.Strlen: key required")
	}

	result := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Controller: "ms",
		Action:     "strlen",
		Id:         key,
	}

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error.Message != "" {
		return 0, errors.New(res.Error.Message)
	}

	var returnedResult int
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
