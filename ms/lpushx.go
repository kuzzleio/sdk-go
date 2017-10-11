package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Lpushx prepends the specified value to a list,
// only if the key already exists and if it holds a list.
func (ms Ms) Lpushx(key string, value string, options types.QueryOptions) (int, error) {
	if key == "" {
		return 0, types.NewError("Ms.Lpushx: key required")
	}

	result := make(chan *types.KuzzleResponse)

	type body struct {
		Value string `json:"value"`
	}

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "lpushx",
		Id:         key,
		Body:       &body{Value: value},
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
