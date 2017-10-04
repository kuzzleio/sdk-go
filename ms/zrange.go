package ms

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kuzzleio/sdk-go/types"
	"strconv"
	"strings"
)

// Zrange returns elements from a sorted set depending on their position in the set, from a start position index to a stop position index (inclusive).
// First position starts at 0.
func (ms Ms) Zrange(key string, start int, stop int, options types.QueryOptions) ([]*types.MSSortedSet, error) {
	if key == "" {
		return []*types.MSSortedSet{}, errors.New("Ms.Zrange: key required")
	}

	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "zrange",
		Id:         key,
		Start:      start,
		Stop:       stop,
	}

	assignZrangeOptions(query, options)

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error.Message != "" {
		return []*types.MSSortedSet{}, errors.New(res.Error.Message)
	}

	var returnedResult []string
	json.Unmarshal(res.Result, &returnedResult)

	return mapZrangeResults(returnedResult), nil
}

func assignZrangeOptions(query *types.KuzzleRequest, options types.QueryOptions) {
	opts := make([]interface{}, 0, 1)

	opts = append(opts, "withscores")

	if options != nil {
		if len(options.GetLimit()) != 0 {
			query.Limit = strings.Trim(strings.Join(strings.Fields(fmt.Sprint(options.GetLimit())), ","), "[]")
		}
	}

	query.Options = []interface{}(opts)
}

func mapZrangeResults(results []string) []*types.MSSortedSet {
	var buffer = ""
	var sortedSet = make([]*types.MSSortedSet{}, results.len())

	for _, value := range results {
		if buffer == "" {
			buffer = value
		} else {
			var score, _ = strconv.ParseFloat(value, 64)
			sortedSet = append(sortedSet, &types.MSSortedSet{Member: buffer, Score: score})
			buffer = ""
		}
	}

	return sortedSet
}
