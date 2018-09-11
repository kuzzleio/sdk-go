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

	"github.com/kuzzleio/sdk-go/event"
	"github.com/kuzzleio/sdk-go/types"
)

// Login send login request to kuzzle with credentials.
// If login success, store the jwtToken into kuzzle object.
func (a *Auth) Login(strategy string, credentials json.RawMessage, expiresIn *int) (string, error) {
	if strategy == "" {
		return "", types.NewError("Kuzzle.Login: cannot authenticate to Kuzzle without an authentication strategy", 400)
	}

	type loginResult struct {
		Jwt string `json:"jwt"`
	}

	var token loginResult
	var body interface{}

	if credentials != nil {
		body = credentials
	}

	q := &types.KuzzleRequest{
		Controller: "auth",
		Action:     "login",
		Body:       body,
		Strategy:   strategy,
	}

	if expiresIn != nil {
		q.ExpiresIn = *expiresIn
	}

	result := make(chan *types.KuzzleResponse)

	go a.kuzzle.Query(q, nil, result)

	res := <-result

	json.Unmarshal(res.Result, &token)

	if res.Error.Error() != "" {
		a.kuzzle.EmitEvent(event.LoginAttempt, &types.LoginAttempt{Success: false, Error: res.Error})
		return "", res.Error
	}

	a.kuzzle.SetJwt(token.Jwt)

	return token.Jwt, nil
}
