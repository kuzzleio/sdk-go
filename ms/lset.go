package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

// Lset sets the list element at index with the provided value.
func (ms Ms) Lset(key string, index int, value string, options types.QueryOptions) (string, error) {
	if key == "" {
		return "", errors.New("Ms.Lset: key required")
	}

	result := make(chan types.KuzzleResponse)

	type body struct {
		Value string `json:"value"`
		Index int    `json:"index"`
	}

	query := types.KuzzleRequest{
		Controller: "ms",
		Action:     "lset",
		Id:         key,
		Body:       &body{Value: value, Index: index},
	}

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error.Message != "" {
		return "", errors.New(res.Error.Message)
	}
	var returnedResult string
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
