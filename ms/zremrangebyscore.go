package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
	"strconv"
)

/*
  Removes members from a sorted set with a score between min and max (inclusive by default).
*/
func (ms Ms) ZremRangeByScore(key string, min float64, max float64, options types.QueryOptions) (int, error) {
	if key == "" {
		return 0, errors.New("Ms.ZremRangeByScore: key required")
	}

	result := make(chan types.KuzzleResponse)

	type body struct {
		Min string `json:"min"`
		Max string `json:"max"`
	}

	query := types.KuzzleRequest{
		Controller: "ms",
		Action:     "zremrangebyscore",
		Id:         key,
		Body:       &body{Min: strconv.FormatFloat(min, 'f', 6, 64), Max: strconv.FormatFloat(max, 'f', 6, 64)},
	}

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error.Message != "" {
		return 0, errors.New(res.Error.Message)
	}

	var returnedResult int
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
