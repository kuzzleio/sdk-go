package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Zrevrangebylex is identical to zrangebylex except that the sorted set is traversed in descending order.
func (ms *Ms) Zrevrangebylex(key string, min string, max string, options types.QueryOptions) ([]string, error) {
	if min == "" || max == "" {
		return nil, types.NewError("Ms.Zrevrangebylex: an empty string is not a valid string range item", 400)
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
		return nil, res.Error
	}

	var returnedResult []string
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
