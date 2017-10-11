package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// SisMember checks if member is a member of the set of unique values stored at key.
func (ms Ms) SisMember(key string, member string, options types.QueryOptions) (int, error) {
	if key == "" {
		return 0, types.NewError("Ms.SisMember: key required")
	}
	if member == "" {
		return 0, types.NewError("Ms.SisMember: member required")
	}

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
