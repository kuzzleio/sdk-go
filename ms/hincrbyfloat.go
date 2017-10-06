package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
	"strconv"
)

// Hincrbyfloat increments the number stored in a hash field by the provided float value.
func (ms Ms) Hincrbyfloat(key string, field string, value float64, options types.QueryOptions) (float64, error) {
	if key == "" {
		return 0, errors.New("Ms.Hincrbyfloat: key required")
	}
	if field == "" {
		return 0, errors.New("Ms.Hincrbyfloat: field required")
	}

	result := make(chan types.KuzzleResponse)

	type body struct {
		Value float64 `json:"value"`
		Field string  `json:"field"`
	}

	query := types.KuzzleRequest{
		Controller: "ms",
		Action:     "hincrbyfloat",
		Id:         key,
		Body:       &body{Value: value, Field: field},
	}

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return 0, errors.New(res.Error.Message)
	}

	var stringResult string
	json.Unmarshal(res.Result, &stringResult)

	return strconv.ParseFloat(stringResult, 64)
}
