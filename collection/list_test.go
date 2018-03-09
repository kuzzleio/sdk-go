package collection_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/kuzzleio/sdk-go/collection"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func TestListIndexNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	nc := collection.NewCollection(k)
	from := 0
	size := 1
	lo := collection.ListOptions{Type: "stored", From: &from, Size: &size}
	_, err := nc.List("", &lo)

	assert.NotNil(t, err)
}

func TestListError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: types.NewError("Unit test error")}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	nc := collection.NewCollection(k)
	from := 0
	size := 1
	lo := collection.ListOptions{Type: "stored", From: &from, Size: &size}
	_, err := nc.List("index", &lo)
	assert.NotNil(t, err)
	assert.Equal(t, "Unit test error", err.(*types.KuzzleError).Message)
}

func TestList(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "collection", parsedQuery.Controller)
			assert.Equal(t, "list", parsedQuery.Action)
			assert.Equal(t, "index", parsedQuery.Index)

			res := types.KuzzleResponse{Result: []byte(`
				{
					"collections": [
						{
							"name": "stored_n", "type": "stored"
						}
					],
					"type": "all"
				}`),
			}

			r, _ := json.Marshal(res.Result)
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	nc := collection.NewCollection(k)
	from := 0
	size := 1
	lo := collection.ListOptions{Type: "stored", From: &from, Size: &size}
	res, err := nc.List("index", &lo)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func ExampleCollection_List() {
	c := &internal.MockedConnection{}
	k, _ := kuzzle.NewKuzzle(c, nil)

	nc := collection.NewCollection(k)

	from := 0
	size := 1
	lo := collection.ListOptions{Type: "stored", From: &from, Size: &size}
	res, err := nc.List("index", &lo)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res)
}
