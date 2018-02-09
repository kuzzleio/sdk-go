package collection

import (
	"testing"

	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/state"
	"github.com/stretchr/testify/assert"
)

func TestUnsubscribeSubscribing(t *testing.T) {
	var k *kuzzle.Kuzzle

	c := &internal.MockedConnection{}
	k, _ = kuzzle.NewKuzzle(c, nil)
	c.SetState(state.Connected)

	r := NewRoom(NewCollection(k, "collection", "index"), nil, nil)
	r.internalState = subscribing
	assert.NotNil(t, r.Unsubscribe())
}

func TestUnsubscribe(t *testing.T) {
	var k *kuzzle.Kuzzle
	var removedListener bool
	c := &internal.MockedConnection{
		MockRemoveListener: func(e int, c chan<- interface{}) {
			removedListener = true
		},
	}
	k, _ = kuzzle.NewKuzzle(c, nil)
	c.SetState(state.Connected)

	r := NewRoom(NewCollection(k, "collection", "index"), nil, nil)
	r.isListening = true
	r.internalState = active
	r.Unsubscribe()
	assert.Equal(t, true, removedListener)
}
