package kuzzle

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
	"time"
)

// GetStatistics get Kuzzle usage statistics
func (k *Kuzzle) GetStatistics(startTime *time.Time, stopTime *time.Time, options types.QueryOptions) (*types.Statistics, error) {
	result := make(chan *types.KuzzleResponse)

	type data struct {
		startTime string `json:"startTime"`
		stopTime  string `json:"stopTime"`
	}

	var d data
	if startTime != nil {
		d = data{
			startTime.String(),
			stopTime.String(),
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
