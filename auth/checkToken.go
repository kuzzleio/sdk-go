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

	"github.com/kuzzleio/sdk-go/types"
)

// TokenValidity provides a representation for JWT validity
type TokenValidity struct {
	Valid     bool   `json:"valid"`
	State     string `json:"state"`
	ExpiresAt int    `json:"expiresAt"`
}

// CheckToken checks the validity of a JSON Web Token.
func (a *Auth) CheckToken(token string) (*TokenValidity, error) {
	if token == "" {
		return nil, types.NewError("Kuzzle.CheckToken: token required", 400)
	}

	result := make(chan *types.KuzzleResponse)

	type body struct {
		Token string `json:"token"`
	}

	query := &types.KuzzleRequest{
		Controller: "auth",
		Action:     "checkToken",
		Body:       &body{token},
	}
	go a.kuzzle.Query(query, nil, result)

	res := <-result

	if res.Error.Error() != "" {
		return nil, res.Error
	}

	tokenValidity := &TokenValidity{}
	json.Unmarshal(res.Result, tokenValidity)

	return tokenValidity, nil
}
