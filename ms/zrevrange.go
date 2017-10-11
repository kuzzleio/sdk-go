package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// ZrevRange is identical to zrange, except that the sorted set is traversed in descending order.
func (ms Ms) ZrevRange(key string, start int, stop int, options types.QueryOptions) ([]*types.MSSortedSet, error) {
	if key == "" {
		return []*types.MSSortedSet{}, types.NewError("Ms.ZrevRange: key required")
	}

	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "zrevrange",
		Id:         key,
		Start:      start,
		Stop:       stop,
	}

	assignZrangeOptions(query, options)

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return []*types.MSSortedSet{}, res.Error
	}

	var returnedResult []string
	json.Unmarshal(res.Result, &returnedResult)

	return mapZrangeResults(returnedResult), nil
}
