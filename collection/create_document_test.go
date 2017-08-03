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
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := collection.NewCollection(k, "collection", "index").CreateDocument("id", types.Document{Source: []byte(`{"title":"yolo"}`)}, nil)
	assert.NotNil(t, err)
}

func TestCreateDocumentWrongOptionError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	newCollection := collection.NewCollection(k, "collection", "index")
	opts := types.NewQueryOptions()
	opts.SetIfExist("unknown")

	_, err := newCollection.CreateDocument("id", types.Document{Source: []byte(`{"title":"yolo"}`)}, opts)
	assert.Equal(t, "Invalid value for the 'ifExist' option: 'unknown'", fmt.Sprint(err))
}

func TestCreateDocument(t *testing.T) {
	id := "myId"

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "document", parsedQuery.Controller)
			assert.Equal(t, "create", parsedQuery.Action)
			assert.Equal(t, "index", parsedQuery.Index)
			assert.Equal(t, "collection", parsedQuery.Collection)
			assert.Equal(t, id, parsedQuery.Id)

			var body = make(map[string]interface{}, 0)
			body["title"] = "yolo"

			assert.Equal(t, body, parsedQuery.Body)

			res := types.Document{Id: id, Source: []byte(`{"title":"yolo"}`)}
			r, _ := json.Marshal(res)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := collection.NewCollection(k, "collection", "index").CreateDocument(id, types.Document{Source: []byte(`{"title":"yolo"}`)}, nil)
	assert.Equal(t, id, res.Id)
}

func TestCreateDocumentReplace(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "createOrReplace", parsedQuery.Action)

			res := types.Document{Id: "id", Source: []byte(`{"Title":"yolo"}`)}
			r, _ := json.Marshal(res)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	newCollection := collection.NewCollection(k, "collection", "index")
	opts := types.NewQueryOptions()
	opts.SetIfExist("replace")
	newCollection.CreateDocument("id", types.Document{Source: []byte(`{"Title":"yolo"}`)}, opts)
}

func TestCreateDocumentCreate(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "create", parsedQuery.Action)

			res := types.Document{Id: "id", Source: []byte(`{"Title":"yolo"}`)}
			r, _ := json.Marshal(res)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	newCollection := collection.NewCollection(k, "collection", "index")
	opts := types.NewQueryOptions()
	opts.SetIfExist("error")

	newCollection.CreateDocument("id", types.Document{Source: []byte(`{"Title":"yolo"}`)}, opts)
}

func TestMCreateDocumentError(t *testing.T) {
	documents := []types.Document{
		{Id: "foo", Source: []byte(`{"title":"Foo"}`)},
		{Id: "bar", Source: []byte(`{"title":"Bar"}`)},
	}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
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
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
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
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
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
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
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
