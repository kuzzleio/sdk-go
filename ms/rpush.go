package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Rpush appends the specified value at the end of a list,
// only if the key already exists and if it holds a list.
func (ms Ms) Rpush(source string, values []string, options types.QueryOptions) (int, error) {
	if source == "" {
		return 0, types.NewError("Ms.Rpush: source required")
	}
	if len(values) == 0 {
		return 0, types.NewError("Ms.Rpush: please provide at least one value")
	}

	result := make(chan *types.KuzzleResponse)

	type body struct {
		Values []string `json:"values"`
	}

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "rpush",
		Id:         source,
		Body:       &body{Values: values},
	}
	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return 0, res.Error
	}
	var returnedResult int
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
