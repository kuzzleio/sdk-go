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

func TestUpdateSelfQueryError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: types.KuzzleError{Message: "error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	_, err := k.Auth.UpdateSelf(nil, nil)
	assert.NotNil(t, err)
}

func TestUpdateSelf(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)
			assert.Equal(t, "auth", request.Controller)
			assert.Equal(t, "updateSelf", request.Action)

			return &types.KuzzleResponse{Result: []byte(`{
				"_id": "login",
				"_source": {
					"username": "foo"
				}
			}`)}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	res, _ := k.Auth.UpdateSelf(json.RawMessage("{\"foo\":\"bar\"}"), nil)

	assert.Equal(t, "login", res.Id)
}

func ExampleKuzzle_UpdateSelf() {
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

	res, _ := k.Auth.UpdateSelf(json.RawMessage("{\"foo\":\"bar\"}"), nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res)
}
