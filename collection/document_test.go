package collection_test

import (
	"encoding/json"
	"fmt"
	"github.com/kuzzleio/sdk-go/collection"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/state"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
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

func ExampleDocument_SetContent() {
	c := &internal.MockedConnection{}
	k, _ := kuzzle.NewKuzzle(c, nil)
	dc := collection.NewCollection(k, "collection", "index")

	d := dc.Document()

	d = d.SetContent(collection.DocumentContent{
		"subfield": collection.DocumentContent{
			"john": "cena",
		},
	}, false)

	fmt.Println(d.Content)
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

	headers := make(map[string]interface{})

	assert.Equal(t, headers, k.GetHeaders())

	headers["foo"] = "bar"
	headers["bar"] = "foo"

	d.SetHeaders(headers, false)

	newHeaders := make(map[string]interface{})
	newHeaders["foo"] = "rab"

	d.SetHeaders(newHeaders, false)

	headers["foo"] = "rab"

	assert.Equal(t, headers, k.GetHeaders())
	assert.NotEqual(t, newHeaders, k.GetHeaders())
}

func ExampleDocument_SetHeaders() {
	c := &internal.MockedConnection{}
	k, _ := kuzzle.NewKuzzle(c, nil)
	d := collection.NewCollection(k, "collection", "index").Document()

	headers := make(map[string]interface{})

	headers["foo"] = "bar"
	headers["bar"] = "foo"

	d.SetHeaders(headers, true)

	fmt.Println(k.GetHeaders())
}

func ExampleDocument_SetDocumentId() {
	c := &internal.MockedConnection{}
	k, _ := kuzzle.NewKuzzle(c, nil)
	d := collection.NewCollection(k, "collection", "index").Document()

	d.SetDocumentId("newId")

	fmt.Println(d.Id)
}

func TestDocumentSetHeadersReplace(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	d := collection.NewCollection(k, "collection", "index").Document()

	headers := make(map[string]interface{})

	assert.Equal(t, headers, k.GetHeaders())

	headers["foo"] = "bar"
	headers["bar"] = "foo"

	d.SetHeaders(headers, false)

	newHeaders := make(map[string]interface{})
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

	assert.Equal(t, "[400] Document.Fetch: missing document id", fmt.Sprint(err))
}

func TestDocumentFetchError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "Not found"}}
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
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "document", parsedQuery.Controller)
			assert.Equal(t, "get", parsedQuery.Action)
			assert.Equal(t, "index", parsedQuery.Index)
			assert.Equal(t, "collection", parsedQuery.Collection)
			assert.Equal(t, id, parsedQuery.Id)

			res := collection.Document{Id: id, Content: []byte(`{"foo":"bar"}`)}
			r, _ := json.Marshal(res)
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	dc := collection.NewCollection(k, "collection", "index")
	d, _ := dc.Document().Fetch(id)
	r := collection.Document{Id: id, Content: []byte(`{"foo":"bar"}`)}

	assert.Equal(t, r.Id, d.Id)
	assert.Equal(t, r.Content, d.Content)
}

func ExampleDocument_Fetch() {
	id := "docid"
	c := &internal.MockedConnection{}
	k, _ := kuzzle.NewKuzzle(c, nil)
	dc := collection.NewCollection(k, "collection", "index")
	res, err := dc.Document().Fetch(id)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res.Id, res.Collection)
}

func TestDocumentSubscribeEmptyId(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	dc := collection.NewCollection(k, "collection", "index")
	cd := dc.Document()

	ch := make(chan *types.KuzzleNotification)
	res := <-cd.Subscribe(types.NewRoomOptions(), ch)

	assert.Nil(t, res.Room)
	assert.Equal(t, "[400] Document.Subscribe: cannot subscribe to a document if no ID has been provided", fmt.Sprint(res.Error))
}

func TestDocumentSubscribe(t *testing.T) {
	id := "docId"
	var k *kuzzle.Kuzzle

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
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

				return &types.KuzzleResponse{Result: r}
			}

			// Subscribe query
			assert.Equal(t, "realtime", parsedQuery.Controller)
			assert.Equal(t, "subscribe", parsedQuery.Action)
			assert.Equal(t, "index", parsedQuery.Index)
			assert.Equal(t, "collection", parsedQuery.Collection)
			assert.Equal(t, map[string]interface{}(map[string]interface{}{"ids": map[string]interface{}{"values": []interface{}{"docId"}}}), parsedQuery.Body)
			room := collection.NewRoom(collection.NewCollection(k, "collection", "index"), nil)
			room.RoomId = "42"

			marshed, _ := json.Marshal(room)

			return &types.KuzzleResponse{Result: marshed}
		},
	}
	k, _ = kuzzle.NewKuzzle(c, nil)
	*k.State = state.Connected
	dc := collection.NewCollection(k, "collection", "index")
	d, _ := dc.Document().Fetch(id)

	ch := make(chan *types.KuzzleNotification)
	subRes := d.Subscribe(types.NewRoomOptions(), ch)
	r := <-subRes

	assert.Nil(t, r.Error)
	assert.NotNil(t, r.Room)
	assert.Equal(t, "42", r.Room.GetRoomId())
}

func ExampleDocument_Subscribe() {
	id := "docId"
	var k *kuzzle.Kuzzle

	c := &internal.MockedConnection{}
	k, _ = kuzzle.NewKuzzle(c, nil)
	*k.State = state.Connected
	dc := collection.NewCollection(k, "collection", "index")
	d, _ := dc.Document().Fetch(id)

	ch := make(chan *types.KuzzleNotification)
	subRes := d.Subscribe(types.NewRoomOptions(), ch)

	notification := <-subRes

	fmt.Println(notification.Room.GetRoomId(), notification.Error)
}

func TestDocumentSaveEmptyId(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	dc := collection.NewCollection(k, "collection", "index")
	_, err := dc.Document().Save(nil)

	assert.NotNil(t, err)
	assert.Equal(t, "[400] Document.Save: missing document id", fmt.Sprint(err))
}

func TestDocumentSaveError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "Unit test error"}}
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
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "document", parsedQuery.Controller)
			assert.Equal(t, "createOrReplace", parsedQuery.Action)
			assert.Equal(t, "index", parsedQuery.Index)
			assert.Equal(t, "collection", parsedQuery.Collection)
			assert.Equal(t, id, parsedQuery.Id)

			res := collection.Document{Id: id, Content: []byte(`{"foo":"bar"}`)}
			r, _ := json.Marshal(res)
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	dc := collection.NewCollection(k, "collection", "index")

	documentContent := collection.DocumentContent{"foo": "bar"}

	d, _ := dc.Document().SetDocumentId(id).SetContent(documentContent, true).Save(nil)

	assert.Equal(t, id, d.Id)
	assert.Equal(t, documentContent.ToString(), string(d.Content))
}

func ExampleDocument_Save() {
	id := "myId"
	c := &internal.MockedConnection{}
	k, _ := kuzzle.NewKuzzle(c, nil)
	dc := collection.NewCollection(k, "collection", "index")

	documentContent := collection.DocumentContent{"foo": "bar"}

	res, err := dc.Document().SetDocumentId(id).SetContent(documentContent, true).Save(nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res.Id, res.Collection)
}

func TestDocumentRefreshEmptyId(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	dc := collection.NewCollection(k, "collection", "index")
	_, err := dc.Document().Refresh(nil)

	assert.NotNil(t, err)
	assert.Equal(t, "[400] Document.Refresh: missing document id", fmt.Sprint(err))
}

func TestDocumentRefreshError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "Unit test error"}}
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
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "document", parsedQuery.Controller)
			assert.Equal(t, "get", parsedQuery.Action)
			assert.Equal(t, "index", parsedQuery.Index)
			assert.Equal(t, "collection", parsedQuery.Collection)
			assert.Equal(t, id, parsedQuery.Id)

			res := collection.Document{Id: id, Content: []byte(`{"name":"Anakin","function":"Jedi"}`)}
			r, _ := json.Marshal(res)
			return &types.KuzzleResponse{Result: r}
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

func ExampleDocument_Refresh() {
	id := "myId"
	c := &internal.MockedConnection{}
	k, _ := kuzzle.NewKuzzle(c, nil)
	dc := collection.NewCollection(k, "collection", "index")
	res, err := dc.Document().SetDocumentId(id).Refresh(nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res.Id, res.Collection)
}

func TestCollectionDocumentExistsEmptyId(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	dc := collection.NewCollection(k, "collection", "index")
	_, err := dc.Document().Exists(nil)

	assert.NotNil(t, err)
	assert.Equal(t, "[400] Document.Exists: missing document id", fmt.Sprint(err))
}

func TestCollectionDocumentExistsError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "Unit test error"}}
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
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "document", parsedQuery.Controller)
			assert.Equal(t, "exists", parsedQuery.Action)
			assert.Equal(t, "index", parsedQuery.Index)
			assert.Equal(t, "collection", parsedQuery.Collection)
			assert.Equal(t, id, parsedQuery.Id)

			r, _ := json.Marshal(true)
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	dc := collection.NewCollection(k, "collection", "index")
	exists, _ := dc.Document().SetDocumentId(id).Exists(nil)

	assert.Equal(t, true, exists)
}

func ExampleDocument_Exists() {
	id := "myId"
	c := &internal.MockedConnection{}
	k, _ := kuzzle.NewKuzzle(c, nil)
	dc := collection.NewCollection(k, "collection", "index")
	res, err := dc.Document().SetDocumentId(id).Exists(nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res)
}

func TestDocumentPublishError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	dc := collection.NewCollection(k, "realtime", "publish")
	_, err := dc.Document().SetDocumentId("myId").Publish(nil)

	assert.NotNil(t, err)
}

func TestDocumentPublish(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "realtime", parsedQuery.Controller)
			assert.Equal(t, "publish", parsedQuery.Action)
			assert.Equal(t, "index", parsedQuery.Index)
			assert.Equal(t, "collection", parsedQuery.Collection)

			r, _ := json.Marshal(types.RealtimeResponse{Published: true})
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	dc := collection.NewCollection(k, "collection", "index")
	res, _ := dc.Document().SetDocumentId("myId").Publish(nil)

	assert.Equal(t, true, res)
}

func ExampleDocument_Publish() {
	c := &internal.MockedConnection{}
	k, _ := kuzzle.NewKuzzle(c, nil)
	dc := collection.NewCollection(k, "collection", "index")
	res, err := dc.Document().SetDocumentId("myId").Publish(nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res)
}

func TestDocumentDeleteEmptyId(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	dc := collection.NewCollection(k, "collection", "index")
	_, err := dc.Document().Delete(nil)

	assert.NotNil(t, err)
	assert.Equal(t, "[400] Document.Delete: missing document id", fmt.Sprint(err))
}

func TestDocumentDeleteError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "Unit test error"}}
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
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "document", parsedQuery.Controller)
			assert.Equal(t, "delete", parsedQuery.Action)
			assert.Equal(t, "index", parsedQuery.Index)
			assert.Equal(t, "collection", parsedQuery.Collection)
			assert.Equal(t, id, parsedQuery.Id)

			r, _ := json.Marshal(collection.Document{Id: id})
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	dc := collection.NewCollection(k, "collection", "index")
	res, _ := dc.Document().SetDocumentId("myId").Delete(nil)

	assert.Equal(t, id, res)
}

func ExampleDocument_Delete() {
	c := &internal.MockedConnection{}
	k, _ := kuzzle.NewKuzzle(c, nil)
	dc := collection.NewCollection(k, "collection", "index")
	res, err := dc.Document().SetDocumentId("myId").Delete(nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res)
}
