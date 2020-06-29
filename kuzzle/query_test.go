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

package kuzzle_test

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

func TestQueryDefaultOptions(t *testing.T) {
	var k *kuzzle.Kuzzle

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)

			assert.NotNil(t, request.Volatile)
			assert.Equal(t, 0, request.From)
			assert.Equal(t, 10, request.Size)
			assert.Equal(t, "", request.Scroll)
			assert.Equal(t, "", request.ScrollId)

			return &types.KuzzleResponse{}
		},
	}
	k, _ = kuzzle.NewKuzzle(c, nil)

	ch := make(chan *types.KuzzleResponse)
	options := types.NewQueryOptions()
	go k.Query(&types.KuzzleRequest{}, options, ch)
	<-ch
}

func TestQueryWithOptions(t *testing.T) {
	var k *kuzzle.Kuzzle

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)

			assert.NotNil(t, request.Volatile)
			assert.Equal(t, 42, request.From)
			assert.Equal(t, 24, request.Size)
			assert.Equal(t, "5m", request.Scroll)
			assert.Equal(t, "f00b4r", request.ScrollId)

			rawRequest := map[string]interface{}{}
			json.Unmarshal(query, &rawRequest)

			assert.Equal(t, "wait_for", rawRequest["refresh"])
			assert.Equal(t, "wait_for", rawRequest["refresh"])
			assert.Equal(t, 7.0, rawRequest["retryOnConflict"])

			return &types.KuzzleResponse{}
		},
	}
	k, _ = kuzzle.NewKuzzle(c, nil)

	ch := make(chan *types.KuzzleResponse)
	options := types.NewQueryOptions()

	options.SetFrom(42)
	options.SetSize(24)
	options.SetScroll("5m")
	options.SetScrollId("f00b4r")
	options.SetRefresh("wait_for")
	options.SetRetryOnConflict(7)
	query := types.KuzzleRequest{}
	query.AddCustomArg("cert", "cert")
	query.AddCustomArg("foo", "bar")

	go k.Query(&query, options, ch)
	<-ch
}

func TestQueryWithVolatile(t *testing.T) {
	var k *kuzzle.Kuzzle
	var volatileData = types.VolatileData(`{"modifiedBy":"awesome me","reason":"it needed to be modified","sdkName":"go@3.0.0"}`)

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)

			fmt.Printf("%s %s\n", request.Volatile, volatileData)
			assert.Equal(t, volatileData, request.Volatile)
			assert.NotNil(t, request.Volatile)

			return &types.KuzzleResponse{}
		},
	}
	k, _ = kuzzle.NewKuzzle(c, nil)

	ch := make(chan *types.KuzzleResponse)
	options := types.NewQueryOptions()
	options.SetVolatile(volatileData)
	go k.Query(&types.KuzzleRequest{}, options, ch)
	<-ch
}

func ExampleKuzzle_Query() {
	conn := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(conn, nil)

	request := types.KuzzleRequest{Controller: "server", Action: "now"}
	resChan := make(chan *types.KuzzleResponse)
	k.Query(&request, nil, resChan)

	now := <-resChan
	if now.Error.Message != "" {
		fmt.Println(now.Error.Message)
		return
	}

	fmt.Println(now.Result)
}
