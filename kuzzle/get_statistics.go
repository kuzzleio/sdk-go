package kuzzle

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
	"time"
)

// GetStatistics get Kuzzle usage statistics
func (k Kuzzle) GetStatistics(timestamp *time.Time, options types.QueryOptions) (*types.Statistics, error) {
	result := make(chan *types.KuzzleResponse)

	type data struct {
		since string `json:"since"`
	}

	var d data
	if timestamp != nil {
		d = data{
			timestamp.String(),
		}
	}

	query := &types.KuzzleRequest{
		Controller: "server",
		Action:     "getLastStats",
		Body:       d,
	}

	go k.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return nil, res.Error
	}

	s := &types.Statistics{}
	json.Unmarshal(res.Result, s)

	return s, nil
}
