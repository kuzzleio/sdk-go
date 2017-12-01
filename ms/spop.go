package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Spop removes and returns one or more elements at random from a set of unique values.
func (ms *Ms) Spop(key string, options types.QueryOptions) ([]string, error) {
	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "spop",
		Id:         key,
	}

	if options != nil {
		if options.Count() != 0 {
			query.Count = options.Count()
		}
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
