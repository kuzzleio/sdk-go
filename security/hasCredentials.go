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

package security

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

// HasCredentials check the existence of the specified strategy credentials for the id
func (s *Security) HasCredentials(strategy, id string, options types.QueryOptions) (bool, error) {
	if strategy == "" || id == "" {
		return false, types.NewError("Security.HasCredentials: strategy and id are required", 400)
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "security",
		Action:     "hasCredentials",
		Strategy:   strategy,
		Id:         id,
	}
	go s.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Error() != "" {
		return false, res.Error
	}

	var hasCreds bool
	json.Unmarshal(res.Result, &hasCreds)

	return hasCreds, nil
}
