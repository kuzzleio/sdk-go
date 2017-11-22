package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Hdel removes fields from a hash
func (ms *Ms) Hdel(key string, fields []string, options types.QueryOptions) (int, error) {
	if len(fields) == 0 {
		return -1, types.NewError("Ms.Hdel: at least one hash field to remove is required", 400)
	}

	result := make(chan *types.KuzzleResponse)

	type body struct {
		Fields []string `json:"fields"`
	}

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "hdel",
		Id:         key,
		Body:       &body{Fields: fields},
	}

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return -1, res.Error
	}

	var returnedResult int
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
