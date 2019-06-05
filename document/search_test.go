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

func TestSearchIndexNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	d := document.NewDocument(k)

	_, err := d.Search("", "collection", json.RawMessage(`{"foo":"bar"}`), nil)
	assert.NotNil(t, err)
}

func TestSearchCollectionNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	d := document.NewDocument(k)

	_, err := d.Search("index", "", json.RawMessage(`{"foo":"bar"}`), nil)
	assert.NotNil(t, err)
}

func TestSearchBodyNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	d := document.NewDocument(k)

	_, err := d.Search("index", "collection", nil, nil)
	assert.NotNil(t, err)
}

func TestSearchDocumentError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: types.NewError("Unit test error")}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	d := document.NewDocument(k)

	_, err := d.Search("index", "collection", json.RawMessage(`{"foo":"bar"}`), nil)
	assert.NotNil(t, err)
	assert.Equal(t, "Unit test error", err.(types.KuzzleError).Message)
}

func TestSearchDocument(t *testing.T) {

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "document", parsedQuery.Controller)
			assert.Equal(t, "search", parsedQuery.Action)
			assert.Equal(t, "index", parsedQuery.Index)
			assert.Equal(t, "collection", parsedQuery.Collection)

			return &types.KuzzleResponse{Result: []byte(`
			{
				"hits": ["id1", "id2"],
				"total": 42
			}`),
			}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	d := document.NewDocument(k)

	response, err := d.Search("index", "collection", json.RawMessage(`{"foo":"bar"}`), nil)
	assert.Nil(t, err)
	assert.Equal(t, 2, response.Fetched)
	assert.Equal(t, 42, response.Total)
}

func TestSearchDocumentNextFromSize(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Result: json.RawMessage(`{
				"hits": ["id1", "id2"],
				"total": 42
			}`)}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	options := types.NewQueryOptions()
	options.SetFrom(0)
	options.SetSize(2)

	sr, err := k.Document.Search("index", "collection", json.RawMessage(`{}`), options)
	assert.Nil(t, err)
	assert.Equal(t, 42, sr.Total)
	assert.Equal(t, 2, sr.Fetched)

	c.MockSend = func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
		parsedQuery := &types.KuzzleRequest{}
		json.Unmarshal(query, parsedQuery)

		assert.Equal(t, 2, parsedQuery.From)
		assert.Equal(t, 2, parsedQuery.Size)

		return &types.KuzzleResponse{Result: json.RawMessage(`{
			"hits": ["id3"]
		}`)}
	}

	nsr, err := sr.Next()
	assert.Nil(t, err)
	assert.Equal(t, 3, nsr.Fetched)
}

func TestSearchDocumentNextEnd(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Result: json.RawMessage(`{
				"hits": ["id1", "id2"],
				"total": 2
			}`)}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	options := types.NewQueryOptions()
	options.SetFrom(0)
	options.SetSize(2)

	sr, err := k.Document.Search("index", "collection", json.RawMessage(`{}`), options)
	assert.Nil(t, err)

	nsr, err := sr.Next()

	assert.Nil(t, err)
	assert.Nil(t, nsr)
}

func TestSearchDocumentNextSort(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Result: json.RawMessage(`{
				"hits": [
					{
						"_id": "id1",
						"_source": {
							"foo": {
								"bar": 3
							}
						}
					},
					{
						"_id": "id2",
						"_source": {
							"foo": {
								"bar": 12
							}
						}
					}
				],
				"total": 42
			}`)}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	options := types.NewQueryOptions()
	options.SetSize(2)

	sr, err := k.Document.Search("index", "collection", json.RawMessage(`{
		"sort":[ 
			{"_uid": "asc"},
			{ "foo.bar": "desc" }
		]
	}`), options)
	assert.Nil(t, err)

	c.MockSend = func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
		parsedQuery := &types.KuzzleRequest{}
		json.Unmarshal(query, parsedQuery)

		assert.Equal(t, []interface{}{"id2", float64(12)}, parsedQuery.Body.(map[string]interface{})["search_after"])

		return &types.KuzzleResponse{Result: json.RawMessage(`{
			"hits": ["id3", "id4"],
			"total": 42
		}`)}
	}

	nsr, err := sr.Next()
	assert.Nil(t, err)
	assert.Equal(t, 4, nsr.Fetched)
	assert.NotNil(t, nsr)
}

func TestSearchDocumentNextScroll(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Result: json.RawMessage(`{
				"hits": ["id1", "id2"],
				"total": 42,
				"_scroll_id": "scroll_id"
			}`)}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	options := types.NewQueryOptions()
	options.SetScroll("1m")

	sr, err := k.Document.Search("index", "collection", json.RawMessage(`{}`), options)
	assert.Nil(t, err)
	assert.Equal(t, "scroll_id", sr.ScrollId)

	c.MockSend = func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
		parsedQuery := &types.KuzzleRequest{}
		json.Unmarshal(query, parsedQuery)

		assert.Equal(t, "1m", options.Scroll())
		assert.Equal(t, "scroll_id", options.ScrollId())
		assert.Equal(t, "scroll", parsedQuery.Action)

		return &types.KuzzleResponse{Result: json.RawMessage(`{
			"hits": ["id3", "id4"],
			"total": 42,
			"_scroll_id": "new_scroll"
		}`)}
	}

	nsr, err := sr.Next()
	assert.Nil(t, err)
	assert.Equal(t, 4, nsr.Fetched)
	assert.Equal(t, "new_scroll", nsr.ScrollId)
}
