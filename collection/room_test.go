package collection_test

import (
	"github.com/kuzzleio/sdk-go/collection"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitRoomWithoutOptions(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{}
		},
	}

	k, _ := kuzzle.NewKuzzle(c, nil)

	r := collection.NewRoom(*collection.NewCollection(k, "collection", "index"), nil)

	assert.NotNil(t, r)
}
