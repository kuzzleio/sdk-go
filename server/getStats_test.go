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

package server_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/kuzzleio/sdk-go/connection/websocket"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func TestGetStatsQueryError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)
			assert.Equal(t, "server", request.Controller)
			assert.Equal(t, "getLastStats", request.Action)
			return &types.KuzzleResponse{Error: types.KuzzleError{Message: "error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	_, err := k.Server.GetStats(nil, nil, nil)
	assert.NotNil(t, err)
}

func TestGetStats(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)
			assert.Equal(t, "server", request.Controller)
			assert.Equal(t, "getLastStats", request.Action)

			return &types.KuzzleResponse{Result: json.RawMessage(`{"foo": "bar"}`)}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := k.Server.GetStats(nil, nil, nil)

	assert.Equal(t, json.RawMessage(`{"foo": "bar"}`), res)
}

func ExampleKuzzle_GetStats() {
	conn := websocket.NewWebSocket("localhost", nil)
	k, _ := kuzzle.NewKuzzle(conn, nil)

	now := time.Now()
	res, err := k.Server.GetStats(&now, nil, nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res)
}
