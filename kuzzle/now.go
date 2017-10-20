package kuzzle

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Now retrieves the current Kuzzle time.
func (k Kuzzle) Now(options types.QueryOptions) (int, error) {
	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "server",
		Action:     "now",
	}
	go k.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return -1, res.Error
	}

	type now struct {
		Now int `json:"now"`
	}

	n := now{}
	json.Unmarshal(res.Result, &n)

	return n.Now, nil
}
