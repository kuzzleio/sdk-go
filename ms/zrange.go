package ms

import (
	"encoding/json"
	"fmt"
	"github.com/kuzzleio/sdk-go/types"
	"strconv"
	"strings"
)

// Zrange returns elements from a sorted set depending on their position in the set, from a start position index to a stop position index (inclusive).
// First position starts at 0.
func (ms Ms) Zrange(key string, start int, stop int, options types.QueryOptions) ([]*types.MSSortedSet, error) {
	if key == "" {
		return nil, types.NewError("Ms.Zrange: key required", 400)
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

	if res.Error != nil {
		return nil, res.Error
	}

	var returnedResult []string
	json.Unmarshal(res.Result, &returnedResult)

	return mapZrangeResults(returnedResult)
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

func mapZrangeResults(results []string) ([]*types.MSSortedSet, error) {
	buffer := ""
	sortedSet := make([]*types.MSSortedSet, 0, len(results))

	for _, value := range results {
		if buffer == "" {
			buffer = value
		} else {
			score, err := strconv.ParseFloat(value, 64)

			if err != nil {
				return nil, types.NewError(err.Error())
			}

			sortedSet = append(sortedSet, &types.MSSortedSet{Member: buffer, Score: score})
			buffer = ""
		}
	}

	return sortedSet, nil
}
