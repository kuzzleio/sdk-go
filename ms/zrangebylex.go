package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

// ZrangeByLex returns elements in a sorted set where all members have equal score, using lexicographical ordering. The min and max values are inclusive by default. To change this behavior, please check the full documentation.
func (ms Ms) ZrangeByLex(key string, min string, max string, options types.QueryOptions) ([]string, error) {
	if key == "" {
		return []string{}, errors.New("Ms.ZrangeByLex: key required")
	}
	if min == "" {
		return []string{}, errors.New("Ms.ZrangeByLex: min required")
	}
	if max == "" {
		return []string{}, errors.New("Ms.ZrangeByLex: max required")
	}

	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "zrangebylex",
		Id:         key,
		Min:        min,
		Max:        max,
	}

	assignZrangeOptions(query, options)

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error.Message != "" {
		return []string{}, errors.New(res.Error.Message)
	}

	var returnedResult []string
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
