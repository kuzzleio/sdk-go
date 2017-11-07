package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
	"strconv"
)

// Zremrangebyrank Removes members from a sorted set with
// their position in the set between start and stop (inclusive).
// Positions are 0-based, meaning the first member of the set has a position of 0.
func (ms Ms) Zremrangebyrank(key string, min int, max int, options types.QueryOptions) (int, error) {
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

	if res.Error != nil {
		return 0, res.Error
	}

	var returnedResult int
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
