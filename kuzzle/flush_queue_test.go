package kuzzle_test

import (
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlushQueue(t *testing.T) {
	c := internal.MockedConnection{}
	k, _ := kuzzle.NewKuzzle(c, nil)

	*k.GetOfflineQueue() = append(*k.GetOfflineQueue(), types.QueryObject{RequestId: "test"})
	assert.NotEmpty(t, *k.GetOfflineQueue())

	k.FlushQueue()
	assert.Empty(t, *k.GetOfflineQueue())
}
