package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
	"strconv"
)

// ZrangeByScore returns all the elements in the sorted set at key with a score between min and max (inclusive). The elements are considered to be ordered from low to high scores.
func (ms Ms) ZrangeByScore(key string, min float64, max float64, options types.QueryOptions) ([]*types.MSSortedSet, error) {
	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "zrangebyscore",
		Id:         key,
		Min:        strconv.FormatFloat(min, 'f', 6, 64),
		Max:        strconv.FormatFloat(max, 'f', 6, 64),
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
