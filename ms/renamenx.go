package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

/*
  Renames a key to newkey, only if newkey does not already exist.
*/
func (ms Ms) RenameNx(key string, newkey string, options types.QueryOptions) (int, error) {
	if key == "" {
		return 0, errors.New("Ms.RenameNx: key required")
	}
	if newkey == "" {
		return 0, errors.New("Ms.RenameNx: newkey required")
	}

	result := make(chan types.KuzzleResponse)

	type body struct {
		NewKey string `json:"newkey"`
	}

	query := types.KuzzleRequest{
		Controller: "ms",
		Action:     "renamenx",
		Id:         key,
		Body:       &body{NewKey: newkey},
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
