package kuzzle_test

import (
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"testing"
)

func TestAddListener(t *testing.T) {
	c := &internal.MockedConnection{}
	k, _ := kuzzle.NewKuzzle(c, nil)

	ch := make(chan interface{})

	kuzzle.AddListener(*k, 0, ch)

	//c.AssertCalled(t, "AddListener")
	//c.AssertNumberOfCalls(t, "AddListener", 1)
}
