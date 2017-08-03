package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

/*
  Sets multiple fields at once in a hash.
*/
func (ms Ms) Hmset(key string, entries []types.MsHashField, options types.QueryOptions) (string, error) {
	if key == "" {
		return "", errors.New("Ms.Hmset: key required")
	}

	result := make(chan types.KuzzleResponse)

	type body struct {
		Entries []types.MsHashField `json:"entries"`
	}

	query := types.KuzzleRequest{
		Controller: "ms",
		Action:     "hmset",
		Id:         key,
		Body:       &body{Entries: entries},
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
