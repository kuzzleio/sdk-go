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

type jsonRole struct {
	Id     string            `json:"_id"`
	Source types.Controllers `json:"_source"`
}

type jsonRoleSearchResult struct {
	Hits     []*jsonRole `json:"hits"`
	Total    int         `json:"total"`
	ScrollId string      `json:"scrollId"`
}

func (j *jsonRole) jsonRoleToRole() *Role {
	r := &Role{}
	r.Id = j.Id
	r.Controllers = j.Source.Controllers

	return r
}

func (r *Role) RoleToJson() ([]byte, error) {
	j := &jsonRole{
		Id: r.Id,
		Source: types.Controllers{
			Controllers: r.Controllers,
		},
	}
	return json.Marshal(j)
}
