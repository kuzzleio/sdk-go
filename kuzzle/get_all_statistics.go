package kuzzle

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

/*
 * Get all Kuzzle usage statistics frames
 */
func (k Kuzzle) GetAllStatistics(options types.QueryOptions) ([]types.Statistics, error) {
	result := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Controller: "server",
		Action:     "getAllStats",
	}

	type stats struct {
		Hits []json.RawMessage `json:"hits"`
	}

	go k.Query(query, nil, result)

	res := <-result

	if res.Error.Message != "" {
		return nil, errors.New(res.Error.Message)
	}

	s := stats{}
	json.Unmarshal(res.Result, &s)

	var stat []types.Statistics
	for _, hit := range s.Hits {
		h := types.Statistics{}

		json.Unmarshal(hit, &h)
		stat = append(stat, h)
	}

	return stat, nil
}
