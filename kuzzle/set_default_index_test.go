package kuzzle

import (
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/stretchr/testify/assert"
	"testing"
	"fmt"
	"github.com/kuzzleio/sdk-go/connection/websocket"
)

func TestSetDefaultIndexNullIndex(t *testing.T) {
	k, _ := NewKuzzle(internal.MockedConnection{}, nil)
	assert.NotNil(t, k.SetDefaultIndex(""))
}

func TestSetDefaultIndex(t *testing.T) {
	k, _ := NewKuzzle(internal.MockedConnection{}, nil)
	k.SetDefaultIndex("myindex")
	assert.Equal(t, "myindex", k.defaultIndex)
}

func ExampleKuzzle_SetDefaultIndex() {
	conn := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := NewKuzzle(conn, nil)

	err := k.SetDefaultIndex("index")

	if err != nil {
		fmt.Println(err.Error())
		return
	}
}