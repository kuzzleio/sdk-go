package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Sismember checks if member is a member of the set of unique values stored at key.
func (ms Ms) Sismember(key string, member string, options types.QueryOptions) (int, error) {
	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "sismember",
		Id:         key,
		Member:     member,
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
