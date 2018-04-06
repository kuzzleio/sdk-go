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

package kuzzle_test

import (
	"testing"

	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/stretchr/testify/assert"
)

func TestAddListener(t *testing.T) {
	called := false

	c := &internal.MockedConnection{
		MockAddListener: func(e int, c chan<- interface{}) {
			called = true
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	ch := make(chan interface{})

	k.AddListener(0, ch)
	assert.Equal(t, true, called)
}

func TestRemoveListener(t *testing.T) {
	called := false

	c := &internal.MockedConnection{
		MockRemoveListener: func(e int, c chan<- interface{}) {
			called = true
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	ch := make(chan interface{})

	k.RemoveListener(0, ch)
	assert.Equal(t, true, called)
}
func TestRemoveAllListener(t *testing.T) {
	called := false

	c := &internal.MockedConnection{
		MockRemoveAllListeners: func(e int) {
			called = true
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	k.RemoveAllListeners(0)
	assert.Equal(t, true, called)
}

func TestOnce(t *testing.T) {
	called := false

	c := &internal.MockedConnection{
		MockOnce: func(e int, c chan<- interface{}) {
			called = true
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	ch := make(chan interface{})

	k.Once(0, ch)
	assert.Equal(t, true, called)
}

func TestOn(t *testing.T) {
	called := false

	c := &internal.MockedConnection{
		MockAddListener: func(e int, c chan<- interface{}) {
			called = true
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	ch := make(chan interface{})

	k.On(0, ch)
	assert.Equal(t, true, called)
}

func TestListenerCount(t *testing.T) {
	called := false

	c := &internal.MockedConnection{
		MockListenerCount: func(e int) int {
			called = true
			return -1
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	k.ListenerCount(0)
	assert.Equal(t, true, called)
}
