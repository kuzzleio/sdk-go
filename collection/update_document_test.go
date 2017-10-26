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

func TestUpdateDocumentEmptyId(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "Collection.UpdateDocument: document id required"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := collection.NewCollection(k, "collection", "index").UpdateDocument("", &collection.Document{Content: []byte(`{"title": "jonathan"}`)}, nil)
	assert.NotNil(t, err)
}

func TestUpdateDocumentError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := collection.NewCollection(k, "collection", "index").UpdateDocument("id", &collection.Document{Content: []byte(`{"title": "jonathan"}`)}, nil)
	assert.NotNil(t, err)
}

func TestUpdateDocument(t *testing.T) {
	id := "myId"

	type Content struct {
		Title string `json:"title"`
	}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "document", parsedQuery.Controller)
			assert.Equal(t, "update", parsedQuery.Action)
			assert.Equal(t, "index", parsedQuery.Index)
			assert.Equal(t, "collection", parsedQuery.Collection)
			assert.Equal(t, 10, options.GetRetryOnConflict())
			assert.Equal(t, id, parsedQuery.Id)

			assert.Equal(t, "jonathan", parsedQuery.Body.(map[string]interface{})["title"])

			res := collection.Document{Id: id, Content: []byte(`{"title": "arthur"}`)}
			r, _ := json.Marshal(res)
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	qo := types.NewQueryOptions()
	qo.SetRetryOnConflict(10)

	res, _ := collection.NewCollection(k, "collection", "index").UpdateDocument(id, &collection.Document{Content: []byte(`{"title": "jonathan"}`)}, qo)

	assert.Equal(t, id, res.Id)

	var result Content

	json.Unmarshal(res.Content, &result)

	assert.Equal(t, result.Title, "arthur")
}

func ExampleCollection_UpdateDocument() {
	id := "myId"
	c := &internal.MockedConnection{}
	k, _ := kuzzle.NewKuzzle(c, nil)
	qo := types.NewQueryOptions()
	qo.SetRetryOnConflict(10)

	res, err := collection.NewCollection(k, "collection", "index").UpdateDocument(id, &collection.Document{Content: []byte(`{"title": "jonathan"}`)}, qo)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res.Id, res.Collection)
}

func TestMUpdateDocumentEmptyDocuments(t *testing.T) {
	documents := []*collection.Document{}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "Collection.MUpdateDocument: please provide at least one document to update"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := collection.NewCollection(k, "collection", "index").MUpdateDocument(documents, nil)
	assert.NotNil(t, err)
}

func TestMUpdateDocumentError(t *testing.T) {
	documents := []*collection.Document{
		{Id: "foo", Content: []byte(`{"title":"Foo"}`)},
		{Id: "bar", Content: []byte(`{"title":"Bar"}`)},
	}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := collection.NewCollection(k, "collection", "index").MUpdateDocument(documents, nil)
	assert.NotNil(t, err)
}

func TestMUpdateDocument(t *testing.T) {
	documents := []*collection.Document{
		{Id: "foo", Content: []byte(`{"title":"Foo"}`)},
		{Id: "bar", Content: []byte(`{"title":"Bar"}`)},
	}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "document", parsedQuery.Controller)
			assert.Equal(t, "mUpdate", parsedQuery.Action)
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

	res, _ := collection.NewCollection(k, "collection", "index").MUpdateDocument(documents, nil)
	assert.Equal(t, 2, res.Total)

	for index, doc := range res.Hits {
		assert.Equal(t, documents[index].Id, doc.Id)
		assert.Equal(t, documents[index].Content, doc.Content)
	}
}

func ExampleCollection_MUpdateDocument() {
	documents := []*collection.Document{
		{Id: "foo", Content: []byte(`{"title":"Foo"}`)},
		{Id: "bar", Content: []byte(`{"title":"Bar"}`)},
	}

	c := &internal.MockedConnection{}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, err := collection.NewCollection(k, "collection", "index").MUpdateDocument(documents, nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res.Hits[0].Id, res.Hits[0].Collection)
}
