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
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/protocol/websocket"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestZrangebyscoreError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: types.KuzzleError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := k.MemoryStorage.Zrangebyscore("foo", 1, 6, nil)

	assert.NotNil(t, err)
}

func TestZrangebyscore(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "ms", parsedQuery.Controller)
			assert.Equal(t, "zrangebyscore", parsedQuery.Action)

			r, _ := json.Marshal([]string{"bar", "5", "foo", "1.377"})
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := k.MemoryStorage.Zrangebyscore("foo", 1, 6, nil)

	expectedResult := []*types.MSSortedSet{
		{Member: "bar", Score: 5},
		{Member: "foo", Score: 1.377},
	}

	assert.Equal(t, expectedResult, res)
}

func TestZrangebyscoreWithLimits(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "ms", parsedQuery.Controller)
			assert.Equal(t, "zrangebyscore", parsedQuery.Action)
			assert.Equal(t, []interface{}([]interface{}{"withscores"}), parsedQuery.Options)
			assert.Equal(t, "0,1", parsedQuery.Limit)

			r, _ := json.Marshal([]string{"bar", "5"})
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	qo := types.NewQueryOptions()

	qo.SetLimit([]int{0, 1})
	res, _ := k.MemoryStorage.Zrangebyscore("foo", 1, 6, qo)

	expectedResult := []*types.MSSortedSet{
		{Member: "bar", Score: 5},
	}

	assert.Equal(t, expectedResult, res)
}

func ExampleMs_Zrangebyscore() {
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, err := k.MemoryStorage.Zrangebyscore("foo", 1, 6, nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res[0].Member, res[0].Score)
}
