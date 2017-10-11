package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Zrem removes members from a sorted set.
func (ms Ms) Zrem(key string, members []string, options types.QueryOptions) (int, error) {
	if key == "" {
		return 0, types.NewError("Ms.Zrem: key required")
	}
	if len(members) == 0 {
		return 0, types.NewError("Ms.Zrem: please provide at least one member")
	}

	result := make(chan *types.KuzzleResponse)

	type body struct {
		Members []string `json:"members"`
	}

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "zrem",
		Id:         key,
		Body:       &body{Members: members},
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
