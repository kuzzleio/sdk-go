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
	"encoding/json"
	"fmt"
	"testing"

	"github.com/kuzzleio/sdk-go/collection"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func TestSearchSpecificationsError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: types.KuzzleError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	nc := collection.NewCollection(k)

	_, err := nc.SearchSpecifications(nil, nil)
	assert.NotNil(t, err)
}

func TestSearchSpecifications(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "collection", parsedQuery.Controller)
			assert.Equal(t, "searchSpecifications", parsedQuery.Action)

			res := types.SpecificationSearchResult{
				ScrollId: "f00b4r",
				Total:    1,
				Hits:     make([]types.SpecificationSearchResultHit, 1),
			}
			res.Hits[0] = types.SpecificationSearchResultHit{
				Source: types.SpecificationEntry{
					Index:      "index",
					Collection: "collection",
					Validation: &types.Specification{
						Strict: false,
						Fields: types.SpecificationFields{
							"foo": types.SpecificationField{
								Mandatory:    true,
								Type:         "string",
								DefaultValue: "Value found with search",
							},
						},
					},
				},
			}

			r, _ := json.Marshal(res)
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	nc := collection.NewCollection(k)

	res, err := nc.SearchSpecifications(json.RawMessage(`{"foo": "bar"}`), nil)
	assert.Equal(t, 1, res.Total)
	assert.Nil(t, err)
}

func TestSearchSpecificationsNext(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Result: json.RawMessage(`{
				"hits": [
					{
						"_id": "id1",
						"_source": "specification1"
					},
					{
						"_id": "id1",
						"_source": "specification1"
					}
				],
				"total": 42
			}`)}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	options := types.NewQueryOptions()
	options.SetFrom(0)
	options.SetSize(2)

	sr, err := k.Collection.SearchSpecifications(json.RawMessage(`{}`), options)
	assert.Nil(t, err)
	assert.NotNil(t, sr)

	c.MockSend = func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
		assert.Equal(t, 2, options.From())
		assert.Equal(t, 2, options.Size())

		return &types.KuzzleResponse{Result: json.RawMessage(`{
			"hits": [
				{
					"_id": "id3",
					"_source": "specification3"
				},
				{
					"_id": "id4",
					"_source": "specification4"
				}
			],
			"total": 42
		}`)}
	}

	nsr, err := sr.Next()
	assert.Nil(t, err)
	assert.NotNil(t, nsr)
	assert.Equal(t, 42, nsr.Total)
	assert.Equal(t, 4, nsr.Fetched)
}

func ExampleCollection_SearchSpecifications() {
	c := &internal.MockedConnection{}
	k, _ := kuzzle.NewKuzzle(c, nil)

	nc := collection.NewCollection(k)

	res, err := nc.SearchSpecifications(json.RawMessage("{}"), nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res)
}
