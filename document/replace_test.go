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

func TestReplaceIndexNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	d := document.NewDocument(k)

	_, err := d.Replace("", "collection", "id1", json.RawMessage(`{"foo":"bar"}`), nil)
	assert.NotNil(t, err)
}

func TestReplaceCollectionNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	d := document.NewDocument(k)

	_, err := d.Replace("index", "", "id1", json.RawMessage(`{"foo":"bar"}`), nil)
	assert.NotNil(t, err)
}

func TestReplaceIdNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	d := document.NewDocument(k)

	_, err := d.Replace("index", "collection", "", json.RawMessage(`{"foo":"bar"}`), nil)
	assert.NotNil(t, err)
}

func TestReplaceBodyNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	d := document.NewDocument(k)

	_, err := d.Replace("index", "collection", "id1", nil, nil)
	assert.NotNil(t, err)
}

func TestReplaceDocumentError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: types.NewError("Unit test error")}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	d := document.NewDocument(k)

	_, err := d.Replace("index", "collection", "id1", json.RawMessage(`{"foo":"bar"}`), nil)
	assert.NotNil(t, err)
	assert.Equal(t, "Unit test error", err.(types.KuzzleError).Message)
}

func TestReplaceDocument(t *testing.T) {
	id := "myId"

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "document", parsedQuery.Controller)
			assert.Equal(t, "replace", parsedQuery.Action)
			assert.Equal(t, "index", parsedQuery.Index)
			assert.Equal(t, "collection", parsedQuery.Collection)
			assert.Equal(t, id, parsedQuery.Id)

			return &types.KuzzleResponse{Result: []byte(`
				{
					"_id": "<documentId>",
					"_version": "<number>",// The new version number of this document
					"created": false
				}`),
			}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	d := document.NewDocument(k)

	_, err := d.Replace("index", "collection", id, json.RawMessage(`{"foo":"bar"}`), nil)
	assert.Nil(t, err)
}
