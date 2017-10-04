package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

// Lindex returns all keys matching the provided pattern.
func (ms Ms) Lindex(key string, index int, options types.QueryOptions) (string, error) {
	if key == "" {
		return "", errors.New("Ms.Lindex: key required")
	}

	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "lindex",
		Id:         key,
		Idx:        index,
	}

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error.Message != "" {
		return "", errors.New(res.Error.Message)
	}
	var returnedResult string
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
