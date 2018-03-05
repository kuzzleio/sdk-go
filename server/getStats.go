package server

import (
	"encoding/json"
	"time"

	"github.com/kuzzleio/sdk-go/types"
)

//GetStats get Kuzzle usage statistics
func (s *Server) GetStats(startTime *time.Time, stopTime *time.Time, options types.QueryOptions) (json.RawMessage, error) {
	result := make(chan *types.KuzzleResponse)

	type data struct {
		StartTime string `json:"startTime"`
		StopTime  string `json:"stopTime"`
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

	go s.k.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return nil, res.Error
	}

	return res.Result, nil
}
