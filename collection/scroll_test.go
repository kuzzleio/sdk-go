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

func TestScrollEmptyScrollId(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "Collection.Scroll: scroll id required"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := collection.NewCollection(k, "collection", "index").Scroll("", nil)
	assert.NotNil(t, err)
}

func TestScrollError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := collection.NewCollection(k, "collection", "index").Scroll("f00b4r", nil)
	assert.NotNil(t, err)
}

func TestScroll(t *testing.T) {
	hits := []*collection.Document{
		{Id: "doc42", Content: json.RawMessage(`{"foo":"bar"}`)},
	}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "document", parsedQuery.Controller)
			assert.Equal(t, "scroll", parsedQuery.Action)

			res := map[string]interface{}{"Total": 42, "Hits": hits}
			r, _ := json.Marshal(res)
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := collection.NewCollection(k, "collection", "index").Scroll("f00b4r", nil)
	assert.Equal(t, 42, res.Total)
	assert.Equal(t, hits[0].Id, res.Documents[0].Id)
	assert.Equal(t, hits[0].Content, res.Documents[0].Content)
	assert.Equal(t, len(hits), len(res.Documents))
}

func ExampleCollection_Scroll(t *testing.T) {
	hits := make([]*collection.Document, 1)
	hits[0] = &collection.Document{Id: "doc42", Content: json.RawMessage(`{"foo":"bar"}`)}

	c := &internal.MockedConnection{}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, err := collection.NewCollection(k, "collection", "index").Scroll("f00b4r", nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res.Documents[0].Id, res.Documents[0].Content)
}
