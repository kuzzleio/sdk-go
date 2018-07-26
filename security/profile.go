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
	"github.com/kuzzleio/sdk-go/types"
)

type Profile struct {
	Id       string `json:"_id"`
	Policies []*types.Policy
}

type ProfileSearchResult struct {
	Hits     []*Profile
	Total    int    `json:"total"`
	ScrollId string `json:"scrollId"`
}

func NewProfile(id string, policies []*types.Policy) *Profile {
	return &Profile{
		Id:       id,
		Policies: policies,
	}
}
