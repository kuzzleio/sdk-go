package collection_test

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/collection"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReplaceDocumentEmptyId(t *testing.T) {
	type Document struct {
		Title string
	}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Collection.ReplaceDocument: document id required"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := collection.NewCollection(k, "collection", "index").ReplaceDocument("", Document{Title: "jonathan"}, nil)
	assert.NotNil(t, err)
}

func TestReplaceDocumentError(t *testing.T) {
	type Document struct {
		Title string
	}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := collection.NewCollection(k, "collection", "index").ReplaceDocument("id", Document{Title: "jonathan"}, nil)
	assert.NotNil(t, err)
}

func TestReplaceDocument(t *testing.T) {
	type Document struct {
		Title string
	}

	id := "myId"

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "document", parsedQuery.Controller)
			assert.Equal(t, "replace", parsedQuery.Action)
			assert.Equal(t, "index", parsedQuery.Index)
			assert.Equal(t, "collection", parsedQuery.Collection)
			assert.Equal(t, id, parsedQuery.Id)

			assert.Equal(t, "jonathan", parsedQuery.Body.(map[string]interface{})["Title"])

			res := types.Document{Id: id, Source: []byte(`{"title": "jonathan"}`)}
			r, _ := json.Marshal(res)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := collection.NewCollection(k, "collection", "index").ReplaceDocument(id, Document{Title: "jonathan"}, nil)
	assert.Equal(t, id, res.Id)
}

func TestMReplaceDocumentEmptyDocuments(t *testing.T) {
	documents := []types.Document{}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Collection.MReplaceDocument: please provide at least one document to replace"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := collection.NewCollection(k, "collection", "index").MReplaceDocument(documents, nil)
	assert.NotNil(t, err)
}

func TestMReplaceDocumentError(t *testing.T) {
	documents := []types.Document{
		{Id: "foo", Source: []byte(`{"title":"Foo"}`)},
		{Id: "bar", Source: []byte(`{"title":"Bar"}`)},
	}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := collection.NewCollection(k, "collection", "index").MReplaceDocument(documents, nil)
	assert.NotNil(t, err)
}

func TestMReplaceDocument(t *testing.T) {
	documents := []types.Document{
		{Id: "foo", Source: []byte(`{"title":"Foo"}`)},
		{Id: "bar", Source: []byte(`{"title":"Bar"}`)},
	}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "document", parsedQuery.Controller)
			assert.Equal(t, "mReplace", parsedQuery.Action)
			assert.Equal(t, "index", parsedQuery.Index)
			assert.Equal(t, "collection", parsedQuery.Collection)

			results := []types.KuzzleResult{
				{Id: "foo", Source: []byte(`{"title":"Foo"}`)},
				{Id: "bar", Source: []byte(`{"title":"Bar"}`)},
			}

			res := types.KuzzleSearchResult{
				Total: 2,
				Hits:  results,
			}
			r, _ := json.Marshal(res)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := collection.NewCollection(k, "collection", "index").MReplaceDocument(documents, nil)
	assert.Equal(t, 2, res.Total)

	for index, doc := range res.Hits {
		assert.Equal(t, documents[index].Id, doc.Id)
		assert.Equal(t, documents[index].Source, doc.Source)
	}
}
