package kuzzle

import (
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/stretchr/testify/assert"
	"testing"
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
