package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

// ZremRangeByLex removes members from a sorted set where all elements have the same score, using lexicographical ordering. The min and max interval are inclusive, see the Redis documentation to change this behavior.
func (ms Ms) ZremRangeByLex(key string, min string, max string, options types.QueryOptions) (int, error) {
	if key == "" {
		return 0, errors.New("Ms.ZremRangeByLex: key required")
	}
	if min == "" {
		return 0, errors.New("Ms.ZremRangeByLex: min required")
	}
	if max == "" {
		return 0, errors.New("Ms.ZremRangeByLex: max required")
	}

	result := make(chan *types.KuzzleResponse)

	type body struct {
		Min string `json:"min"`
		Max string `json:"max"`
	}

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "zremrangebylex",
		Id:         key,
		Body:       &body{Min: min, Max: max},
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
