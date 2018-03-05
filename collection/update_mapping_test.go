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

func TestUpdateSpecificationsIndexNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	nc := collection.NewCollection(k)
	err := nc.UpdateSpecifications("", "collection", "body")
	assert.NotNil(t, err)
}

func TestUpdateSpecificationsCollectionNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	nc := collection.NewCollection(k)
	err := nc.UpdateSpecifications("index", "", "body")
	assert.NotNil(t, err)
}

func TestUpdateSpecificationsBodyNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	nc := collection.NewCollection(k)
	err := nc.UpdateSpecifications("index", "collection", "")
	assert.NotNil(t, err)
}

func TestUpdateSpecificationsError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	nc := collection.NewCollection(k)
	err := nc.UpdateSpecifications("index", "collection", "body")
	assert.NotNil(t, err)
}

func TestUpdateSpecifications(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Result: []byte(`{}`)}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	nc := collection.NewCollection(k)
	err := nc.UpdateSpecifications("index", "collection", "body")
	assert.Nil(t, err)
}

func ExampleCollection_UpdateSpecifications() {
	c := &internal.MockedConnection{}
	k, _ := kuzzle.NewKuzzle(c, nil)

	nc := collection.NewCollection(k)
	err := nc.UpdateSpecifications("index", "collection", "body")

	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
