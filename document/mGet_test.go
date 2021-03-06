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

package document_test

import (
	"encoding/json"
	"testing"

	"github.com/kuzzleio/sdk-go/document"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func TestMGetIndexNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	d := document.NewDocument(k)
	var ids []string
	ids = append(ids, "id1")
	_, err := d.MGet("", "collection", ids, nil)
	assert.NotNil(t, err)
}

func TestMGetCollectionNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	d := document.NewDocument(k)
	var ids []string
	ids = append(ids, "id1")
	_, err := d.MGet("index", "", ids, nil)
	assert.NotNil(t, err)
}

func TestMGetIdsNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	d := document.NewDocument(k)
	var ids []string
	_, err := d.MGet("index", "collection", ids, nil)
	assert.NotNil(t, err)
}

func TestMGetDocumentError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: types.NewError("Unit test error")}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	d := document.NewDocument(k)
	var ids []string
	ids = append(ids, "id1")
	_, err := d.MGet("index", "collection", ids, nil)
	assert.NotNil(t, err)
	assert.Equal(t, "Unit test error", err.(types.KuzzleError).Message)
}

func TestMGetDocument(t *testing.T) {

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "document", parsedQuery.Controller)
			assert.Equal(t, "mGet", parsedQuery.Action)
			assert.Equal(t, "index", parsedQuery.Index)
			assert.Equal(t, "collection", parsedQuery.Collection)

			return &types.KuzzleResponse{Result: []byte(`
			{
				"successes": [
					{
						"_id": "id1",
						"_index": "index",
						"_shards": {
							"failed": 0,
							"successful": 1,
							"total": 2
						},
						"_source": {
							"document": "body"
						},
						"_meta": {
							"active": true,
							"author": "-1",
							"createdAt": 1484225532686,
							"deletedAt": null,
							"updatedAt": null,
							"updater": null
						},
						"_type": "collection",
						"_version": 1,
						"created": true,
						"result": "created"
					}
				],
				"errors": []
			}`),
			}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	d := document.NewDocument(k)
	var ids []string
	ids = append(ids, "id1")
	_, err := d.MGet("index", "collection", ids, nil)
	assert.Nil(t, err)
}
