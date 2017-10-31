package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
	"strconv"
)

// Hincrbyfloat increments the number stored in a hash field by the provided float value.
func (ms Ms) Hincrbyfloat(key string, field string, value float64, options types.QueryOptions) (float64, error) {
	result := make(chan *types.KuzzleResponse)

	type body struct {
		Value float64 `json:"value"`
		Field string  `json:"field"`
	}

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "hincrbyfloat",
		Id:         key,
		Body:       &body{Value: value, Field: field},
	}

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return 0, res.Error
	}

	var stringResult string
	json.Unmarshal(res.Result, &stringResult)

	return strconv.ParseFloat(stringResult, 64)
}
