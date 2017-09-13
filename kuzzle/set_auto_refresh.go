package kuzzle

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

// autorefresh status setter for the provided data index name
func (k Kuzzle) SetAutoRefresh(index string, autoRefresh bool, options types.QueryOptions) (bool, error) {
	if index == "" {
		if k.defaultIndex == "" {
			return false, errors.New("Kuzzle.SetAutoRefresh: index required")
		}
		index = k.defaultIndex
	}

	result := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Controller: "index",
		Action:     "setAutoRefresh",
		Index:      index,
		Body: struct {
			AutoRefresh bool `json:"autoRefresh"`
		}{autoRefresh},
	}
	go k.Query(query, options, result)

	res := <-result

	if res.Error.Message != "" {
		return false, errors.New(res.Error.Message)
	}

	type autoRefreshResponse struct {
		Response bool `json:"response"`
	}

	var r autoRefreshResponse
	json.Unmarshal(res.Result, &r)

	return r.Response, nil
}
