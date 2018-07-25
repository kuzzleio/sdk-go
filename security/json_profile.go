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

type jsonProfile struct {
	Id     string         `json:"_id"`
	Source types.Policies `json:"_source"`
}

type jsonProfileSearchResult struct {
	Total    int            `json:"total"`
	Hits     []*jsonProfile `json:"hits"`
	ScrollId string         `json:"scrollId"`
}

func (j *jsonProfile) jsonProfileToProfile() *Profile {
	p := &Profile{
		Id:       j.Id,
		Policies: j.Source.Policies,
	}

	return p
}

func (p *Profile) ProfileToJson() ([]byte, error) {
	j := &jsonProfile{
		Id: p.Id,
		Source: types.Policies{
			Policies: p.Policies,
		},
	}

	return json.Marshal(j)
}
