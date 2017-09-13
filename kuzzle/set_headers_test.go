package kuzzle

import (
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/stretchr/testify/assert"
	"testing"
	"fmt"
	"github.com/kuzzleio/sdk-go/connection/websocket"
)

func TestSetHeaders(t *testing.T) {
	k, _ := NewKuzzle(internal.MockedConnection{}, nil)

	m := make(map[string]interface{})
	m["foo"] = "bar"
	k.SetHeaders(m, false)
	assert.Equal(t, "bar", k.headers["foo"])
}

func TestSetHeadersReplace(t *testing.T) {
	k, _ := NewKuzzle(internal.MockedConnection{}, nil)

	m := make(map[string]interface{})
	m["foo"] = "bar"
	k.SetHeaders(m, false)
	assert.Equal(t, "bar", k.headers["foo"])
	delete(m, "foo")
	m["oof"] = "bar"
	k.SetHeaders(m, true)
	assert.Nil(t, k.headers["foo"])
}

func ExampleKuzzle_SetHeaders() {
	conn := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := NewKuzzle(conn, nil)

	headers := make(map[string]interface{})
	headers["foo"] = "bar"
	k.SetHeaders(headers, false)
	res := k.GetHeaders()

	fmt.Println(res["foo"])
}