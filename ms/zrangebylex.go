package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// ZrangeByLex returns elements in a sorted set where all members have equal score, using lexicographical ordering. The min and max values are inclusive by default. To change this behavior, please check the full documentation.
func (ms Ms) ZrangeByLex(key string, min string, max string, options types.QueryOptions) ([]string, error) {
	if key == "" {
		return nil, types.NewError("Ms.ZrangeByLex: key required", 400)
	}
	if min == "" {
		return nil, types.NewError("Ms.ZrangeByLex: min required", 400)
	}
	if max == "" {
		return nil, types.NewError("Ms.ZrangeByLex: max required", 400)
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

	if res.Error != nil {
		return nil, res.Error
	}

	var returnedResult []string
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
