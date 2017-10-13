package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// RenameNx renames a key to newkey, only if newkey does not already exist.
func (ms Ms) RenameNx(key string, newkey string, options types.QueryOptions) (int, error) {
	if key == "" {
		return 0, types.NewError("Ms.RenameNx: key required", 400)
	}
	if newkey == "" {
		return 0, types.NewError("Ms.RenameNx: newkey required", 400)
	}

	result := make(chan *types.KuzzleResponse)

	type body struct {
		NewKey string `json:"newkey"`
	}

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "renamenx",
		Id:         key,
		Body:       &body{NewKey: newkey},
	}
	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return 0, res.Error
	}
	var returnedResult int
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
