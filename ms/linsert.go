package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

/*
  Inserts a value in a list, either before or after the reference pivot value.
*/
func (ms Ms) Linsert(key string, position string, pivot string, value string, options types.QueryOptions) (int, error) {
	if key == "" {
		return 0, errors.New("Ms.Linsert: key required")
	}

	result := make(chan types.KuzzleResponse)

	type body struct {
		Position string `json:"position"`
		Pivot string `json:"pivot"`
		Value string `json:"value"`
	}

	query := types.KuzzleRequest{
		Controller: "ms",
		Action:     "linsert",
		Id:         key,
		Body: 	    &body{Position: position, Pivot: pivot, Value: value},
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
