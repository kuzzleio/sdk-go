package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Hval returns all values contained in a hash.
func (ms Ms) Hvals(key string, options types.QueryOptions) ([]string, error) {
	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "hvals",
		Id:         key,
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
