package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Zrevrange is identical to zrange, except that the sorted set is traversed in descending order.
func (ms *Ms) Zrevrange(key string, start int, stop int, options types.QueryOptions) ([]*types.MSSortedSet, error) {
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
		return nil, res.Error
	}

	var returnedResult []string
	json.Unmarshal(res.Result, &returnedResult)

	return mapZrangeResults(returnedResult)
}
