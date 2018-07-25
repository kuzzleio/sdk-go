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

package auth

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

// CreateMyCredentials create credentials of the specified strategy for the current user.
func (a *Auth) CreateMyCredentials(strategy string, credentials json.RawMessage, options types.QueryOptions) (json.RawMessage, error) {
	if strategy == "" {
		return nil, types.NewError("Kuzzle.CreateMyCredentials: strategy is required", 400)
	}
	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "auth",
		Action:     "createMyCredentials",
		Body:       credentials,
		Strategy:   strategy,
	}
	go a.kuzzle.Query(query, options, result)

	res := <-result

	if res.Error.Error() != "" {
		return nil, res.Error
	}

	return res.Result, nil
}
