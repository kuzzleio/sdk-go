package kuzzle

import (
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/stretchr/testify/assert"
	"testing"
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
