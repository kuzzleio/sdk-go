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

	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/protocol/websocket"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func TestSearchProfileError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{
				Error: types.NewError("Test error"),
			}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	_, err := k.Security.SearchProfiles(nil, nil)
	assert.NotNil(t, err)
}

func TestSearchProfile(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "searchProfiles", parsedQuery.Action)

			return &types.KuzzleResponse{Result: []byte(`{
				"total": 42,
				"hits": [
					{ "_id": "profile42", "_source": { "policies": [{"roleId": "admin"}] } }
				]
			}`)}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	res, _ := k.Security.SearchProfiles(nil, nil)

	assert.Equal(t, 42, res.Total)
	assert.Equal(t, "profile42", res.Hits[0].Id)
	assert.Equal(t, []*types.Policy{
		{RoleId: "admin"},
	}, res.Hits[0].Policies)
}

func ExampleSearchProfiles() {
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	res, err := k.Security.SearchProfiles(nil, nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res.Hits[0].Id, res.Hits[0].Policies)
}

func TestSearchProfileWithScroll(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "searchProfiles", parsedQuery.Action)

			return &types.KuzzleResponse{Result: []byte(`{
				"total": 42,
				"hits": [
					{"_id": "profile42", "_source": {"policies": [{"roleId": "admin"}]}}
				],
				"scrollId": "f00b4r"
			}`)}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	opts := types.NewQueryOptions()
	opts.SetFrom(2)
	opts.SetSize(4)
	opts.SetScroll("1m")

	res, _ := k.Security.SearchProfiles(nil, opts)
	assert.Equal(t, 42, res.Total)
	assert.Equal(t, 1, len(res.Hits))
	assert.Equal(t, []*types.Policy{
		{RoleId: "admin"},
	}, res.Hits[0].Policies)
	assert.Equal(t, "f00b4r", res.ScrollId)
	assert.Equal(t, "profile42", res.Hits[0].Id)
}
