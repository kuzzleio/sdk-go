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

	"github.com/kuzzleio/sdk-go/protocol/websocket"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func TestUpdateMyCredentialsQueryError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)
			assert.Equal(t, "auth", request.Controller)
			assert.Equal(t, "updateMyCredentials", request.Action)
			return &types.KuzzleResponse{Error: types.KuzzleError{Message: "error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	_, err := k.Auth.UpdateMyCredentials("local", nil, nil)
	assert.NotNil(t, err)
}

func TestUpdateMyCredentialsEmptyStrategy(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: types.KuzzleError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	_, err := k.Auth.UpdateMyCredentials("", nil, nil)
	assert.NotNil(t, err)
}

func TestUpdateMyCredentials(t *testing.T) {
	type credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	myCredentials := credentials{"foo", "bar"}
	marsh, _ := json.Marshal(myCredentials)

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			ack := credentials{Username: "foo", Password: "bar"}
			r, _ := json.Marshal(ack)

			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)
			assert.Equal(t, "auth", request.Controller)
			assert.Equal(t, "updateMyCredentials", request.Action)
			assert.Equal(t, "local", request.Strategy)
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := k.Auth.UpdateMyCredentials("local", marsh, nil)
	creds := credentials{}
	json.Unmarshal(res, &creds)

	assert.Equal(t, "foo", creds.Username)
	assert.Equal(t, "bar", creds.Password)
}

func ExampleKuzzle_UpdateMyCredentials() {
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

	newCredentials := credentials{"foo", "new"}
	marsh, _ = json.Marshal(newCredentials)

	res, err := k.Auth.UpdateMyCredentials("local", marsh, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res)
}
