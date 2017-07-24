package connection

import (
  "testing"
  "github.com/kuzzleio/sdk-go/event"
  "github.com/magiconair/properties/assert"
)

func TestRemoveAllListeners(t *testing.T)  {
  c := WebSocket{eventListeners: make(map[int] chan<- interface{})}
  c.Connect()
  c.AddListener(event.LoginAttempt, nil)
  c.AddListener(event.Disconnected, nil)
  assert.Equal(t, 2, len(c.eventListeners))
  c.RemoveAllListeners()
  assert.Equal(t, 0, len(c.eventListeners))
}