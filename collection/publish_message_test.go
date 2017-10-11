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

func TestPublishMessageError(t *testing.T) {

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	message := make(map[string]interface{})
	message["title"] = interface{}("yolo")
	_, err := collection.NewCollection(k, "collection", "index").PublishMessage(message, nil)
	assert.NotNil(t, err)
}

func TestPublishMessage(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "realtime", parsedQuery.Controller)
			assert.Equal(t, "publish", parsedQuery.Action)
			assert.Equal(t, "index", parsedQuery.Index)
			assert.Equal(t, "collection", parsedQuery.Collection)

			assert.Equal(t, "yolo", parsedQuery.Body.(map[string]interface{})["title"])

			res := types.KuzzleResponse{Result: []byte(`{"published":true}`)}
			r, _ := json.Marshal(res.Result)
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	message := make(map[string]interface{})
	message["title"] = interface{}("yolo")
	res, _ := collection.NewCollection(k, "collection", "index").PublishMessage(message, nil)

	assert.Equal(t, true, res.Published)
}

func ExampleCollection_PublishMessage() {
	type Document struct {
		Title string
	}

	c := &internal.MockedConnection{}
	k, _ := kuzzle.NewKuzzle(c, nil)

	message := make(map[string]interface{})
	message["title"] = interface{}("yolo")
	res, err := collection.NewCollection(k, "collection", "index").PublishMessage(message, nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res.Published)
}
