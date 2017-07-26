package kuzzle_test

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestListCollectionsIndexNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	_, err := k.ListCollections("", nil)
	assert.NotNil(t, err)
}

func TestListCollectionsQueryError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)
			assert.Equal(t, "collection", request.Controller)
			assert.Equal(t, "index", request.Index)
			assert.Equal(t, "list", request.Action)
			return types.KuzzleResponse{Error: types.MessageError{Message: "error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	_, err := k.ListCollections("index", nil)
	assert.NotNil(t, err)
}

func TestListCollections(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)
			assert.Equal(t, "collection", request.Controller)
			assert.Equal(t, "list", request.Action)

			type collections struct {
				Collections []types.CollectionsList `json:"collections"`
			}

			list := make([]types.CollectionsList, 0)
			list = append(list, types.CollectionsList{Name: "collection1", Type: "stored"})
			list = append(list, types.CollectionsList{Name: "collection2", Type: "stored"})

			c := collections{
				Collections: list,
			}

			h, err := json.Marshal(c)
			if err != nil {
				log.Fatal(err)
			}

			return types.KuzzleResponse{Result: h}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := k.ListCollections("index", nil)

	assert.Equal(t, "collection1", res[0].Name)
	assert.Equal(t, "collection2", res[1].Name)
	assert.Equal(t, "stored", res[0].Type)
	assert.Equal(t, "stored", res[1].Type)

}
