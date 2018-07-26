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

package security_test

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

func TestMGetRolesEmptyId(t *testing.T) {
	c := &internal.MockedConnection{}
	k, _ := kuzzle.NewKuzzle(c, nil)
	_, err := k.Security.MGetRoles([]string{}, nil)

	assert.NotNil(t, err)
}

func TestMGetRolesError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{
				Error: types.NewError("Test error"),
			}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	_, err := k.Security.MGetRoles([]string{"id"}, nil)
	assert.NotNil(t, err)
}

func TestMGetRoles(t *testing.T) {

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "mGetRoles", parsedQuery.Action)

			return &types.KuzzleResponse{Result: []byte(`{
           "_shards": {
             "failed": 0,
             "successful": 5,
             "total": 5
           },
           "hits": [
             {
               "_id": "id",
               "_index": "%kuzzle",
               "_score": 1,
               "_source": {
                 "controllers": {}
               },
               "_type": "roles"
             }
           ],
           "max_score": null,
           "timed_out": false,
           "took": 1,
           "total": 2
        }`),
			}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	res, err := k.Security.MGetRoles([]string{"id"}, nil)
	assert.NoError(t, err)
	assert.Equal(t, "id", res[0].Id)

}

func ExampleMGetRoles() {
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	res, err := k.Security.MGetRoles([]string{"id"}, nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res[0])
}
