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

package kuzzle

import (
	"encoding/json"
	"testing"

	"github.com/kuzzleio/sdk-go/internal"
	"github.com/stretchr/testify/assert"
)

func TestAddListener(t *testing.T) {
	c := &internal.MockedConnection{}
	k, _ := NewKuzzle(c, nil)

	ch := make(chan json.RawMessage)

	k.AddListener(0, ch)
	assert.Equal(t, 1, len(k.eventListeners[0]))
}

func TestRemoveListener(t *testing.T) {
	c := &internal.MockedConnection{}
	k, _ := NewKuzzle(c, nil)

	ch := make(chan json.RawMessage)

	k.AddListener(0, ch)
	k.AddListener(1, ch)
	k.RemoveListener(0, ch)

	assert.Equal(t, 0, len(k.eventListeners[0]))
}
func TestRemoveAllListener(t *testing.T) {
	c := &internal.MockedConnection{}
	k, _ := NewKuzzle(c, nil)

	ch := make(chan json.RawMessage)
	ch2 := make(chan json.RawMessage)

	k.AddListener(0, ch)
	k.AddListener(42, ch)
	k.AddListener(42, ch2)
	k.RemoveAllListeners(42)

	assert.Equal(t, 0, len(k.eventListeners[42]))
	assert.Equal(t, 1, len(k.eventListeners[0]))
}

func TestOnce(t *testing.T) {
	c := &internal.MockedConnection{}
	k, _ := NewKuzzle(c, nil)

	ch := make(chan json.RawMessage)

	k.Once(0, ch)
	assert.True(t, k.eventListenersOnce[0][ch])
}

func TestListenerCount(t *testing.T) {
	c := &internal.MockedConnection{}
	k, _ := NewKuzzle(c, nil)

	ch := make(chan json.RawMessage)
	ch2 := make(chan json.RawMessage)

	k.AddListener(0, ch)
	k.AddListener(0, ch2)

	assert.Equal(t, 2, k.ListenerCount(0))
}
