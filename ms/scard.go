package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Scard returns the number of members stored in a set of unique values.
func (ms Ms) Scard(key string, options types.QueryOptions) (int, error) {
	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "scard",
		Id:         key,
	}

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return 0, res.Error
	}

	var scanResponse int
	json.Unmarshal(res.Result, &scanResponse)

	return scanResponse, nil
}
