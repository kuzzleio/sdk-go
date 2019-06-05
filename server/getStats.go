// Copyright 2015-2018 Kuzzle
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	"encoding/json"
	"time"

	"github.com/kuzzleio/sdk-go/types"
)

// GetStats get Kuzzle usage statistics
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
		Action:     "getStats",
		Body:       d,
	}

	go s.k.Query(query, options, result)

	res := <-result

	if res.Error.Error() != "" {
		return nil, res.Error
	}

	return res.Result, nil
}
