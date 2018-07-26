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

func TestMDeleteIndexNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	d := document.NewDocument(k)

	var ids []string
	ids = append(ids, "id1")
	_, err := d.MDelete("", "collection", ids, nil)
	assert.NotNil(t, err)
}

func TestMDeleteCollectionNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	d := document.NewDocument(k)

	var ids []string
	ids = append(ids, "id1")
	_, err := d.MDelete("index", "", ids, nil)
	assert.NotNil(t, err)
}

func TestMDeleteIdsNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	d := document.NewDocument(k)

	var ids []string
	_, err := d.MDelete("index", "collection", ids, nil)
	assert.NotNil(t, err)
}

func TestMDeleteDocumentError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: types.NewError("Unit test error")}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	d := document.NewDocument(k)

	var ids []string
	ids = append(ids, "id1")
	_, err := d.MDelete("index", "collection", ids, nil)
	assert.NotNil(t, err)
	assert.Equal(t, "Unit test error", err.(types.KuzzleError).Message)
}

func TestMDeleteDocument(t *testing.T) {

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "document", parsedQuery.Controller)
			assert.Equal(t, "mDelete", parsedQuery.Action)
			assert.Equal(t, "index", parsedQuery.Index)
			assert.Equal(t, "collection", parsedQuery.Collection)

			return &types.KuzzleResponse{Result: []byte(`
			{
				"ids": ["id1", "id2"]
			}`),
			}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	d := document.NewDocument(k)

	var ids []string
	ids = append(ids, "id1")
	_, err := d.MDelete("index", "collection", ids, nil)
	assert.Nil(t, err)
}
