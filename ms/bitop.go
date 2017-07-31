package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

func (ms Ms) Bitop(key string, operation string, keys []string, options types.QueryOptions) (int, error) {
	if key == "" {
		return 0, errors.New("Ms.Bitop: key required")
	}

	result := make(chan types.KuzzleResponse)

	type body struct {
		Operation string `json:"operation"`
		Keys []string `json:"keys"`
	}

	query := types.KuzzleRequest{
		Controller: "ms",
		Action:     "bitop",
		Id:         key,
		Body:       &body{Operation: operation, Keys:keys},
	}
	go ms.Kuzzle.Query(query, nil, result)

	res := <-result

	if res.Error.Message != "" {
		return 0, errors.New(res.Error.Message)
	}
	var bitopResult int
	json.Unmarshal(res.Result, &bitopResult)

	return bitopResult, nil
}
