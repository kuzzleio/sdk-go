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

func TestUpdateDocumentEmptyId(t *testing.T) {
	type Document struct {
		Name     string
		Function string
	}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Collection.UpdateDocument: document id required"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := collection.NewCollection(k, "collection", "index").UpdateDocument("", Document{Name: "Obi Wan", Function: "Legend"}, nil)
	assert.NotNil(t, err)
}

func TestUpdateDocumentError(t *testing.T) {
	type Document struct {
		Name     string
		Function string
	}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := collection.NewCollection(k, "collection", "index").UpdateDocument("id", Document{Name: "Obi Wan", Function: "Legend"}, nil)
	assert.NotNil(t, err)
}

func TestUpdateDocument(t *testing.T) {
	id := "myId"

	type InitialContent struct {
		Name     string
		Function string
	}
	initialContent := InitialContent{
		Name:     "Anakin",
		Function: "Padawan",
	}

	type NewContent struct {
		Function string
	}
	updatePart := NewContent{"Jedi Knight"}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "document", parsedQuery.Controller)
			assert.Equal(t, "update", parsedQuery.Action)
			assert.Equal(t, "index", parsedQuery.Index)
			assert.Equal(t, "collection", parsedQuery.Collection)
			assert.Equal(t, id, parsedQuery.Id)

			assert.Equal(t, "Jedi Knight", parsedQuery.Body.(map[string]interface{})["Function"])

			res := types.Document{Id: id, Source: []byte(`{"Name":"Anakin","Function":"Jedi Knight"}`)}
			r, _ := json.Marshal(res)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := collection.NewCollection(k, "collection", "index").UpdateDocument(id, updatePart, nil)

	assert.Equal(t, id, res.Id)

	var result InitialContent
	json.Unmarshal(res.Source, &result)

	assert.Equal(t, initialContent.Name, result.Name)
	assert.NotEqual(t, initialContent.Function, result.Name)
	assert.Equal(t, updatePart.Function, result.Function)
}

func TestMUpdateDocumentEmptyDocuments(t *testing.T) {
	documents := []types.Document{}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Collection.MUpdateDocument: please provide at least one document to update"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := collection.NewCollection(k, "collection", "index").MUpdateDocument(documents, nil)
	assert.NotNil(t, err)
}

func TestMUpdateDocumentError(t *testing.T) {
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

	_, err := collection.NewCollection(k, "collection", "index").MUpdateDocument(documents, nil)
	assert.NotNil(t, err)
}

func TestMUpdateDocument(t *testing.T) {
	documents := []types.Document{
		{Id: "foo", Source: []byte(`{"title":"Foo"}`)},
		{Id: "bar", Source: []byte(`{"title":"Bar"}`)},
	}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "document", parsedQuery.Controller)
			assert.Equal(t, "mUpdate", parsedQuery.Action)
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

	res, _ := collection.NewCollection(k, "collection", "index").MUpdateDocument(documents, nil)
	assert.Equal(t, 2, res.Total)

	for index, doc := range res.Hits {
		assert.Equal(t, documents[index].Id, doc.Id)
		assert.Equal(t, documents[index].Source, doc.Source)
	}
}
