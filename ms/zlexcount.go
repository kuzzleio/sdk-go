package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

// ZlexCount counts elements in a sorted set where all members have equal score, using lexicographical ordering. The min and max values are inclusive by default. To change this behavior, please check the syntax detailed in the Redis documentation.
func (ms Ms) ZlexCount(key string, min string, max string, options types.QueryOptions) (int, error) {
	if key == "" {
		return 0, errors.New("Ms.ZlexCount: key required")
	}
	if min == "" {
		return 0, errors.New("Ms.ZlexCount: min required")
	}
	if max == "" {
		return 0, errors.New("Ms.ZlexCount: max required")
	}

	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "zlexcount",
		Id:         key,
		Min:        min,
		Max:        max,
	}

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return 0, errors.New(res.Error.Message)
	}

	var returnedResult int
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
