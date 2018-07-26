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

func TestGetMappingIndexNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	nc := collection.NewCollection(k)
	_, err := nc.GetMapping("", "collection", nil)
	assert.NotNil(t, err)
}

func TestGetMappingCollectionNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	nc := collection.NewCollection(k)
	_, err := nc.GetMapping("index", "", nil)
	assert.NotNil(t, err)
}

func TestGetMappingError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: types.NewError("Unit test error")}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	nc := collection.NewCollection(k)
	_, err := nc.GetMapping("index", "collection", nil)
	assert.NotNil(t, err)
	assert.Equal(t, "Unit test error", err.(types.KuzzleError).Message)
}

func TestGetMapping(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "collection", parsedQuery.Controller)
			assert.Equal(t, "getMapping", parsedQuery.Action)
			assert.Equal(t, "index", parsedQuery.Index)
			assert.Equal(t, "collection", parsedQuery.Collection)

			res := types.KuzzleResponse{Result: []byte(`
				{
					"index":{
						"mappings":{
							"collection":{
								"properties":{
									"foo":{
										"type":"text",
										"ignore_above":255
									}
								}
							}
						}
					}
				}`),
			}

			r, _ := json.Marshal(res.Result)
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	nc := collection.NewCollection(k)
	res, err := nc.GetMapping("index", "collection", nil)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func ExampleCollection_GetMapping() {
	c := &internal.MockedConnection{}
	k, _ := kuzzle.NewKuzzle(c, nil)

	nc := collection.NewCollection(k)
	res, err := nc.GetMapping("index", "collection", nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res)
}
