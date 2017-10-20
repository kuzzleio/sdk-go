package collection

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
	"strconv"
)

type IDocument interface {
	Save()
	Refresh()
	SetContent()
	Publish()
	Exists()
	SetHeaders()
	Delete()
}

type Document struct {
	Id         string            `json:"_id"`
	Index      string 					 `json:"_index"`
	Meta       *types.KuzzleMeta `json:"_meta"`
	Shards     map[string]int    `json:"_shards"`
	Content    json.RawMessage   `json:"_source"`
	Version    int               `json:"_version"`
	Result     string            `json:"result"`
	Created    bool 						 `json:"created"`
	Collection string            `json:"_collection"`
	collection *Collection       `json:"-"`
}

type DocumentContent map[string]interface{}

func (documentContent DocumentContent) ToString() string {
	s, _ := json.Marshal(documentContent)

	return string(s)
}

func (d Document) SourceToMap() DocumentContent {
	sourceMap := DocumentContent{}

	json.Unmarshal(d.Content, &sourceMap)

	return sourceMap
}

// Helper function to initialize a document into Document using fetch query.
func (d *Document) Fetch(id string) (*Document, error) {
	if id == "" {
		return d, types.NewError("Document.Fetch: missing document id", 400)
	}

	doc, err := d.collection.FetchDocument(id, nil)

	if err != nil {
		err.(*types.KuzzleError).Message = "Document.Fetch: an error occurred: " + err.(*types.KuzzleError).Message
		return d, err
	}

	d.Id = id
	d.Index = doc.Index
	d.Meta = doc.Meta
	d.Shards = doc.Shards
	d.Content = doc.Content
	d.Version = doc.Version
	d.Result = doc.Result
	d.Created = doc.Created
	d.Collection = doc.Collection
	d.collection = doc.collection

	return d, nil
}

// Subscribe listens to events concerning this document. Has no effect if the document does not have an ID
// (i.e. if the document has not yet been created as a persisted document).
func (d Document) Subscribe(options types.RoomOptions, ch chan<- *types.KuzzleNotification) chan *types.SubscribeResponse {
	if d.Id == "" {
		errorResponse := make(chan *types.SubscribeResponse, 1)
		errorResponse <- &types.SubscribeResponse{Error: types.NewError("Document.Subscribe: cannot subscribe to a document if no ID has been provided", 400)}

		return errorResponse
	}

	filters := map[string]map[string][]string{
		"ids": {
			"values": []string{d.Id},
		},
	}

	return d.collection.Subscribe(filters, options, ch)
}

/*
  Saves the document into Kuzzle.

  If this is a new document, will create it in Kuzzle and the id property will be made available.
  Otherwise, will replace the latest version of the document in Kuzzle by the current content of this object.
*/
func (d *Document) Save(options types.QueryOptions) (*Document, error) {
	if d.Id == "" {
		return d, types.NewError("Document.Save: missing document id", 400)
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Index:      d.collection.index,
		Collection: d.collection.collection,
		Controller: "document",
		Action:     "createOrReplace",
		Id:         d.Id,
		Body:       d.Content,
	}

	go d.collection.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return d, res.Error
	}

	return d, nil
}

/*
  Replaces the document with the latest version stored in Kuzzle.
*/
func (d *Document) Refresh(options types.QueryOptions) (*Document, error) {
	if d.Id == "" {
		return d, types.NewError("Document.Refresh: missing document id", 400)
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Index:      d.collection.index,
		Collection: d.collection.collection,
		Controller: "document",
		Action:     "get",
		Id:         d.Id,
	}

	go d.collection.Kuzzle.Query(query, options, ch)

	res := <-ch
	if res.Error != nil {
		return d, res.Error
	}

	json.Unmarshal(res.Result, d)

	return d, nil
}

/*
  Sets the document id.
*/
func (d *Document) SetDocumentId(id string) *Document {
	if id != "" {
		d.Id = id
	}

	return d
}

/*
  Replaces the current document content with provided data.
  Changes made by this function won’t be applied until the save method is called.
  If replace is set to true, the entire content will be replaced, otherwise, only existing and new fields will be impacted.
*/
func (d *Document) SetContent(content DocumentContent, replace bool) *Document {
	if replace {
		d.Content, _ = json.Marshal(content)
	} else {
		source := DocumentContent{}
		json.Unmarshal(d.Content, &source)

		for attr, value := range content {
			source[attr] = value
		}

		d.Content, _ = json.Marshal(source)
	}

	return d
}

/*
  Helper function allowing to set headers while chaining calls.

  If the replace argument is set to true, replaces the current headers with the provided ones.
  Otherwise, appends the content to the current headers, only replacing already existing values.
*/
func (d Document) SetHeaders(content map[string]interface{}, replace bool) {
	d.collection.Kuzzle.SetHeaders(content, replace)
}

/*
  Sends the content of the document as a realtime message.
*/
func (d Document) Publish(options types.QueryOptions) (bool, error) {
	ch := make(chan *types.KuzzleResponse)

	type message struct {
		Id      string            `json:"_id,omitempty"`
		Version int               `json:"_version,omitempty"`
		Body    json.RawMessage   `json:"body"`
		Meta    *types.KuzzleMeta `json:"meta"`
	}

	query := &types.KuzzleRequest{
		Index:      d.collection.index,
		Collection: d.collection.collection,
		Controller: "realtime",
		Action:     "publish",
		Body: message{
			Id:      d.Id,
			Version: d.Version,
			Body:    d.Content,
			Meta:    d.Meta,
		},
	}

	go d.collection.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return false, res.Error
	}

	response := types.RealtimeResponse{}

	json.Unmarshal(res.Result, &response)

	return response.Published, nil
}

/*
  Checks if the document exists in Kuzzle.
*/
func (d Document) Exists(options types.QueryOptions) (bool, error) {
	if d.Id == "" {
		return false, types.NewError("Document.Exists: missing document id", 400)
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Index:      d.collection.index,
		Collection: d.collection.collection,
		Controller: "document",
		Action:     "exists",
		Id:         d.Id,
	}

	go d.collection.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return false, res.Error
	}

	exists, _ := strconv.ParseBool(string(res.Result))

	return exists, nil
}

/*
  Deletes the document in Kuzzle.
*/
func (d Document) Delete(options types.QueryOptions) (string, error) {
	if d.Id == "" {
		return "", types.NewError("Document.Delete: missing document id", 400)
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Index:      d.collection.index,
		Collection: d.collection.collection,
		Controller: "document",
		Action:     "delete",
		Id:         d.Id,
	}

	go d.collection.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return "", res.Error
	}

	document := Document{collection: d.collection}
	json.Unmarshal(res.Result, &document)

	return document.Id, nil
}
