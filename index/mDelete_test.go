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

package index_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/kuzzleio/sdk-go/index"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/protocol/websocket"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func TestMDeleteNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	i := index.NewIndex(k)
	indexes := []string{}
	_, err := i.MDelete(indexes, nil)
	assert.NotNil(t, err)
}

func TestMDeleteQueryError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: types.KuzzleError{Message: "error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	i := index.NewIndex(k)
	indexes := []string{"index"}
	_, err := i.MDelete(indexes, nil)
	assert.NotNil(t, err)
}

func TestMDelete(t *testing.T) {
	type ackResult struct {
		Acknowledged       bool
		ShardsAcknowledged bool
	}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			q := &types.KuzzleRequest{}
			json.Unmarshal(query, q)

			assert.Equal(t, "index", q.Controller)
			assert.Equal(t, "mDelete", q.Action)

			return &types.KuzzleResponse{Result: []byte(`{"deleted":["index1"]}`)}
		},
	}

	k, _ := kuzzle.NewKuzzle(c, nil)
	i := index.NewIndex(k)
	indexes := []string{"index1"}
	res, err := i.MDelete(indexes, nil)

	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, indexes, res)
}

func ExampleIndex_MDelete() {
	conn := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(conn, nil)
	i := index.NewIndex(k)
	i.Create("index1", nil)
	i.Create("index2", nil)
	indexes := []string{
		"index1",
		"index2",
	}
	_, err := i.MDelete(indexes, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
