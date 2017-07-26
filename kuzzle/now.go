package kuzzle

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

/*
 * Returns the current Kuzzle UTC timestamp
 */
func (k Kuzzle) Now(options types.QueryOptions) (int, error) {
	result := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Controller: "server",
		Action:     "now",
	}
	go k.Query(query, options, result)

	res := <-result

	if res.Error.Message != "" {
		return 0, errors.New(res.Error.Message)
	}

	type now struct {
		Now int `json:"now"`
	}

	n := now{}
	json.Unmarshal(res.Result, &n)

	return n.Now, nil
}
