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

func TestReplaceDocumentEmptyId(t *testing.T) {

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.MessageError{Message: "Collection.ReplaceDocument: document id required"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := collection.NewCollection(k, "collection", "index").ReplaceDocument("", &collection.Document{Content: []byte(`{"title": "jonathan"}`)}, nil)
	assert.NotNil(t, err)
}

func TestReplaceDocumentError(t *testing.T) {

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := collection.NewCollection(k, "collection", "index").ReplaceDocument("id", &collection.Document{Content: []byte(`{"title": "jonathan"}`)}, nil)
	assert.NotNil(t, err)
}

func TestReplaceDocument(t *testing.T) {
	type Document struct {
		Title string
	}

	id := "myId"

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "document", parsedQuery.Controller)
			assert.Equal(t, "replace", parsedQuery.Action)
			assert.Equal(t, "index", parsedQuery.Index)
			assert.Equal(t, "collection", parsedQuery.Collection)
			assert.Equal(t, id, parsedQuery.Id)

			assert.Equal(t, "jonathan", parsedQuery.Body.(map[string]interface{})["title"])

			res := collection.Document{Id: id, Content: []byte(`{"title": "jonathan"}`)}
			r, _ := json.Marshal(res)
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := collection.NewCollection(k, "collection", "index").ReplaceDocument(id, &collection.Document{Content: []byte(`{"title": "jonathan"}`)}, nil)
	assert.Equal(t, id, res.Id)
}

func ExampleCollection_ReplaceDocument() {
	type Document struct {
		Title string
	}

	id := "myId"
	c := &internal.MockedConnection{}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, err := collection.NewCollection(k, "collection", "index").ReplaceDocument(id, &collection.Document{Content: []byte(`{"title": "jonathan"}`)}, nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res.Id, res.Collection)
}

func TestMReplaceDocumentEmptyDocuments(t *testing.T) {
	documents := []*collection.Document{}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.MessageError{Message: "Collection.MReplaceDocument: please provide at least one document to replace"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := collection.NewCollection(k, "collection", "index").MReplaceDocument(documents, nil)
	assert.NotNil(t, err)
}

func TestMReplaceDocumentError(t *testing.T) {
	documents := []*collection.Document{
		{Id: "foo", Content: []byte(`{"title":"Foo"}`)},
		{Id: "bar", Content: []byte(`{"title":"Bar"}`)},
	}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := collection.NewCollection(k, "collection", "index").MReplaceDocument(documents, nil)
	assert.NotNil(t, err)
}

func TestMReplaceDocument(t *testing.T) {
	documents := []*collection.Document{
		{Id: "foo", Content: []byte(`{"title":"Foo"}`)},
		{Id: "bar", Content: []byte(`{"title":"Bar"}`)},
	}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "document", parsedQuery.Controller)
			assert.Equal(t, "mReplace", parsedQuery.Action)
			assert.Equal(t, "index", parsedQuery.Index)
			assert.Equal(t, "collection", parsedQuery.Collection)

			results := []*collection.Document{
				{Id: "foo", Content: []byte(`{"title":"Foo"}`)},
				{Id: "bar", Content: []byte(`{"title":"Bar"}`)},
			}

			res := collection.SearchResult{
				Total: 2,
				Hits:  results,
			}
			r, _ := json.Marshal(res)
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := collection.NewCollection(k, "collection", "index").MReplaceDocument(documents, nil)
	assert.Equal(t, 2, res.Total)

	for index, doc := range res.Hits {
		assert.Equal(t, documents[index].Id, doc.Id)
		assert.Equal(t, documents[index].Content, doc.Content)
	}
}

func ExampleCollection_MReplaceDocument() {
	documents := []*collection.Document{
		{Id: "foo", Content: []byte(`{"title":"Foo"}`)},
		{Id: "bar", Content: []byte(`{"title":"Bar"}`)},
	}

	c := &internal.MockedConnection{}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, err := collection.NewCollection(k, "collection", "index").MReplaceDocument(documents, nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res.Hits[0].Id, res.Hits[0].Collection)
}
