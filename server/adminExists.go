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

	"github.com/kuzzleio/sdk-go/types"
)

// AdminExists checks if an administrator account has been created, and return a boolean as a result.
func (s *Server) AdminExists(options types.QueryOptions) (bool, error) {
	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "server",
		Action:     "adminExists",
	}

	go s.k.Query(query, options, result)

	res := <-result

	if res.Error.Error() != "" {
		return false, res.Error
	}

	type exists struct {
		Exists bool `json:"exists"`
	}

	r := exists{}
	json.Unmarshal(res.Result, &r)
	return r.Exists, nil
}
