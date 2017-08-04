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
	"github.com/kuzzleio/sdk-go/state"
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


func TestDocumentFetchEmptyId(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	dc := collection.NewCollection(k, "collection", "index")
	_, err := dc.CollectionDocument().Fetch("")

	assert.Equal(t, "CollectionDocument.Fetch: missing document id", fmt.Sprint(err))
}

func TestDocumentFetchError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Not found"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	dc := collection.NewCollection(k, "collection", "index")
	_, err := dc.CollectionDocument().Fetch("docId")

	assert.Equal(t, "CollectionDocument.Fetch: an error occurred: Not found", fmt.Sprint(err))
}

func TestDocumentFetch(t *testing.T) {
	id := "docid"

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "document", parsedQuery.Controller)
			assert.Equal(t, "get", parsedQuery.Action)
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
	cd, _ := dc.CollectionDocument().Fetch(id)

	assert.Equal(t, types.Document{Id: id, Source: []byte(`{"foo":"bar"}`)}, cd.Document)
}

func TestDocumentSubscribeEmptyId(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	dc := collection.NewCollection(k, "collection", "index")
	cd := dc.CollectionDocument()

	ch := make(chan types.KuzzleNotification)
	res := <- cd.Subscribe(types.NewRoomOptions(), ch)

	assert.Nil(t, res.Room)
	assert.Equal(t, "CollectionDocument.Subscribe: cannot subscribe to a document if no ID has been provided", fmt.Sprint(res.Error))
}

func TestDocumentSubscribe(t *testing.T) {
	id := "docId"
	var k *kuzzle.Kuzzle

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			// Fetch query
			if parsedQuery.Controller == "document" {
				assert.Equal(t, "get", parsedQuery.Action)
				assert.Equal(t, "index", parsedQuery.Index)
				assert.Equal(t, "collection", parsedQuery.Collection)
				assert.Equal(t, id, parsedQuery.Id)

				res := types.Document{Id: id, Source: []byte(`{"foo":"bar"}`)}
				r, _ := json.Marshal(res)

				return types.KuzzleResponse{Result: r}
			}

			// Subscribe query
			assert.Equal(t, "realtime", parsedQuery.Controller)
			assert.Equal(t, "subscribe", parsedQuery.Action)
			assert.Equal(t, "index", parsedQuery.Index)
			assert.Equal(t, "collection", parsedQuery.Collection)
			assert.Equal(t, map[string]interface {}(map[string]interface {}{"ids":map[string]interface {}{"values":[]interface {}{"docId"}}}), parsedQuery.Body)
			room := collection.NewRoom(*collection.NewCollection(k, "collection", "index"), nil)
			room.RoomId = "42"

			marshed, _ := json.Marshal(room)

			return types.KuzzleResponse{Result: marshed}
		},
	}
	k, _ = kuzzle.NewKuzzle(c, nil)
	*k.State = state.Connected
	dc := collection.NewCollection(k, "collection", "index")
	cd, _ := dc.CollectionDocument().Fetch(id)

	ch := make(chan types.KuzzleNotification)
	subRes := cd.Subscribe(types.NewRoomOptions(), ch)
	r := <-subRes

	assert.Nil(t, r.Error)
	assert.NotNil(t, r.Room)
	assert.Equal(t, "42", r.Room.GetRoomId())
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
