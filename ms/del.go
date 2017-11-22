package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Del deletes keys
func (ms *Ms) Del(keys []string, options types.QueryOptions) (int, error) {
	if len(keys) == 0 {
		return 0, types.NewError("Ms.Del: at least one key is required", 400)
	}

	result := make(chan *types.KuzzleResponse)

	type body struct {
		Keys []string `json:"keys"`
	}

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "del",
		Body:       &body{Keys: keys},
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
