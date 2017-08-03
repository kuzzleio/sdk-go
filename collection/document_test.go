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

func TestDocumentSetContent(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	dc := collection.NewCollection(k, "collection", "index")

	cd := dc.CollectionDocument()
	cd.Document.Source = []byte(`{"foo":"bar","subfield":{"john":"smith"}}`)

	assert.Equal(t, json.RawMessage([]byte(`{"foo":"bar","subfield":{"john":"smith"}}`)), cd.Document.Source)

	cd = cd.SetContent(collection.DocumentContent{
		"subfield": collection.DocumentContent{
			"john": "cena",
		},
	}, false)

	assert.Equal(t, string(json.RawMessage([]byte(`{"foo":"bar","subfield":{"john":"cena"}}`))), string(cd.Document.Source))
}

func TestDocumentSetContentReplace(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	dc := collection.NewCollection(k, "collection", "index")

	cd := dc.CollectionDocument()
	cd.Document.Source = []byte(`{"foo":"bar","subfield":{"john":"smith"}}`)

	assert.Equal(t, json.RawMessage([]byte(`{"foo":"bar","subfield":{"john":"smith"}}`)), cd.Document.Source)

	cd = cd.SetContent(collection.DocumentContent{
		"subfield": collection.DocumentContent{
			"john": "cena",
			"subsubfield": collection.DocumentContent{
				"hi": "there",
			},
		},
	}, true)

	assert.Equal(t, string(json.RawMessage([]byte(`{"subfield":{"john":"cena","subsubfield":{"hi":"there"}}}`))), string(cd.Document.Source))
}

func TestDocumentSetHeaders(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	cd := collection.NewCollection(k, "collection", "index").CollectionDocument()

	var headers = make(map[string]interface{}, 0)

	assert.Equal(t, headers, k.GetHeaders())

	headers["foo"] = "bar"
	headers["bar"] = "foo"

	cd.SetHeaders(headers, false)

	var newHeaders = make(map[string]interface{}, 0)
	newHeaders["foo"] = "rab"

	cd.SetHeaders(newHeaders, false)

	headers["foo"] = "rab"

	assert.Equal(t, headers, k.GetHeaders())
	assert.NotEqual(t, newHeaders, k.GetHeaders())
}

func TestDocumentSetHeadersReplace(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	cd := collection.NewCollection(k, "collection", "index").CollectionDocument()

	var headers = make(map[string]interface{}, 0)

	assert.Equal(t, headers, k.GetHeaders())

	headers["foo"] = "bar"
	headers["bar"] = "foo"

	cd.SetHeaders(headers, false)

	var newHeaders = make(map[string]interface{}, 0)
	newHeaders["foo"] = "rab"

	cd.SetHeaders(newHeaders, true)

	headers["foo"] = "rab"

	assert.Equal(t, newHeaders, k.GetHeaders())
	assert.NotEqual(t, headers, k.GetHeaders())
}

func TestDocumentSaveEmptyId(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	dc := collection.NewCollection(k, "collection", "index")
	_, err := dc.CollectionDocument().Save(nil)

	assert.NotNil(t, err)
	assert.Equal(t, "CollectionDocument.Save: missing document id", fmt.Sprint(err))
}

func TestDocumentSaveError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	dc := collection.NewCollection(k, "collection", "index")
	_, err := dc.CollectionDocument().SetDocumentId("myId").Save(nil)

	assert.NotNil(t, err)
}

func TestDocumentSave(t *testing.T) {
	id := "myId"

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "document", parsedQuery.Controller)
			assert.Equal(t, "createOrReplace", parsedQuery.Action)
			assert.Equal(t, "index", parsedQuery.Index)
			assert.Equal(t, "collection", parsedQuery.Collection)
			assert.Equal(t, id, parsedQuery.Id)

			res := types.Document{Id: id, Source: []byte(`{"foo":"bar"}`)}
			r, _ := json.Marshal(res)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	dc := collection.NewCollection(k, "collection", "index")

	documentSource := collection.DocumentContent{"foo": "bar"}

	cd, _ := dc.CollectionDocument().SetDocumentId(id).SetContent(documentSource, true).Save(nil)

	assert.Equal(t, id, cd.Document.Id)
	assert.Equal(t, dc, &cd.Collection)
	assert.Equal(t, documentSource.ToString(), string(cd.Document.Source))
}

func TestDocumentRefreshEmptyId(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	dc := collection.NewCollection(k, "collection", "index")
	_, err := dc.CollectionDocument().Refresh(nil)

	assert.NotNil(t, err)
	assert.Equal(t, "CollectionDocument.Refresh: missing document id", fmt.Sprint(err))
}

func TestDocumentRefreshError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	dc := collection.NewCollection(k, "collection", "index")
	_, err := dc.CollectionDocument().SetDocumentId("myId").Refresh(nil)

	assert.NotNil(t, err)
}

func TestDocumentRefresh(t *testing.T) {
	id := "myId"

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "document", parsedQuery.Controller)
			assert.Equal(t, "get", parsedQuery.Action)
			assert.Equal(t, "index", parsedQuery.Index)
			assert.Equal(t, "collection", parsedQuery.Collection)
			assert.Equal(t, id, parsedQuery.Id)

			res := types.Document{Id: id, Source: []byte(`{"name":"Anakin","function":"Jedi"}`)}
			r, _ := json.Marshal(res)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	dc := collection.NewCollection(k, "collection", "index")

	documentSource := collection.DocumentContent{
		"name":     "Anakin",
		"function": "Padawan",
	}

	cd, _ := dc.CollectionDocument().SetDocumentId(id).SetContent(documentSource, true).Refresh(nil)

	result := types.Document{}
	json.Unmarshal(cd.Document.Source, &result.Source)

	ic := collection.DocumentContent{}
	json.Unmarshal(result.Source, &ic)

	assert.Equal(t, id, cd.Document.Id)
	assert.Equal(t, dc, &cd.Collection)
	assert.Equal(t, "Padawan", documentSource["function"])
	assert.Equal(t, "Jedi", ic["function"])
	assert.NotEqual(t, documentSource["function"], ic["function"])
}

func TestDocumentPublishError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	dc := collection.NewCollection(k, "realtime", "publish")
	_, err := dc.CollectionDocument().SetDocumentId("myId").Publish(nil)

	assert.NotNil(t, err)
}

func TestDocumentPublish(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "realtime", parsedQuery.Controller)
			assert.Equal(t, "publish", parsedQuery.Action)
			assert.Equal(t, "index", parsedQuery.Index)
			assert.Equal(t, "collection", parsedQuery.Collection)

			r, _ := json.Marshal(types.RealtimeResponse{Published: true})
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	dc := collection.NewCollection(k, "collection", "index")
	result, _ := dc.CollectionDocument().SetDocumentId("myId").Publish(nil)

	assert.Equal(t, true, result)
}

func TestDocumentDeleteEmptyId(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	dc := collection.NewCollection(k, "collection", "index")
	_, err := dc.CollectionDocument().Delete(nil)

	assert.NotNil(t, err)
	assert.Equal(t, "CollectionDocument.Delete: missing document id", fmt.Sprint(err))
}

func TestDocumentDeleteError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	dc := collection.NewCollection(k, "collection", "index")
	_, err := dc.CollectionDocument().SetDocumentId("myId").Delete(nil)

	assert.NotNil(t, err)
}

func TestDocumentDelete(t *testing.T) {
	id := "myId"

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "document", parsedQuery.Controller)
			assert.Equal(t, "delete", parsedQuery.Action)
			assert.Equal(t, "index", parsedQuery.Index)
			assert.Equal(t, "collection", parsedQuery.Collection)
			assert.Equal(t, id, parsedQuery.Id)

			r, _ := json.Marshal(types.Document{Id: id})
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	dc := collection.NewCollection(k, "collection", "index")
	result, _ := dc.CollectionDocument().SetDocumentId("myId").Delete(nil)

	assert.Equal(t, id, result)
}
