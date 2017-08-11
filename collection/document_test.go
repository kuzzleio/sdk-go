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

	d := dc.Document()
	d.Content = []byte(`{"foo":"bar","subfield":{"john":"smith"}}`)

	assert.Equal(t, json.RawMessage([]byte(`{"foo":"bar","subfield":{"john":"smith"}}`)), d.Content)

	d = d.SetContent(collection.DocumentContent{
		"subfield": collection.DocumentContent{
			"john": "cena",
		},
	}, false)

	assert.Equal(t, string(json.RawMessage([]byte(`{"foo":"bar","subfield":{"john":"cena"}}`))), string(d.Content))
}

func TestDocumentSetContentReplace(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	dc := collection.NewCollection(k, "collection", "index")

	d := dc.Document()
	d.Content = []byte(`{"foo":"bar","subfield":{"john":"smith"}}`)

	assert.Equal(t, json.RawMessage([]byte(`{"foo":"bar","subfield":{"john":"smith"}}`)), d.Content)

	d = d.SetContent(collection.DocumentContent{
		"subfield": collection.DocumentContent{
			"john": "cena",
			"subsubfield": collection.DocumentContent{
				"hi": "there",
			},
		},
	}, true)

	assert.Equal(t, string(json.RawMessage([]byte(`{"subfield":{"john":"cena","subsubfield":{"hi":"there"}}}`))), string(d.Content))
}

func TestDocumentSetHeaders(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	d := collection.NewCollection(k, "collection", "index").Document()

	var headers = make(map[string]interface{}, 0)

	assert.Equal(t, headers, k.GetHeaders())

	headers["foo"] = "bar"
	headers["bar"] = "foo"

	d.SetHeaders(headers, false)

	var newHeaders = make(map[string]interface{}, 0)
	newHeaders["foo"] = "rab"

	d.SetHeaders(newHeaders, false)

	headers["foo"] = "rab"

	assert.Equal(t, headers, k.GetHeaders())
	assert.NotEqual(t, newHeaders, k.GetHeaders())
}

func TestDocumentSetHeadersReplace(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	d := collection.NewCollection(k, "collection", "index").Document()

	var headers = make(map[string]interface{}, 0)

	assert.Equal(t, headers, k.GetHeaders())

	headers["foo"] = "bar"
	headers["bar"] = "foo"

	d.SetHeaders(headers, false)

	var newHeaders = make(map[string]interface{}, 0)
	newHeaders["foo"] = "rab"

	d.SetHeaders(newHeaders, true)

	headers["foo"] = "rab"

	assert.Equal(t, newHeaders, k.GetHeaders())
	assert.NotEqual(t, headers, k.GetHeaders())
}


func TestDocumentFetchEmptyId(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	dc := collection.NewCollection(k, "collection", "index")
	_, err := dc.Document().Fetch("")

	assert.Equal(t, "Document.Fetch: missing document id", fmt.Sprint(err))
}

func TestDocumentFetchError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Not found"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	dc := collection.NewCollection(k, "collection", "index")
	_, err := dc.Document().Fetch("docId")

	assert.Equal(t, "Document.Fetch: an error occurred: Not found", fmt.Sprint(err))
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

			res := collection.Document{Id: id, Content: []byte(`{"foo":"bar"}`)}
			r, _ := json.Marshal(res)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	dc := collection.NewCollection(k, "collection", "index")
	d, _ := dc.Document().Fetch(id)

	assert.Equal(t, id, d.Id)
	assert.Equal(t, []byte(`{"foo":"bar"}`), d.Content)
}

func TestDocumentSubscribeEmptyId(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	dc := collection.NewCollection(k, "collection", "index")
	cd := dc.Document()

	ch := make(chan types.KuzzleNotification)
	res := <- cd.Subscribe(types.NewRoomOptions(), ch)

	assert.Nil(t, res.Room)
	assert.Equal(t, "Document.Subscribe: cannot subscribe to a document if no ID has been provided", fmt.Sprint(res.Error))
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

				res := collection.Document{Id: id, Content: []byte(`{"foo":"bar"}`)}
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
	cd, _ := dc.Document().Fetch(id)

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
	_, err := dc.Document().Save(nil)

	assert.NotNil(t, err)
	assert.Equal(t, "Document.Save: missing document id", fmt.Sprint(err))
}

func TestDocumentSaveError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	dc := collection.NewCollection(k, "collection", "index")
	_, err := dc.Document().SetDocumentId("myId").Save(nil)

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

			res := collection.Document{Id: id, Content: []byte(`{"foo":"bar"}`)}
			r, _ := json.Marshal(res)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	dc := collection.NewCollection(k, "collection", "index")

	documentContent := collection.DocumentContent{"foo": "bar"}

	d, _ := dc.Document().SetDocumentId(id).SetContent(documentContent, true).Save(nil)

	assert.Equal(t, id, d.Id)
	assert.Equal(t, documentContent.ToString(), string(d.Content))
}

func TestDocumentRefreshEmptyId(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	dc := collection.NewCollection(k, "collection", "index")
	_, err := dc.Document().Refresh(nil)

	assert.NotNil(t, err)
	assert.Equal(t, "Document.Refresh: missing document id", fmt.Sprint(err))
}

func TestDocumentRefreshError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	dc := collection.NewCollection(k, "collection", "index")
	_, err := dc.Document().SetDocumentId("myId").Refresh(nil)

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

			res := collection.Document{Id: id, Content: []byte(`{"name":"Anakin","function":"Jedi"}`)}
			r, _ := json.Marshal(res)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	dc := collection.NewCollection(k, "collection", "index")

	documentContent := collection.DocumentContent{
		"name":     "Anakin",
		"function": "Padawan",
	}

	d, _ := dc.Document().SetDocumentId(id).SetContent(documentContent, true).Refresh(nil)

	result := collection.Document{}
	json.Unmarshal(d.Content, &result.Content)

	ic := collection.DocumentContent{}
	json.Unmarshal(result.Content, &ic)

	assert.Equal(t, id, d.Id)
	assert.Equal(t, "Padawan", documentContent["function"])
	assert.Equal(t, "Jedi", ic["function"])
	assert.NotEqual(t, documentContent["function"], ic["function"])
}

func TestCollectionDocumentExistsEmptyId(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	dc := collection.NewCollection(k, "collection", "index")
	_, err := dc.Document().Exists(nil)

	assert.NotNil(t, err)
	assert.Equal(t, "Document.Exists: missing document id", fmt.Sprint(err))
}

func TestCollectionDocumentExistsError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	dc := collection.NewCollection(k, "collection", "index")
	_, err := dc.Document().SetDocumentId("myId").Exists(nil)

	assert.NotNil(t, err)
}

func TestCollectionDocumentExists(t *testing.T) {
	id := "myId"

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "document", parsedQuery.Controller)
			assert.Equal(t, "exists", parsedQuery.Action)
			assert.Equal(t, "index", parsedQuery.Index)
			assert.Equal(t, "collection", parsedQuery.Collection)
			assert.Equal(t, id, parsedQuery.Id)

			r, _ := json.Marshal(true)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	dc := collection.NewCollection(k, "collection", "index")
	exists, _ := dc.Document().SetDocumentId("myId").Exists(nil)

	assert.Equal(t, true, exists)
}

func TestDocumentPublishError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	dc := collection.NewCollection(k, "realtime", "publish")
	_, err := dc.Document().SetDocumentId("myId").Publish(nil)

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
	result, _ := dc.Document().SetDocumentId("myId").Publish(nil)

	assert.Equal(t, true, result)
}

func TestDocumentDeleteEmptyId(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	dc := collection.NewCollection(k, "collection", "index")
	_, err := dc.Document().Delete(nil)

	assert.NotNil(t, err)
	assert.Equal(t, "Document.Delete: missing document id", fmt.Sprint(err))
}

func TestDocumentDeleteError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	dc := collection.NewCollection(k, "collection", "index")
	_, err := dc.Document().SetDocumentId("myId").Delete(nil)

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

			r, _ := json.Marshal(collection.Document{Id: id})
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	dc := collection.NewCollection(k, "collection", "index")
	result, _ := dc.Document().SetDocumentId("myId").Delete(nil)

	assert.Equal(t, id, result)
}
