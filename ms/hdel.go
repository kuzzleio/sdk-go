package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

// Hdel removes fields from a hash
func (ms Ms) Hdel(key string, fields []string, options types.QueryOptions) (int, error) {
	if key == "" {
		return 0, errors.New("Ms.Hdel: key required")
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

	if res.Error.Message != "" {
		return 0, errors.New(res.Error.Message)
	}

	var returnedResult int
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
