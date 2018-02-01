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
