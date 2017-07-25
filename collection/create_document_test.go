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

func TestCreateDocumentError(t *testing.T) {
	type Document struct {
		Title string
	}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := collection.NewCollection(k, "collection", "index").CreateDocument("id", Document{Title: "yolo"}, nil)
	assert.NotNil(t, err)
}

func TestCreateDocumentWrongOptionError(t *testing.T) {
	type Document struct {
		Title string
	}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
			return types.KuzzleResponse{}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	newCollection := collection.NewCollection(k, "collection", "index")
	_, err := newCollection.CreateDocument("id", Document{Title: "yolo"}, &types.Options{IfExist: "unknown"})
	assert.Equal(t, "Invalid value for the 'ifExist' option: 'unknown'", fmt.Sprint(err))
}

func TestCreateDocument(t *testing.T) {
	type Document struct {
		Title string
	}

	id := "myId"

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "document", parsedQuery.Controller)
			assert.Equal(t, "create", parsedQuery.Action)
			assert.Equal(t, "index", parsedQuery.Index)
			assert.Equal(t, "collection", parsedQuery.Collection)
			assert.Equal(t, id, parsedQuery.Id)

			assert.Equal(t, "yolo", parsedQuery.Body.(map[string]interface{})["Title"])

			res := types.Document{Id: id, Source: []byte(`{"title": "yolo"}`)}
			r, _ := json.Marshal(res)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := collection.NewCollection(k, "collection", "index").CreateDocument(id, Document{Title: "yolo"}, nil)
	assert.Equal(t, id, res.Id)
}

func TestCreateDocumentReplace(t *testing.T) {
	type Document struct {
		Title string
	}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "createOrReplace", parsedQuery.Action)

			res := types.Document{Id: "id", Source: []byte(`{"title": "yolo"}`)}
			r, _ := json.Marshal(res)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	newCollection := collection.NewCollection(k, "collection", "index")
	newCollection.CreateDocument("id", Document{Title: "yolo"}, &types.Options{IfExist: "replace"})
}

func TestCreateDocumentCreate(t *testing.T) {
	type Document struct {
		Title string
	}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "create", parsedQuery.Action)

			res := types.Document{Id: "id", Source: []byte(`{"title": "yolo"}`)}
			r, _ := json.Marshal(res)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	newCollection := collection.NewCollection(k, "collection", "index")
	newCollection.CreateDocument("id", Document{Title: "yolo"}, &types.Options{IfExist: "error"})
}

func TestMCreateDocumentError(t *testing.T) {
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

	_, err := collection.NewCollection(k, "collection", "index").MCreateDocument(documents, nil)
	assert.NotNil(t, err)
}

func TestMCreateDocument(t *testing.T) {
	documents := []types.Document{
		{Id: "foo", Source: []byte(`{"title":"Foo"}`)},
		{Id: "bar", Source: []byte(`{"title":"Bar"}`)},
	}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "document", parsedQuery.Controller)
			assert.Equal(t, "mCreate", parsedQuery.Action)
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

	res, _ := collection.NewCollection(k, "collection", "index").MCreateDocument(documents, nil)
	assert.Equal(t, 2, res.Total)

	for index, doc := range res.Hits {
		assert.Equal(t, documents[index].Id, doc.Id)
		assert.Equal(t, documents[index].Source, doc.Source)
	}
}

func TestMCreateOrReplaceDocumentError(t *testing.T) {
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

	_, err := collection.NewCollection(k, "collection", "index").MCreateOrReplaceDocument(documents, nil)
	assert.NotNil(t, err)
}

func TestMCreateOrReplaceDocument(t *testing.T) {
	documents := []types.Document{
		{Id: "foo", Source: []byte(`{"title":"Foo"}`)},
		{Id: "bar", Source: []byte(`{"title":"Bar"}`)},
	}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "document", parsedQuery.Controller)
			assert.Equal(t, "mCreateOrReplace", parsedQuery.Action)
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

	res, _ := collection.NewCollection(k, "collection", "index").MCreateOrReplaceDocument(documents, nil)
	assert.Equal(t, 2, res.Total)

	for index, doc := range res.Hits {
		assert.Equal(t, documents[index].Id, doc.Id)
		assert.Equal(t, documents[index].Source, doc.Source)
	}
}
