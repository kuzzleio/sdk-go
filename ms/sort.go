package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Sort sorts and returns elements contained in a list, a set of unique values or a sorted set.
// By default, sorting is numeric and elements are compared by their value interpreted
// as double precision floating point number.
func (ms *Ms) Sort(key string, options types.QueryOptions) ([]string, error) {
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

		if options.By() != "" {
			bodyContent.By = options.By()
		}

		if options.Direction() != "" {
			bodyContent.Direction = options.Direction()
		}

		if options.Get() != nil {
			bodyContent.Get = options.Get()
		}

		if options.Limit() != nil {
			bodyContent.Limit = options.Limit()
		}

		bodyContent.Alpha = options.Alpha()

		query.Body = bodyContent
	}

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return nil, res.Error
	}

	var returnedResult []string
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
