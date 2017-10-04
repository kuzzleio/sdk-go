package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
	"strconv"
)

// ZrevRangeByScore is identical to zrangebyscore except that the sorted set is traversed in descending order.
func (ms Ms) ZrevRangeByScore(key string, min float64, max float64, options types.QueryOptions) ([]*types.MSSortedSet, error) {
	if key == "" {
		return []*types.MSSortedSet{}, errors.New("Ms.ZrevRangeByScore: key required")
	}

	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "zrevrangebyscore",
		Id:         key,
		Min:        strconv.FormatFloat(min, 'f', 6, 64),
		Max:        strconv.FormatFloat(max, 'f', 6, 64),
	}

	assignZrangeOptions(query, options)

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return []*types.MSSortedSet{}, errors.New(res.Error.Message)
	}

	var returnedResult []string
	json.Unmarshal(res.Result, &returnedResult)

	return mapZrangeResults(returnedResult), nil
}
