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
	scrollAction string
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
	err = json.Unmarshal(nsr.Hits, npsr.Hits)

	if err != nil {
		return nil, err
	}

	return npsr, nil
}
