package kuzzle

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

// Retrieves the list of indexes stored in Kuzzle.
func (k Kuzzle) ListIndexes(options types.QueryOptions) ([]string, error) {
	result := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Controller: "index",
		Action:     "list",
	}

	go k.Query(query, options, result)

	res := <-result

	type indexes struct {
		Indexes []string `json:"indexes"`
	}

	if res.Error.Message != "" {
		return nil, errors.New(res.Error.Message)
	}

	var i indexes
	json.Unmarshal(res.Result, &i)

	return i.Indexes, nil
}
