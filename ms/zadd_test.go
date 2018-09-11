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

func TestZaddEmptyElements(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)

	_, err := k.MemoryStorage.Zadd("foo", []*types.MSSortedSet{}, nil)

	assert.NotNil(t, err)
	assert.Equal(t, "[400] Ms.Zadd: please provide at least one element", fmt.Sprint(err))
}

func TestZaddError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: types.KuzzleError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	sortedSet := []*types.MSSortedSet{
		{Score: 10, Member: "bar"},
		{Score: 5, Member: "foo"},
	}

	_, err := k.MemoryStorage.Zadd("foo", sortedSet, nil)

	assert.NotNil(t, err)
}

func TestZadd(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "ms", parsedQuery.Controller)
			assert.Equal(t, "zadd", parsedQuery.Action)

			r, _ := json.Marshal(2)
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	sortedSet := []*types.MSSortedSet{
		{Score: 10, Member: "bar"},
		{Score: 5, Member: "foo"},
	}

	res, _ := k.MemoryStorage.Zadd("foo", sortedSet, nil)

	assert.Equal(t, 2, res)
}

func TestZaddWithOptions(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "ms", parsedQuery.Controller)
			assert.Equal(t, "zadd", parsedQuery.Action)
			assert.Equal(t, true, options.Ch())
			assert.Equal(t, true, options.Incr())
			assert.Equal(t, true, options.Nx())
			assert.Equal(t, true, options.Xx())

			r, _ := json.Marshal(2)
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	qo := types.NewQueryOptions()

	sortedSet := []*types.MSSortedSet{
		{Score: 10, Member: "bar"},
		{Score: 5, Member: "foo"},
	}

	qo.SetCh(true)
	qo.SetIncr(true)
	qo.SetNx(true)
	qo.SetXx(true)
	res, _ := k.MemoryStorage.Zadd("foo", sortedSet, qo)

	assert.Equal(t, 2, res)
}

func ExampleMs_Zadd() {
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	qo := types.NewQueryOptions()

	sortedSet := []*types.MSSortedSet{
		{Score: 10, Member: "bar"},
		{Score: 5, Member: "foo"},
	}

	qo.SetCh(true)
	qo.SetIncr(true)
	qo.SetNx(true)
	qo.SetXx(true)

	res, err := k.MemoryStorage.Zadd("foo", sortedSet, qo)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res)
}
