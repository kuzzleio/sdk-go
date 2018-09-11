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

	"github.com/kuzzleio/sdk-go/security"
	"github.com/kuzzleio/sdk-go/types"
)

// UpdateSelf updates the current User object in Kuzzle's database layer.
func (a *Auth) UpdateSelf(data json.RawMessage, options types.QueryOptions) (*security.User, error) {
	if options == nil {
		options = types.NewQueryOptions()
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "auth",
		Action:     "updateSelf",
		Body:       data,
	}
	go a.kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Error() != "" {
		return nil, res.Error
	}

	type jsonUser struct {
		Id         string          `json:"_id"`
		Source     json.RawMessage `json:"_source"`
		ProfileIds []string        `json:"profileIds"`
	}
	u := &jsonUser{}
	json.Unmarshal(res.Result, u)

	var content map[string]interface{}
	json.Unmarshal(u.Source, &content)
	user := &security.User{
		Id:         u.Id,
		Content:    content,
		ProfileIds: u.ProfileIds,
	}
	return user, nil
}
