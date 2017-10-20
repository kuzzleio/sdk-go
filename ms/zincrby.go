package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
	"strconv"
)

// ZincrBy increments the score of a member in a sorted set by the provided value.
func (ms Ms) ZincrBy(key string, member string, increment float64, options types.QueryOptions) (float64, error) {
	if key == "" {
		return 0, types.NewError("Ms.ZincrBy: key required", 400)
	}
	if member == "" {
		return 0, types.NewError("Ms.ZincrBy: member required", 400)
	}

	result := make(chan *types.KuzzleResponse)

	type body struct {
		Member    string  `json:"member"`
		Increment float64 `json:"value"`
	}

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "zincrby",
		Id:         key,
		Body:       &body{Member: member, Increment: increment},
	}

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return 0, res.Error
	}

	var returnedResult string
	json.Unmarshal(res.Result, &returnedResult)

	converted, err := strconv.ParseFloat(returnedResult, 64)

	if (err != nil) {
		err = types.NewError(err.Error())
	}

	return converted, err
}
