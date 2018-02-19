package collection

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/kuzzleio/sdk-go/types"
)

type IDocument interface {
	Save()
	Refresh()
	SetContent()
	Publish()
	Exists()
	Delete()
}

type Document struct {
	Id         string          `json:"_id"`
	Index      string          `json:"_index"`
	Meta       *types.Meta     `json:"_meta"`
	Shards     *types.Shards   `json:"_shards"`
	Content    json.RawMessage `json:"_source"`
	Version    int             `json:"_version"`
	Result     string          `json:"result"`
	Created    bool            `json:"created"`
	Collection string          `json:"_collection"`
	collection *Collection     `json:"-"`
}

type DocumentContent map[string]interface{}

/*
 * Instanciate a new document
 */
func NewDocument(col *Collection, id string) *Document {
	return &Document{
		collection: col,
		Id:         id,
		Collection: col.collection,
		Index:      col.index,
	}
}

func (documentContent *DocumentContent) ToString() string {
	s, _ := json.Marshal(documentContent)

	return string(s)
}

func (d *Document) SourceToMap() DocumentContent {
	sourceMap := DocumentContent{}

	json.Unmarshal(d.Content, &sourceMap)

	return sourceMap
}

// Subscribe listens to events concerning this document. Has no effect if the document does not have an ID
// (i.e. if the document has not yet been created as a persisted document).
func (d *Document) Subscribe(options types.RoomOptions, ch chan<- types.KuzzleNotification) (*Room, error) {
	if d.Id == "" {
		return nil, types.NewError("Document.Subscribe: cannot subscribe to a document if no ID has been provided", 400)
	}

	filters := map[string]map[string][]string{
		"ids": {
			"values": []string{d.Id},
		},
	}

	return d.collection.Subscribe(filters, options, ch), nil
}

// Create a new document in Kuzzle.
// Takes an optional argument object with the following properties:
//   - volatile (object, default: null):
//       Additional information passed to notifications to other users
//   - ifExist (string, allowed values: "error" (default), "replace"):
//       If the same document already exists:
//         - resolves with an error if set to "error".
//         - replaces the existing document if set to "replace"
func (d *Document) Create(options types.QueryOptions) (*Document, error) {
	ch := make(chan *types.KuzzleResponse)

	action := "create"

	if options != nil {
		if options.IfExist() == "replace" {
			action = "createOrReplace"
		} else if options.IfExist() != "error" {
			return nil, types.NewError(fmt.Sprintf("Invalid value for the 'ifExist' option: '%s'", options.IfExist()), 400)
		}
	}

	query := &types.KuzzleRequest{
		Collection: d.collection.collection,
		Index:      d.collection.index,
		Controller: "document",
		Action:     action,
		Body:       d.Content,
		Id:         d.Id,
	}
	go d.collection.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return nil, res.Error
	}

	documentResponse := &Document{collection: d.collection}
	json.Unmarshal(res.Result, documentResponse)

	return documentResponse, nil
}

// Replace a document in Kuzzle.
func (d *Document) Replace(options types.QueryOptions) (*Document, error) {
	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Collection: d.collection.collection,
		Index:      d.collection.index,
		Controller: "document",
		Action:     "replace",
		Body:       d.Content,
		Id:         d.Id,
	}
	go d.collection.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return nil, res.Error
	}

	documentResponse := &Document{collection: d.collection}
	json.Unmarshal(res.Result, documentResponse)

	return documentResponse, nil
}

// Replace a document in Kuzzle.
func (d *Document) Update(options types.QueryOptions) (*Document, error) {
	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Collection: d.collection.collection,
		Index:      d.collection.index,
		Controller: "document",
		Action:     "update",
		Body:       d.Content,
		Id:         d.Id,
	}
	go d.collection.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return nil, res.Error
	}

	documentResponse := &Document{collection: d.collection}
	json.Unmarshal(res.Result, documentResponse)

	return documentResponse, nil
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
			if source[attr] == nil {
				source[attr] = value
			}
		}

		d.Content, _ = json.Marshal(source)
	}

	return d
}

/*
  Sends the content of the document as a realtime message.
*/
func (d *Document) Publish(options types.QueryOptions) (bool, error) {
	ch := make(chan *types.KuzzleResponse)

	type message struct {
		Id      string          `json:"_id,omitempty"`
		Version int             `json:"_version,omitempty"`
		Body    json.RawMessage `json:"body"`
		Meta    *types.Meta     `json:"meta"`
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
func (d *Document) Exists(options types.QueryOptions) (bool, error) {
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
func (d *Document) Delete(options types.QueryOptions) (string, error) {
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
