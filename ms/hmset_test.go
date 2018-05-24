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

package ms_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/kuzzleio/sdk-go/connection/websocket"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func TestHmsetEmptyEntries(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)

	err := k.MemoryStorage.Hmset("", []*types.MsHashField{}, nil)

	assert.NotNil(t, err)
	assert.Equal(t, "[400] Ms.Hmset: at least one entry field to set is required", fmt.Sprint(err))
}

func TestHmsetError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: types.KuzzleError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	err := k.MemoryStorage.Hmset("foo", []*types.MsHashField{}, nil)

	assert.NotNil(t, err)
}

func TestHmset(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "ms", parsedQuery.Controller)
			assert.Equal(t, "hmset", parsedQuery.Action)
			assert.Equal(t, "foo", parsedQuery.Id)
			assert.Equal(t, "foo", parsedQuery.Body.(map[string]interface{})["entries"].([]interface{})[0].(map[string]interface{})["field"].(string))
			assert.Equal(t, "bar", parsedQuery.Body.(map[string]interface{})["entries"].([]interface{})[0].(map[string]interface{})["value"].(string))

			r, _ := json.Marshal("result")
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	err := k.MemoryStorage.Hmset("foo", []*types.MsHashField{{Field: "foo", Value: "bar"}}, nil)

	assert.Nil(t, err)
}

func ExampleMs_Hmset() {
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)

	err := k.MemoryStorage.Hmset("foo", []*types.MsHashField{{Field: "foo", Value: "bar"}}, nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("success")
}
