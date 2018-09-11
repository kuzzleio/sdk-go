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

// GetCurrentUser retrieves user linked to JWT
func (a *Auth) GetCurrentUser() (*security.User, error) {
	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "auth",
		Action:     "getCurrentUser",
	}

	go a.kuzzle.Query(query, nil, result)

	res := <-result

	if res.Error.Error() != "" {
		return nil, res.Error
	}

	type jsonUser struct {
		Id     string          `json:"_id"`
		Source json.RawMessage `json:"_source"`
	}
	ju := &jsonUser{}
	json.Unmarshal(res.Result, ju)

	var unmarsh map[string]interface{}
	json.Unmarshal(ju.Source, &unmarsh)
	u := &security.User{
		Id:      ju.Id,
		Content: unmarsh,
	}
	return u, nil
}
