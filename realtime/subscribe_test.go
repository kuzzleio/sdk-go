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
	"github.com/kuzzleio/sdk-go/state"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func TestSubscribeIndexNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	nr := realtime.NewRealtime(k)

	notifChan := make(chan<- types.NotificationResult)
	_, err := nr.Subscribe("", "collection", json.RawMessage(`{"filter": "filter"}`), notifChan, nil)

	assert.NotNil(t, err)
}

func TestSubscribeCollectionNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	nr := realtime.NewRealtime(k)

	notifChan := make(chan<- types.NotificationResult)
	_, err := nr.Subscribe("index", "", json.RawMessage(`{"filter": "filter"}`), notifChan, nil)

	assert.NotNil(t, err)
}

func TestSubscribeNotifChannelNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "realtime", parsedQuery.Controller)
			assert.Equal(t, "publish", parsedQuery.Action)
			assert.Equal(t, "index", parsedQuery.Index)
			assert.Equal(t, "collection", parsedQuery.Collection)
			assert.NotNil(t, parsedQuery.Body)

			res := types.KuzzleResponse{Result: []byte(`{}`)}

			r, _ := json.Marshal(res.Result)
			return &types.KuzzleResponse{Result: r}
		},
	}, nil)
	nr := realtime.NewRealtime(k)

	_, err := nr.Subscribe("index", "collection", json.RawMessage(`{"filter": "filter"}`), nil, nil)

	assert.NotNil(t, err)
}

func TestSubscribeQueryError(t *testing.T) {
	var k *kuzzle.Kuzzle

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: types.KuzzleError{Message: "ah!"}}
		},
	}
	k, _ = kuzzle.NewKuzzle(c, nil)
	nr := realtime.NewRealtime(k)
	c.SetState(state.Connected)

	notifChan := make(chan<- types.NotificationResult)
	_, err := nr.Subscribe("index", "collection", json.RawMessage(`{"filter": "filter"}`), notifChan, nil)
	assert.Equal(t, "ah!", err.Error())
}

func TestRenewWithSubscribeToSelf(t *testing.T) {
	var k *kuzzle.Kuzzle

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			roomRaw := []byte(`{"requestId": "rqid", "channel": "foo", "roomId": "42"}`)
			return &types.KuzzleResponse{RequestId: "ah!", Result: roomRaw}
		},
	}
	k, _ = kuzzle.NewKuzzle(c, nil)
	nr := realtime.NewRealtime(k)
	c.SetState(state.Connected)

	notifChan := make(chan<- types.NotificationResult)
	res, err := nr.Subscribe("index", "collection", json.RawMessage(`{"filter": "filter"}`), notifChan, nil)
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, "42", res.Room)
}

func TestRoomSubscribe(t *testing.T) {
	var k *kuzzle.Kuzzle

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			roomRaw := []byte(`{"requestId": "rqid", "channel": "foo", "roomId": "42"}`)
			return &types.KuzzleResponse{Result: roomRaw}
		},
	}
	k, _ = kuzzle.NewKuzzle(c, nil)
	nr := realtime.NewRealtime(k)
	c.SetState(state.Connected)

	notifChan := make(chan<- types.NotificationResult)
	res, err := nr.Subscribe("index", "collection", json.RawMessage(`{"filter": "filter"}`), notifChan, nil)
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, "42", res.Room)

}

func TestRoomSubscribeNotConnected(t *testing.T) {
	var k *kuzzle.Kuzzle

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: types.KuzzleError{Message: "Not Connected"}}
		},
	}

	k, _ = kuzzle.NewKuzzle(c, nil)
	nr := realtime.NewRealtime(k)
	c.SetState(state.Connected)

	notifChan := make(chan<- types.NotificationResult)
	_, err := nr.Subscribe("collection", "index", json.RawMessage("{}"), notifChan, nil)
	assert.NotNil(t, err)
	assert.Equal(t, "Not Connected", err.Error())
}

func ExampleRealtime_Subscribe() {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			roomRaw := []byte(`{"requestId": "rqid", "channel": "foo", "roomId": "42"}`)
			return &types.KuzzleResponse{Result: roomRaw}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	nr := realtime.NewRealtime(k)
	c.SetState(state.Connected)

	notifChan := make(chan<- types.NotificationResult)
	nr.Subscribe("collection", "index", json.RawMessage(""), notifChan, nil)
}
