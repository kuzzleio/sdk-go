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

type jsonUser struct {
	Id     string          `json:"_id"`
	Source json.RawMessage `json:"_source"`
}

type jsonUserSearchResult struct {
	Total    int         `json:"total"`
	Hits     []*jsonUser `json:"hits"`
	ScrollId string      `json:"scrollId"`
}

func (j *jsonUser) jsonUserToUser() *User {
	u := &User{
		Id: j.Id,
	}

	userData := &types.UserData{}
	json.Unmarshal(j.Source, userData)
	if userData.ProfileIds != nil {
		u.ProfileIds = userData.ProfileIds
	}

	m := map[string]interface{}{}
	json.Unmarshal(j.Source, &m)

	for k, v := range m {
		if k != "profileIds" && k != "credentials" && v != nil {
			if u.Content == nil {
				u.Content = make(map[string]interface{})
			}
			u.Content[k] = v
		}
	}

	return u
}
