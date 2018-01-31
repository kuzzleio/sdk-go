package websocket

import (
	"github.com/kuzzleio/sdk-go/event"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveAllListeners(t *testing.T) {
	c := webSocket{eventListeners: make(map[int]map[chan<- interface{}]struct{})}
	c.Connect()
	c.AddListener(event.LoginAttempt, nil)
	c.AddListener(event.Disconnected, nil)
	assert.Equal(t, 2, len(c.eventListeners))
	c.RemoveAllListeners(event.LoginAttempt)
	c.RemoveAllListeners(event.Disconnected)
	assert.Equal(t, 0, len(c.eventListeners))
}

func TestRemoveListener(t *testing.T) {
	c := webSocket{eventListeners: make(map[int]map[chan<- interface{}]struct{})}
	listener := make(chan interface{})
	c.AddListener(event.LoginAttempt, listener)
	c.AddListener(event.Disconnected, make(chan interface{}))
	assert.Equal(t, 2, len(c.eventListeners))
	c.RemoveListener(event.LoginAttempt, listener)
	assert.Nil(t, c.eventListeners[event.LoginAttempt])
	assert.NotNil(t, c.eventListeners[event.Disconnected])
}
