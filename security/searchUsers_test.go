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

	"github.com/kuzzleio/sdk-go/security"

	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/protocol/websocket"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func TestSearchUsersError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: types.NewError("Test error")}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	_, err := k.Security.SearchUsers(nil, nil)
	assert.NotNil(t, err)
}

func TestSearchUsers(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "searchUsers", parsedQuery.Action)

			return &types.KuzzleResponse{Result: []byte(`{
				"total": 42,
				"hits": [
					{"_id": "user42", "_source": {"profileIds": ["admin", "other"], "foo": "bar"}}
				]
			}`)}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	res, _ := k.Security.SearchUsers(nil, nil)

	assert.Equal(t, 42, res.Total)
	assert.Equal(t, 1, len(res.Hits))
	assert.Equal(t, res.Hits[0].Id, "user42")
	assert.Equal(t, []string{"admin", "other"}, res.Hits[0].ProfileIds)
	assert.Equal(t, map[string]interface{}{"foo": "bar"}, res.Hits[0].Content)
}

func TestSearchUsersNext(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Result: json.RawMessage(`{
				"hits": [
					{
						"_id": "id1",
						"_source": {
							"profileIds": [ "admin", "other" ],
							"foo": "bar"
						}
					},
					{
						"_id": "id2",
						"_source": {
							"profileIds": [ "default" ],
							"foo": "baz"
						}
					}
				],
				"total": 42,
				"_scroll_id": "scroll_id"
			}`)}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	options := types.NewQueryOptions()
	options.SetScroll("1m")

	sr, err := k.Security.SearchUsers(json.RawMessage(`{}`), options)
	assert.Nil(t, err)
	assert.NotNil(t, sr)

	c.MockSend = func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
		parsedQuery := &types.KuzzleRequest{}
		json.Unmarshal(query, parsedQuery)

		assert.Equal(t, "1m", options.Scroll())
		assert.Equal(t, "scroll_id", options.ScrollId())

		return &types.KuzzleResponse{Result: json.RawMessage(`{
			"hits": [
				{
					"_id": "id3",
					"_source": {
						"profileIds": [ "admin", "other" ],
						"foo": "bar"
					}
				},
				{
					"_id": "id4",
					"_source": {
						"profileIds": [ "default" ],
						"foo": "baz"
					}
				}
			],
			"total": 42,
			"_scroll_id": "new_scroll"
		}`)}
	}

	nsr, err := sr.Next()
	assert.Nil(t, err)
	assert.NotNil(t, nsr)
	assert.Equal(t, 4, nsr.Fetched)

	u := security.NewUser("id", nil)
	assert.IsType(t, u, nsr.Hits[1])
}

func ExampleSearchUsers() {
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	res, err := k.Security.SearchUsers(nil, nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res.Hits[0].Id, res.Hits[0].Content)
}
