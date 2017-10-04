package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

// ZrevRangeByLex is identical to zrangebylex except that the sorted set is traversed in descending order.
func (ms Ms) ZrevRangeByLex(key string, min string, max string, options types.QueryOptions) ([]string, error) {
	if key == "" {
		return []string{}, errors.New("Ms.ZrevRangeByLex: key required")
	}
	if min == "" {
		return []string{}, errors.New("Ms.ZrevRangeByLex: min required")
	}
	if max == "" {
		return []string{}, errors.New("Ms.ZrevRangeByLex: max required")
	}

	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "zrevrangebylex",
		Id:         key,
		Min:        min,
		Max:        max,
	}

	assignZrangeOptions(query, options)

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return []string{}, errors.New(res.Error.Message)
	}

	var returnedResult []string
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
