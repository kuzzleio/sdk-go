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

func TestFetchDocumentEmptyId(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Collection.FetchDocument: document id required"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := collection.NewCollection(k, "collection", "index").FetchDocument("", nil)
	assert.NotNil(t, err)
}

func TestFetchDocumentError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Unit test error"}}
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
		MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "document", parsedQuery.Controller)
			assert.Equal(t, "get", parsedQuery.Action)
			assert.Equal(t, "index", parsedQuery.Index)
			assert.Equal(t, "collection", parsedQuery.Collection)
			assert.Equal(t, id, parsedQuery.Id)

			res := types.Document{Id: id, Source: []byte(`{"foo": "bar"}`)}
			r, _ := json.Marshal(res)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := collection.NewCollection(k, "collection", "index").FetchDocument(id, nil)
	assert.Equal(t, id, res.Id)
}

func TestMGetDocumentEmptyIds(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Collection.MGetDocument: please provide at least one id of document to retrieve"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := collection.NewCollection(k, "collection", "index").MGetDocument([]string{}, nil)
	assert.NotNil(t, err)
}

func TestMGetDocumentError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := collection.NewCollection(k, "collection", "index").MGetDocument([]string{"foo", "bar"}, nil)
	assert.NotNil(t, err)
}

func TestMGetDocument(t *testing.T) {
	hits := make([]types.KuzzleResult, 2)
	hits[0] = types.KuzzleResult{Id: "foo", Source: json.RawMessage(`{"title":"foo"}`)}
	hits[1] = types.KuzzleResult{Id: "bar", Source: json.RawMessage(`{"title":"bar"}`)}
	var results = types.KuzzleSearchResult{Total: 2, Hits: hits}

	ids := []string{"foo", "bar"}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "document", parsedQuery.Controller)
			assert.Equal(t, "mGet", parsedQuery.Action)
			assert.Equal(t, "index", parsedQuery.Index)
			assert.Equal(t, "collection", parsedQuery.Collection)
			assert.Equal(t, []interface{}{"foo", "bar"}, parsedQuery.Body.(map[string]interface{})["ids"])

			res := types.KuzzleSearchResult{Total: results.Total, Hits: results.Hits}
			r, _ := json.Marshal(res)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := collection.NewCollection(k, "collection", "index").MGetDocument(ids, nil)
	assert.Equal(t, results.Total, res.Total)
	assert.Equal(t, hits, res.Hits)
}
