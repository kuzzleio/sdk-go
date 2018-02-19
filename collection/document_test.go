package collection_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/kuzzleio/sdk-go/collection"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/state"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
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

	assert.Equal(t, string(json.RawMessage([]byte(`{"foo":"bar","subfield":{"john":"smith"}}`))), string(d.Content))
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

func TestDocumentSubscribeEmptyId(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	dc := collection.NewCollection(k, "collection", "index")
	cd := dc.Document()

	ch := make(chan types.KuzzleNotification)
	_, err := cd.Subscribe(types.NewRoomOptions(), ch)

	assert.Equal(t, "[400] Document.Subscribe: cannot subscribe to a document if no ID has been provided", err.Error())
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
			roomRaw := []byte(`{"requestId": "rqid", "channel": "foo", "roomId": "42"}`)
			return &types.KuzzleResponse{Result: roomRaw}
		},
	}
	k, _ = kuzzle.NewKuzzle(c, nil)
	c.SetState(state.Connected)
	dc := collection.NewCollection(k, "collection", "index")
	d := dc.Document()
	d.Id = id

	ch := make(chan types.KuzzleNotification)
	room, _ := d.Subscribe(types.NewRoomOptions(), ch)
	r := <-room.ResponseChannel()

	assert.Nil(t, r.Error)
	assert.NotNil(t, r.Room)
	assert.Equal(t, "42", r.Room.RoomId())
}

func ExampleDocument_Subscribe() {
	id := "docId"
	var k *kuzzle.Kuzzle

	c := &internal.MockedConnection{}
	k, _ = kuzzle.NewKuzzle(c, nil)
	c.SetState(state.Connected)
	dc := collection.NewCollection(k, "collection", "index")
	d := dc.Document()
	d.Id = id

	ch := make(chan types.KuzzleNotification)
	room, _ := d.Subscribe(types.NewRoomOptions(), ch)

	notification := <-room.ResponseChannel()

	fmt.Println(notification.Room.RoomId(), notification.Error)
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

//Create tests
func TestDocumentCreateError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: types.NewError("Unit test error")}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	col := collection.NewCollection(k, "collection", "index")
	doc := collection.NewDocument(col, "")
	_, err := doc.Create(nil)
	assert.NotNil(t, err)
	assert.Equal(t, "Unit test error", err.(*types.KuzzleError).Message)
}

func TestCreateDocumentWrongOptionError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	newCollection := collection.NewCollection(k, "collection", "index")
	opts := types.NewQueryOptions()
	opts.SetIfExist("unknown")

	doc := collection.NewDocument(newCollection, "")
	_, err := doc.Create(opts)
	assert.Equal(t, "[400] Invalid value for the 'ifExist' option: 'unknown'", fmt.Sprint(err))
}

func TestCreateDocument(t *testing.T) {
	done := make(chan bool)
	id := "myId"

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "document", parsedQuery.Controller)
			assert.Equal(t, "create", parsedQuery.Action)
			assert.Equal(t, "index", parsedQuery.Index)
			assert.Equal(t, "collection", parsedQuery.Collection)
			assert.Equal(t, id, parsedQuery.Id)

			body := make(map[string]interface{}, 0)
			body["title"] = "yolo"

			assert.Equal(t, body, parsedQuery.Body)

			res := collection.Document{Id: id, Content: []byte(`{"title":"yolo"}`)}
			r, _ := json.Marshal(res)
			done <- true
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res := collection.NewCollection(k, "collection", "index")
	doc := collection.NewDocument(res, id)
	content := make(collection.DocumentContent)
	content["title"] = "yolo"
	doc.SetContent(content, false)
	go doc.Create(nil)
	<-done
	assert.Equal(t, id, doc.Id)
}

func TestCreateDocumentReplace(t *testing.T) {
	done := make(chan bool)
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "createOrReplace", parsedQuery.Action)

			res := collection.Document{Id: "id", Content: []byte(`{"title":"yolo"}`)}
			r, _ := json.Marshal(res)
			done <- true
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	newCollection := collection.NewCollection(k, "collection", "index")
	opts := types.NewQueryOptions()
	opts.SetIfExist("replace")
	doc := collection.NewDocument(newCollection, "")
	content := make(collection.DocumentContent)
	content["title"] = "yolo"
	doc.SetContent(content, false)
	go doc.Create(opts)
	<-done
}

func TestCreateDocumentCreate(t *testing.T) {
	done := make(chan bool)
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "create", parsedQuery.Action)

			res := collection.Document{Id: "id", Content: []byte(`{"Title":"yolo"}`)}
			r, _ := json.Marshal(res)
			done <- true
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	newCollection := collection.NewCollection(k, "collection", "index")
	opts := types.NewQueryOptions()
	opts.SetIfExist("error")

	doc := collection.NewDocument(newCollection, "id")
	content := make(collection.DocumentContent)
	content["title"] = "yolo"
	doc.SetContent(content, false)
	go doc.Create(opts)
	<-done
}

func ExampleCollection_CreateDocument() {
	c := &internal.MockedConnection{}
	k, _ := kuzzle.NewKuzzle(c, nil)
	id := "myId"

	newCollection := collection.NewCollection(k, "collection", "index")
	doc := collection.NewDocument(newCollection, id)
	content := make(collection.DocumentContent)
	content["title"] = "foo"
	doc.SetContent(content, false)
	doc, err := doc.Create(nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(doc.Id, doc.Content)
}
