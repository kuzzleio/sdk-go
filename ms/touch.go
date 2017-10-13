package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Touch alters the last access time of one or multiple keys. A key is ignored if it does not exist.
func (ms Ms) Touch(keys []string, options types.QueryOptions) (int, error) {
	if len(keys) == 0 {
		return 0, types.NewError("Ms.Touch: please provide at least one key", 400)
	}

	result := make(chan *types.KuzzleResponse)

	type body struct {
		Keys []string `json:"keys"`
	}

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "touch",
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
