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

package server_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/protocol/websocket"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func TestNowQueryError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: types.KuzzleError{Message: "error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	k.Connect()
	_, err := k.Server.Now(nil)
	assert.NotNil(t, err)
}

func TestNow(t *testing.T) {
	type now struct {
		Now int64 `json:"now"`
	}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			n := now{Now: int64(1488540242465)}

			marsh, _ := json.Marshal(n)
			return &types.KuzzleResponse{Result: marsh}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := k.Server.Now(nil)
	assert.Equal(t, int64(1488540242465), res)
}

func ExampleKuzzle_Now() {
	conn := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(conn, nil)

	now, err := k.Server.Now(nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(now)
}
