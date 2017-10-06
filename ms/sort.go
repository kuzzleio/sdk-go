package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

// Sort sorts and returns elements contained in a list, a set of unique values or a sorted set.
// By default, sorting is numeric and elements are compared by their value interpreted
// as double precision floating point number.
func (ms Ms) Sort(key string, options types.QueryOptions) ([]interface{}, error) {
	if key == "" {
		return []interface{}{}, errors.New("Ms.Sort: key required")
	}

	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "sort",
		Id:         key,
	}

	if options != nil {
		type body struct {
			Limit     []int    `json:"limit,omitempty"`
			By        string   `json:"by,omitempty"`
			Direction string   `json:"direction,omitempty"`
			Get       []string `json:"get,omitempty"`
			Alpha     bool     `json:"alpha,omitempty"`
		}

		bodyContent := &body{}

		if options.GetBy() != "" {
			bodyContent.By = options.GetBy()
		}

		if options.GetDirection() != "" {
			bodyContent.Direction = options.GetDirection()
		}

		if options.GetGet() != nil {
			bodyContent.Get = options.GetGet()
		}

		if options.GetLimit() != nil {
			bodyContent.Limit = options.GetLimit()
		}

		bodyContent.Alpha = options.GetAlpha()

		query.Body = bodyContent
	}

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return []interface{}{}, errors.New(res.Error.Message)
	}

	var returnedResult []interface{}
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
