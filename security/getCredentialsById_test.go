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
	"testing"

	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func TestGetCredentialsByIDStrategyNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	_, err := k.Security.GetCredentialsByID("", "id", nil)
	assert.Error(t, err)
}

func TestGetCredentialsByIDIDNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	_, err := k.Security.GetCredentialsByID("strategy", "", nil)
	assert.Error(t, err)
}

func TestGetCredentialsByIDError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: types.KuzzleError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := k.Security.GetCredentialsByID("strategy", "id", nil)
	assert.NotNil(t, err)
}

func TestGetCredentialsByID(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "getCredentialsById", parsedQuery.Action)
			assert.Equal(t, "id", parsedQuery.Id)
			assert.Equal(t, "strategy", parsedQuery.Strategy)

			return &types.KuzzleResponse{Result: []byte(`{"username": "user", "kuid": "id"}`)}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, err := k.Security.GetCredentialsByID("strategy", "id", nil)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}
