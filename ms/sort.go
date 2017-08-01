package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

/*
  Sorts and returns elements contained in a list, a set of unique values or a sorted set. By default, sorting is numeric and elements are compared by their value interpreted as double precision floating point number.
*/
func (ms Ms) Sort(key string, options types.QueryOptions) ([]interface{}, error) {
	if key == "" {
		return []interface{}{}, errors.New("Ms.Sort: key required")
	}

	result := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Controller: "ms",
		Action:     "sort",
		Id:         key,
	}

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error.Message != "" {
		return []interface{}{}, errors.New(res.Error.Message)
	}

	var returnedResult []interface{}
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
