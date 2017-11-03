package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Srandmember returns one or more members of a set of unique values, at random.
// If count is provided and is positive, the returned values are unique.
// If count is negative, a set member can be returned multiple times.
func (ms Ms) Srandmember(key string, options types.QueryOptions) ([]string, error) {
	count := 1

	if options != nil {
		count = options.GetCount()

		if count < 1 {
			count = 1
		}
	}

	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "srandmember",
		Id:         key,
		Count:			count,
	}

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return nil, res.Error
	}

	if count == 1 {
		var returnedResult string
		json.Unmarshal(res.Result, &returnedResult)

		return []string{returnedResult}, nil
	} else {
		var returnedResult []string
		json.Unmarshal(res.Result, &returnedResult)

		return returnedResult, nil
	}
}
