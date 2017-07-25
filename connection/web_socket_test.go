package connection

import (
	"github.com/kuzzleio/sdk-go/event"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveAllListeners(t *testing.T) {
	c := WebSocket{eventListeners: make(map[int]chan<- interface{})}
	c.Connect()
	c.AddListener(event.LoginAttempt, nil)
	c.AddListener(event.Disconnected, nil)
	assert.Equal(t, 2, len(c.eventListeners))
	c.RemoveAllListeners()
	assert.Equal(t, 0, len(c.eventListeners))
}

func TestRemoveListener(t *testing.T) {
	c := WebSocket{eventListeners: make(map[int]chan<- interface{})}
	c.AddListener(event.LoginAttempt, make(chan interface{}))
	c.AddListener(event.Disconnected, make(chan interface{}))
	assert.Equal(t, 2, len(c.eventListeners))
	c.RemoveListener(event.LoginAttempt)
	assert.Nil(t, c.eventListeners[event.LoginAttempt])
	assert.NotNil(t, c.eventListeners[event.Disconnected])
}
