package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

/*
  Trims an existing list so that it will
  contain only the specified range of elements specified.
*/
func (ms Ms) Ltrim(key string, start int, stop int, options types.QueryOptions) (string, error) {
	if key == "" {
		return "", errors.New("Ms.Ltrim: key required")
	}

	result := make(chan types.KuzzleResponse)

	type body struct {
		Start int `json:"start"`
		Stop  int `json:"stop"`
	}

	query := types.KuzzleRequest{
		Controller: "ms",
		Action:     "ltrim",
		Id:         key,
		Body:       &body{Start: start, Stop: stop},
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
