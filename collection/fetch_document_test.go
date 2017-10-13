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

func TestFetchDocumentEmptyId(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "Collection.FetchDocument: document id required"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := collection.NewCollection(k, "collection", "index").FetchDocument("", nil)
	assert.NotNil(t, err)
}

func TestFetchDocumentError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := collection.NewCollection(k, "collection", "index").FetchDocument("myId", nil)
	assert.NotNil(t, err)
}

func TestFetchDocument(t *testing.T) {
	type Document struct {
		Title string
	}

	id := "myId"

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "document", parsedQuery.Controller)
			assert.Equal(t, "get", parsedQuery.Action)
			assert.Equal(t, "index", parsedQuery.Index)
			assert.Equal(t, "collection", parsedQuery.Collection)
			assert.Equal(t, id, parsedQuery.Id)

			res := collection.Document{Id: id, Content: []byte(`{"foo": "bar"}`)}
			r, _ := json.Marshal(res)
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := collection.NewCollection(k, "collection", "index").FetchDocument(id, nil)
	assert.Equal(t, id, res.Id)
}

func ExampleCollection_FetchDocument() {
	id := "myId"
	c := &internal.MockedConnection{}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, err := collection.NewCollection(k, "collection", "index").FetchDocument(id, nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res.Id, res.Collection)
}

func TestMGetDocumentEmptyIds(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "Collection.MGetDocument: please provide at least one id of document to retrieve"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := collection.NewCollection(k, "collection", "index").MGetDocument([]string{}, nil)
	assert.NotNil(t, err)
}

func TestMGetDocumentError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := collection.NewCollection(k, "collection", "index").MGetDocument([]string{"foo", "bar"}, nil)
	assert.NotNil(t, err)
}

func TestMGetDocument(t *testing.T) {
	hits := []*collection.Document{
		{Id: "foo", Content: json.RawMessage(`{"title":"foo"}`)},
		{Id: "bar", Content: json.RawMessage(`{"title":"bar"}`)},
	}

	results := collection.SearchResult{Total: 2, Hits: hits}

	ids := []string{"foo", "bar"}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "document", parsedQuery.Controller)
			assert.Equal(t, "mGet", parsedQuery.Action)
			assert.Equal(t, "index", parsedQuery.Index)
			assert.Equal(t, "collection", parsedQuery.Collection)
			assert.Equal(t, []interface{}{"foo", "bar"}, parsedQuery.Body.(map[string]interface{})["ids"])

			res := collection.SearchResult{Total: results.Total, Hits: results.Hits}
			r, _ := json.Marshal(res)
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := collection.NewCollection(k, "collection", "index").MGetDocument(ids, nil)
	assert.Equal(t, results.Total, res.Total)

	for i, _ := range res.Hits {
		assert.Equal(t, hits[i].Id, res.Hits[i].Id)
		assert.Equal(t, hits[i].Content, res.Hits[i].Content)
	}
}

func ExampleCollection_MGetDocument() {
	ids := []string{"foo", "bar"}
	c := &internal.MockedConnection{}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, err := collection.NewCollection(k, "collection", "index").MGetDocument(ids, nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res.Hits[0].Id, res.Hits[0].Collection)
}
