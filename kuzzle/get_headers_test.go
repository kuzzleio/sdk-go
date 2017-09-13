package kuzzle

import (
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/stretchr/testify/assert"
	"testing"
	"fmt"
	"github.com/kuzzleio/sdk-go/connection/websocket"
)

func TestGetHeaders(t *testing.T) {
	k, _ := NewKuzzle(internal.MockedConnection{}, nil)

	m := make(map[string]interface{})
	m["foo"] = "bar"
	k.SetHeaders(m, false)
	assert.Equal(t, "bar", k.headers["foo"])
	assert.Equal(t, "bar", k.GetHeader("foo"))
}

func TestGetHeadersReplace(t *testing.T) {
	k, _ := NewKuzzle(internal.MockedConnection{}, nil)

	m := make(map[string]interface{})
	m["foo"] = "bar"
	k.SetHeaders(m, false)
	assert.Equal(t, "bar", k.headers["foo"])
	delete(m, "foo")
	m["oof"] = "bar"
	k.SetHeaders(m, true)
	assert.Nil(t, k.headers["foo"])
	assert.Nil(t, k.GetHeader("foo"))
	assert.Equal(t, "bar", k.GetHeader("oof"))
}

func ExampleKuzzle_GetHeaders() {
	conn := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := NewKuzzle(conn, nil)

	headers := make(map[string]interface{})
	headers["foo"] = "bar"
	k.SetHeaders(headers, false)
	res := k.GetHeaders()

	fmt.Println(res["foo"])
}

func ExampleKuzzle_GetHeader() {
	conn := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := NewKuzzle(conn, nil)

	headers := make(map[string]interface{})
	headers["foo"] = "bar"
	k.SetHeaders(headers, false)
	res := k.GetHeader("foo")

	fmt.Println(res)
}