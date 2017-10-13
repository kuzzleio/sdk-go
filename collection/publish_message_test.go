package collection_test

import (
	"encoding/json"
	"fmt"
	"github.com/kuzzleio/sdk-go/collection"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPublishKuzzleError(t *testing.T) {
	type Document struct {
		Title string
	}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := collection.NewCollection(k, "collection", "index").PublishMessage(&Document{Title: "yolo"}, nil)
	assert.NotNil(t, err)
}

func TestPublishMessage(t *testing.T) {
	type Document struct {
		Title string
	}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "realtime", parsedQuery.Controller)
			assert.Equal(t, "publish", parsedQuery.Action)
			assert.Equal(t, "index", parsedQuery.Index)
			assert.Equal(t, "collection", parsedQuery.Collection)

			assert.Equal(t, "yolo", parsedQuery.Body.(map[string]interface{})["Title"])

			res := types.KuzzleResponse{Result: []byte(`{"published":true}`)}
			r, _ := json.Marshal(res.Result)
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	coll := collection.NewCollection(k, "collection", "index")
	res, _ := coll.PublishMessage(&Document{Title: "yolo"}, nil)
	assert.Equal(t, coll, res)
}

func ExampleCollection_PublishMessage() {
	type Document struct {
		Title string
	}

	c := &internal.MockedConnection{}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, err := collection.NewCollection(k, "collection", "index").PublishMessage(&Document{Title: "yolo"}, nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res)
}
