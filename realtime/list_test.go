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

package realtime_test

import (
	"encoding/json"
	"testing"

	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/realtime"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func TestListIndexNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	nr := realtime.NewRealtime(k)

	_, err := nr.List("", "collection", nil)

	assert.NotNil(t, err)
}

func TestListCollectionNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	nr := realtime.NewRealtime(k)

	_, err := nr.List("index", "", nil)

	assert.NotNil(t, err)
}

func TestListError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: types.NewError("Unit test error")}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	nr := realtime.NewRealtime(k)

	_, err := nr.List("index", "collection", nil)
	assert.NotNil(t, err)
	assert.Equal(t, "Unit test error", err.(types.KuzzleError).Message)
}

func TestList(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "realtime", parsedQuery.Controller)
			assert.Equal(t, "list", parsedQuery.Action)
			assert.Equal(t, "index", parsedQuery.Index)
			assert.Equal(t, "collection", parsedQuery.Collection)

			res := types.KuzzleResponse{Result: []byte(`
				{
          "index": {
            "collection": {
              "afcd909773f197ab859447594bfbd154": 12,
              "4adbc1948ac4dc84ac89d14b488bcad1": 4
            }
          }
			  }`),
			}

			r, _ := json.Marshal(res.Result)
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	nr := realtime.NewRealtime(k)

	res, err := nr.List("index", "collection", nil)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
