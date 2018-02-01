package websocket

import (
	"testing"

	"github.com/kuzzleio/sdk-go/event"
	"github.com/stretchr/testify/assert"
)

func TestAddListener(t *testing.T) {
	c := webSocket{eventListeners: make(map[int]map[chan<- interface{}]bool)}
	c.Connect()
	c.AddListener(event.LoginAttempt, nil)
	c.AddListener(event.Disconnected, nil)
	assert.Equal(t, 2, len(c.eventListeners))
}

func TestRemoveAllListeners(t *testing.T) {
	c := webSocket{eventListeners: make(map[int]map[chan<- interface{}]bool)}
	c.Connect()
	c.AddListener(event.LoginAttempt, nil)
	c.AddListener(event.Disconnected, nil)
	assert.Equal(t, 2, len(c.eventListeners))
	c.RemoveAllListeners(event.LoginAttempt)
	c.RemoveAllListeners(event.Disconnected)
	assert.Equal(t, 0, len(c.eventListeners))
}

func TestRemoveListener(t *testing.T) {
	c := webSocket{eventListeners: make(map[int]map[chan<- interface{}]bool)}
	listener := make(chan interface{})
	c.AddListener(event.LoginAttempt, listener)
	c.AddListener(event.Disconnected, make(chan interface{}))
	assert.Equal(t, 1, len(c.eventListeners[event.LoginAttempt]))
	c.RemoveListener(event.LoginAttempt, listener)
	assert.Equal(t, 0, len(c.eventListeners[event.LoginAttempt]))
	assert.Equal(t, 1, len(c.eventListeners[event.Disconnected]))
}

func TestOnce(t *testing.T) {
	c := webSocket{eventListenersOnce: make(map[int]map[chan<- interface{}]bool)}
	listener := make(chan interface{})
	go func() {
		<-listener
	}()
	c.Once(event.LoginAttempt, listener)
	assert.Equal(t, 1, len(c.eventListenersOnce[event.LoginAttempt]))
	c.EmitEvent(event.LoginAttempt, nil)
	assert.Equal(t, 0, len(c.eventListenersOnce[event.LoginAttempt]))
}
