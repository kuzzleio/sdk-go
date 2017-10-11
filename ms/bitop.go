package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

// Bitop performs a bitwise operation between multiple keys (containing string values) and stores the result in the destination key.
func (ms Ms) Bitop(key string, operation string, keys []string, options types.QueryOptions) (int, error) {
	if key == "" {
		return 0, errors.New("Ms.Bitop: key required")
	}

	result := make(chan *types.KuzzleResponse)

	type body struct {
		Operation string   `json:"operation"`
		Keys      []string `json:"keys"`
	}

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "bitop",
		Id:         key,
		Body:       &body{Operation: operation, Keys: keys},
	}
	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return 0, errors.New(res.Error.Message)
	}
	var returnedResult int
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
