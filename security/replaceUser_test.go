// Copyright 2015-2017 Kuzzle
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
	"testing"

	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func TestReplaceUserIDNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	res, err := k.Security.ReplaceUser("", []byte(`{"body": "test"}`), nil)
	assert.Nil(t, res)
	assert.Error(t, err)
}

func TestReplaceUserContentNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	res, err := k.Security.ReplaceUser("id", nil, nil)
	assert.Nil(t, res)
	assert.Error(t, err)
}

func TestReplaceUserError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: types.KuzzleError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	res, err := k.Security.ReplaceUser("id", []byte(`{"body": "test"}`), nil)
	assert.Nil(t, res)
	assert.Error(t, err)
}

func TestReplaceUser(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {

			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)
			assert.Equal(t, "security", request.Controller)
			assert.Equal(t, "replaceUser", request.Action)
			assert.Equal(t, "id", request.Id)

			return &types.KuzzleResponse{Result: []byte(`{
          "_id": "id",
          "_source": {
            "profileIds": ["profileId"],
            "name": "John Doe"
          }
        }`),
			}
		},
	}

	k, _ := kuzzle.NewKuzzle(c, nil)
	res, err := k.Security.ReplaceUser("id", []byte(`{"body": "test"}`), nil)
	assert.NoError(t, err)
	assert.Equal(t, "id", res.Id)
}
