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

// CreateOrReplaceProfile creates or replaces (if _id matches an existing one) a profile with a list of policies.
func (s *Security) CreateOrReplaceProfile(id string, body json.RawMessage, options types.QueryOptions) (*Profile, error) {
	if body == nil {
		return nil, types.NewError("Kuzzle.CreateOrReplaceProfile: body is required", 400)
	}
	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "security",
		Action:     "createOrReplaceProfile",
		Id:         id,
		Body:       body,
	}

	go s.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error.Error() != "" {
		return nil, res.Error
	}

	var profile *Profile

	json.Unmarshal(res.Result, &profile)

	return profile, nil
}
