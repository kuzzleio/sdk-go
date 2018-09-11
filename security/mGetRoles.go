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

// MGetRoles gets all roles matching with given ids
func (s *Security) MGetRoles(ids []string, options types.QueryOptions) ([]*Role, error) {
	if len(ids) == 0 {
		return nil, types.NewError("Security.MGetRoles: ids array can't be nil", 400)
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "security",
		Action:     "mGetRoles",
		Body: struct {
			Ids []string `json:"ids"`
		}{ids},
	}
	go s.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Error() != "" {
		return nil, res.Error
	}

	var fetchedRaw jsonRoleSearchResult
	var fetchedRoles []*Role
	json.Unmarshal(res.Result, &fetchedRaw)

	for _, jsonRoleRaw := range fetchedRaw.Hits {
		fetchedRoles = append(fetchedRoles, jsonRoleRaw.jsonRoleToRole())
	}

	return fetchedRoles, nil

}
