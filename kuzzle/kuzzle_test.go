// Copyright 2015-2017 Kuzzle
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

package kuzzle_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/kuzzleio/sdk-go/connection/websocket"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func ExampleKuzzle_Connect() {
	opts := types.NewOptions()
	opts.SetConnect(types.Manual)

	conn := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(conn, opts)

	err := k.Connect()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

}

func ExampleKuzzle_Jwt() {
	conn := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(conn, nil)

	jwt := k.Jwt()
	fmt.Println(jwt)
}

func ExampleKuzzle_OfflineQueue() {
	//todo
}

func TestUnsetJwt(t *testing.T) {
	var k *kuzzle.Kuzzle

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)

			assert.Equal(t, "auth", request.Controller)
			assert.Equal(t, "login", request.Action)
			assert.Equal(t, 0, request.ExpiresIn)

			type loginResult struct {
				Jwt string `json:"jwt"`
			}

			loginRes := loginResult{"token"}
			marsh, _ := json.Marshal(loginRes)

			return &types.KuzzleResponse{Result: marsh}
		},
	}

	k, _ = kuzzle.NewKuzzle(c, nil)

	res, _ := k.Auth.Login("local", nil, nil)
	assert.Equal(t, "token", res)
	assert.Equal(t, "token", k.Jwt())
	k.UnsetJwt()
	assert.Equal(t, "", k.Jwt())
}

func ExampleKuzzle_UnsetJwt() {
	conn := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(conn, nil)

	type credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	myCredentials := credentials{"foo", "bar"}
	marsh, _ := json.Marshal(myCredentials)

	_, err := k.Auth.Login("local", marsh, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	k.UnsetJwt()
	fmt.Println(k.Jwt())
}

func TestSetDefaultIndexNullIndex(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	assert.NotNil(t, k.SetDefaultIndex(""))
}

func TestSetDefaultIndex(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)
			assert.Equal(t, "myindex", request.Index)
			return &types.KuzzleResponse{Error: types.KuzzleError{Message: "error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	k.SetDefaultIndex("myindex")
}

func ExampleKuzzle_SetDefaultIndex() {
	conn := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(conn, nil)

	err := k.SetDefaultIndex("index")

	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
