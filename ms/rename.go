package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

/*
  Renames a key to newkey. If newkey already exists, it is overwritten.
*/
func (ms Ms) Rename(key string, newkey string, options types.QueryOptions) (string, error) {
	if key == "" {
		return "", errors.New("Ms.Rename: key required")
	}
	if newkey == "" {
		return "", errors.New("Ms.Rename: newkey required")
	}

	result := make(chan types.KuzzleResponse)

	type body struct {
		NewKey string `json:"newkey"`
	}

	query := types.KuzzleRequest{
		Controller: "ms",
		Action:     "rename",
		Id:         key,
		Body:       &body{NewKey: newkey},
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
