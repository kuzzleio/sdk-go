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

// CreateRestrictedUser create credentials of the specified strategy with given body infos.
func (s *Security) CreateRestrictedUser(body json.RawMessage, options types.QueryOptions) (json.RawMessage, error) {
	if body == nil {
		return nil, types.NewError("Kuzzle.CreateRestrictedUser: strategy, id and body are required", 400)
	}
	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "security",
		Action:     "createRestrictedUser",
		Body:       body,
	}

	if options != nil {
		query.Id = options.ID()
	}

	go s.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error.Error() != "" {
		return nil, res.Error
	}

	return res.Result, nil
}
