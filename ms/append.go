package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Append a value to a key
func (ms Ms) Append(key string, value string, options types.QueryOptions) (int, error) {
	result := make(chan *types.KuzzleResponse)

	type body struct {
		Value string `json:"value"`
	}

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "append",
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
