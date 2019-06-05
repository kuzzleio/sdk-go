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

type ProfileSearchResult struct {
	Aggregations json.RawMessage `json:"aggregations"`
	Hits         []*Profile
	Total        int `json:"total"`
	Fetched      int
	ScrollId     string `json:"scrollId"`
	kuzzle       types.IKuzzle
	request      *types.KuzzleRequest
	response     *types.KuzzleResponse
	options      types.QueryOptions
}

func (sr *ProfileSearchResult) Request() *types.KuzzleRequest {
	return sr.request
}

func (sr *ProfileSearchResult) Response() *types.KuzzleResponse {
	return sr.response
}

func (sr *ProfileSearchResult) Options() types.QueryOptions {
	return sr.options
}

// Next returns the next page of profiles
func (psr *ProfileSearchResult) Next() (*ProfileSearchResult, error) {
	sr, err := types.NewSearchResult(psr.kuzzle, "scrollProfiles", psr.request, psr.options, psr.response)
	if err != nil {
		return nil, err
	}

	nsr, err := sr.Next()

	if err != nil {
		return nil, err
	}

	npsr := &ProfileSearchResult{
		Aggregations: nsr.Aggregations,
		Total:        nsr.Total,
		Fetched:      nsr.Fetched,
		ScrollId:     nsr.ScrollId,
		kuzzle:       psr.kuzzle,
		request:      psr.request,
		response:     psr.response,
		options:      psr.options,
	}
	var jprofiles []jsonProfile
	err = json.Unmarshal(nsr.Hits, &jprofiles)

	if err != nil {
		return nil, err
	}

	for _, jprofile := range jprofiles {
		npsr.Hits = append(npsr.Hits, jprofile.jsonProfileToProfile())
	}

	return npsr, nil
}
