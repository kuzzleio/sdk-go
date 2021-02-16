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

package collection_test

import (
	"testing"

	"github.com/kuzzleio/sdk-go/collection"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func TestRefreshIndexNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	nc := collection.NewCollection(k)
	err := nc.Refresh("", "collection", nil)
	assert.NotNil(t, err)
}

func TestRefreshCollectionNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	nc := collection.NewCollection(k)
	err := nc.Refresh("index", "", nil)
	assert.NotNil(t, err)
}
func TestRefreshError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: types.NewError("Unit test error")}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	nc := collection.NewCollection(k)
	err := nc.Refresh("index", "collection", nil)
	assert.NotNil(t, err)
	assert.Equal(t, "Unit test error", err.(types.KuzzleError).Message)
}

func TestRefresh(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Result: []byte(`{
				"status":200,
				"error": null,
				"index": "index",
				"collection": "collection",
				"action": "refresh",
				"result": null
			}`)}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	nc := collection.NewCollection(k)
	err := nc.Refresh("index", "collection", nil)
	assert.Nil(t, err)
}
