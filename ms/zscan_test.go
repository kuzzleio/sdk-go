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

package ms_test

import (
	"encoding/json"
	"fmt"
	"github.com/kuzzleio/sdk-go/connection/websocket"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestZscanError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: types.KuzzleError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	_, err := k.MemoryStorage.Zscan("foo", 42, nil)

	assert.NotNil(t, err)
}

func TestZscan(t *testing.T) {
	scanResponse := types.MSScanResponse{
		Cursor: 42,
		Values: []string{"bar", "5", "foo", "10"},
	}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "ms", parsedQuery.Controller)
			assert.Equal(t, "zscan", parsedQuery.Action)

			r, _ := json.Marshal([]interface{}{"42", []string{"bar", "5", "foo", "10"}})
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	res, _ := k.MemoryStorage.Zscan("foo", 42, nil)

	assert.Equal(t, &scanResponse, res)
}

func TestZscanWithOptions(t *testing.T) {
	scanResponse := types.MSScanResponse{
		Cursor: 42,
		Values: []string{"bar", "5", "foo", "10"},
	}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "ms", parsedQuery.Controller)
			assert.Equal(t, "zscan", parsedQuery.Action)
			assert.Equal(t, 10, options.Count())
			assert.Equal(t, "*", options.Match())

			r, _ := json.Marshal([]interface{}{"42", []string{"bar", "5", "foo", "10"}})
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	qo := types.NewQueryOptions()

	qo.SetCount(10)
	qo.SetMatch("*")

	res, _ := k.MemoryStorage.Zscan("foo", 42, qo)

	assert.Equal(t, &scanResponse, res)
}

func ExampleMs_Zscan() {
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	qo := types.NewQueryOptions()

	qo.SetCount(10)
	qo.SetMatch("*")

	res, err := k.MemoryStorage.Zscan("foo", 42, qo)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res.Cursor, res.Values)
}
