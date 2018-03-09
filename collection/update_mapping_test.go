package collection_test

import (
	"encoding/json"
	"fmt"

	"github.com/kuzzleio/sdk-go/collection"

	"testing"

	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func TestUpdateMappingIndexNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	nc := collection.NewCollection(k)
	err := nc.UpdateMapping("", "collection", json.RawMessage("body"))
	assert.NotNil(t, err)
}

func TestUpdateMappingCollectionNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	nc := collection.NewCollection(k)
	err := nc.UpdateMapping("index", "", json.RawMessage("body"))
	assert.NotNil(t, err)
}

func TestUpdateMappingBodyNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	nc := collection.NewCollection(k)
	err := nc.UpdateMapping("index", "collection", nil)
	assert.NotNil(t, err)
}

func TestUpdateMappingError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	nc := collection.NewCollection(k)
	err := nc.UpdateMapping("index", "collection", json.RawMessage("body"))
	assert.NotNil(t, err)
}

func TestUpdateMapping(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Result: []byte(`{}`)}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	nc := collection.NewCollection(k)
	err := nc.UpdateMapping("index", "collection", json.RawMessage("body"))
	assert.Nil(t, err)
}

func ExampleCollection_UpdateMapping() {
	c := &internal.MockedConnection{}
	k, _ := kuzzle.NewKuzzle(c, nil)

	nc := collection.NewCollection(k)
	err := nc.UpdateMapping("index", "collection", json.RawMessage("body"))

	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
