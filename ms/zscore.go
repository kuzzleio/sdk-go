package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
	"strconv"
)

// Zscore returns the score of a member in a sorted set.
func (ms Ms) Zscore(key string, member string, options types.QueryOptions) (float64, error) {
	if key == "" {
		return 0, types.NewError("Ms.Zscore: key required", 400)
	}
	if member == "" {
		return 0, types.NewError("Ms.Zscore: member required", 400)
	}

	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "zscore",
		Id:         key,
		Member:     member,
	}

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return 0, res.Error
	}

	var scanResponse string
	json.Unmarshal(res.Result, &scanResponse)

	return strconv.ParseFloat(scanResponse, 64)
}
