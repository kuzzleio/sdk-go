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
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/protocol/websocket"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlushQueue(t *testing.T) {
	c := &internal.MockedConnection{}
	k, _ := kuzzle.NewKuzzle(c, nil)

	c.AddToOfflineQueue(&types.QueryObject{RequestId: "test"})
	assert.NotEmpty(t, k.OfflineQueue())

	k.FlushQueue()
	assert.Empty(t, k.OfflineQueue())
}

func ExampleKuzzle_FlushQueue() {
	conn := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(conn, nil)

	k.FlushQueue()
}
