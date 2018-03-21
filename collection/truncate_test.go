package collection_test

import (
	"fmt"

	"github.com/kuzzleio/sdk-go/collection"

	"testing"

	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func TestTruncateIndexNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	nc := collection.NewCollection(k)
	err := nc.Truncate("", "collection", nil)
	assert.NotNil(t, err)
}

func TestTruncateCollectionNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	nc := collection.NewCollection(k)
	err := nc.Truncate("index", "", nil)
	assert.NotNil(t, err)
}

func TestTruncateError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	nc := collection.NewCollection(k)
	err := nc.Truncate("index", "collection", nil)
	assert.NotNil(t, err)
}

func TestTruncate(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Result: []byte(`{
				"acknowledged": true
			}`)}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	nc := collection.NewCollection(k)
	err := nc.Truncate("index", "collection", nil)
	assert.Nil(t, err)
}

func ExampleCollection_Truncate() {
	c := &internal.MockedConnection{}
	k, _ := kuzzle.NewKuzzle(c, nil)

	nc := collection.NewCollection(k)
	err := nc.Truncate("index", "collection", nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
