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

type RoleSearchResult struct {
	Aggregations json.RawMessage `json:"aggregations"`
	Hits         []*Role
	Total        int `json:"total"`
	Fetched      int
	ScrollId     string `json:"scrollId"`
	kuzzle       types.IKuzzle
	request      *types.KuzzleRequest
	response     *types.KuzzleResponse
	options      types.QueryOptions
}

func (sr *RoleSearchResult) Request() *types.KuzzleRequest {
	return sr.request
}

func (sr *RoleSearchResult) Response() *types.KuzzleResponse {
	return sr.response
}

func (sr *RoleSearchResult) Options() types.QueryOptions {
	return sr.options
}

// Next returns the next page of roles
func (rsr *RoleSearchResult) Next() (*RoleSearchResult, error) {
	sr, err := types.NewSearchResult(rsr.kuzzle, "scrollRoles", rsr.request, rsr.options, rsr.response)
	if err != nil {
		return nil, err
	}

	nsr, err := sr.Next()

	if err != nil {
		return nil, err
	}

	nrsr := &RoleSearchResult{
		Aggregations: nsr.Aggregations,
		Total:        nsr.Total,
		Fetched:      nsr.Fetched,
		ScrollId:     nsr.ScrollId,
		kuzzle:       rsr.kuzzle,
		request:      rsr.request,
		response:     rsr.response,
		options:      rsr.options,
	}
	var jroles []jsonRole
	err = json.Unmarshal(nsr.Hits, &jroles)

	if err != nil {
		return nil, err
	}

	for _, jrole := range jroles {
		nrsr.Hits = append(nrsr.Hits, jrole.jsonRoleToRole())
	}

	return nrsr, nil
}
