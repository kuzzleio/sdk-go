package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

// Spop removes and returns one or more elements at random from a set of unique values.
func (ms Ms) Spop(key string, options types.QueryOptions) (interface{}, error) {
	if key == "" {
		return "", errors.New("Ms.Spop: key required")
	}

	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "spop",
		Id:         key,
	}

	if options != nil {
		if options.GetCount() != 0 {
			query.Count = options.GetCount()
		}
	}

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error.Message != "" {
		return "", errors.New(res.Error.Message)
	}

	var returnedResult interface{}
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
