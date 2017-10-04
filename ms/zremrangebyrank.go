package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
	"strconv"
)

// ZremRangeByRank Removes members from a sorted set with
// their position in the set between start and stop (inclusive).
// Positions are 0-based, meaning the first member of the set has a position of 0.
func (ms Ms) ZremRangeByRank(key string, min int, max int, options types.QueryOptions) (int, error) {
	if key == "" {
		return 0, errors.New("Ms.ZremRangeByRank: key required")
	}

	result := make(chan *types.KuzzleResponse)

	type body struct {
		Start string `json:"start"`
		Stop  string `json:"stop"`
	}

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "zremrangebyrank",
		Id:         key,
		Body:       &body{Start: strconv.Itoa(min), Stop: strconv.Itoa(max)},
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
