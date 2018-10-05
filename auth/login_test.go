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

package auth_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/kuzzleio/sdk-go/event"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/protocol/websocket"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func TestLoginNoStrategy(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	_, err := k.Auth.Login("", json.RawMessage("{}"), nil)
	assert.NotNil(t, err)
}

func TestLoginError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)

			assert.Equal(t, "auth", request.Controller)
			assert.Equal(t, "login", request.Action)
			return &types.KuzzleResponse{Error: types.KuzzleError{Message: "error"}}
		},
		MockEmitEvent: func(e int, arg interface{}) {
			assert.Equal(t, event.LoginAttempt, e)
			assert.Equal(t, "error", arg.(*types.LoginAttempt).Error.Error())
		},
	}

	k, _ := kuzzle.NewKuzzle(c, nil)
	k.Auth.Login("local", json.RawMessage("{}"), nil)
}

func TestLogin(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)

			assert.Equal(t, "auth", request.Controller)
			assert.Equal(t, "login", request.Action)
			assert.Equal(t, 42, request.ExpiresIn)

			type loginResult struct {
				Jwt string `json:"jwt"`
			}

			loginRes := loginResult{"token"}
			marsh, _ := json.Marshal(loginRes)

			return &types.KuzzleResponse{Result: marsh}
		},
		MockEmitEvent: func(e int, arg interface{}) {
			assert.Equal(t, event.LoginAttempt, e)
			assert.Equal(t, true, arg.(*types.LoginAttempt).Success)
			assert.Nil(t, arg.(*types.LoginAttempt).Error)
		},
	}

	k, _ := kuzzle.NewKuzzle(c, nil)
	expiresIn := 42
	token, _ := k.Auth.Login("local", json.RawMessage("{}"), &expiresIn)
	assert.Equal(t, "token", token)
}

func ExampleKuzzle_Login() {
	conn := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(conn, nil)

	type credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	myCredentials := credentials{"foo", "bar"}
	marsh, _ := json.Marshal(myCredentials)

	jwt, err := k.Auth.Login("local", marsh, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(jwt)
}
