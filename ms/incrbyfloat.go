package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
	"strconv"
)

// Incrbyfloat increments the number stored at key by the provided float value.
// If the key does not exist, it is set to 0 before performing the operation.
func (ms Ms) Incrbyfloat(key string, value float64, options types.QueryOptions) (float64, error) {
	if key == "" {
		return 0, errors.New("Ms.Incrbyfloat: key required")
	}

	result := make(chan *types.KuzzleResponse)

	type body struct {
		Value float64 `json:"value"`
	}

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "incrbyfloat",
		Id:         key,
		Body:       &body{Value: value},
	}

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error.Message != "" {
		return 0, errors.New(res.Error.Message)
	}

	var stringResult string
	json.Unmarshal(res.Result, &stringResult)

	return strconv.ParseFloat(stringResult, 64)
}
