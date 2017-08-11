package collection

import (
	"encoding/json"
	"errors"
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
	Id         string           `json:"_id"`
	Meta       types.KuzzleMeta `json:"_meta"`
	Content    json.RawMessage  `json:"_source"`
	Version    int              `json:"_version"`
	Collection string           `json:"collection"`
	collection Collection       `json:"-"`
}

type DocumentContent map[string]interface{}

func (documentContent DocumentContent) ToString() string {
	s, _ := json.Marshal(documentContent)

	return string(s)
}

func (d Document) SourceToMap() map[string]interface{} {
	type SourceMap map[string]interface{}
	sourceMap := SourceMap{}

	json.Unmarshal(d.Content, &sourceMap)

	return sourceMap
}

/*
  Saves the document into Kuzzle.

  If this is a new document, will create it in Kuzzle and the id property will be made available.
  Otherwise, will replace the latest version of the document in Kuzzle by the current content of this object.
*/
func (d Document) Save(options types.QueryOptions) (Document, error) {
	if d.Id == "" {
		return Document{}, errors.New("Document.Save: missing document id")
	}

	ch := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Index:      d.collection.index,
		Collection: d.collection.collection,
		Controller: "document",
		Action:     "createOrReplace",
		Id:         d.Id,
		Body:       d.Content,
	}

	go d.collection.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Message != "" {
		return Document{}, errors.New(res.Error.Message)
	}

	return d, nil
}

/*
  Replaces the document with the latest version stored in Kuzzle.
*/
func (d Document) Refresh(options types.QueryOptions) (Document, error) {
	if d.Id == "" {
		return Document{}, errors.New("Document.Refresh: missing document id")
	}

	ch := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Index:      d.collection.index,
		Collection: d.collection.collection,
		Controller: "document",
		Action:     "get",
		Id:         d.Id,
	}

	go d.collection.Kuzzle.Query(query, options, ch)

	res := <-ch
	if res.Error.Message != "" {
		return Document{}, errors.New(res.Error.Message)
	}

	json.Unmarshal(res.Result, &d)

	return d, nil
}

/*
  Sets the document id.
*/
func (d Document) SetDocumentId(id string) Document {
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
func (d Document) SetContent(content DocumentContent, replace bool) Document {
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
func (d *Document) SetHeaders(content map[string]interface{}, replace bool) {
	d.collection.Kuzzle.SetHeaders(content, replace)
}

/*
  Sends the content of the document as a realtime message.
*/
func (d Document) Publish(options types.QueryOptions) (bool, error) {
	ch := make(chan types.KuzzleResponse)

	type message struct {
		Id      string           `json:"_id,omitempty"`
		Version int              `json:"_version,omitempty"`
		Body    json.RawMessage  `json:"body"`
		Meta    types.KuzzleMeta `json:"meta"`
	}

	query := types.KuzzleRequest{
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

	if res.Error.Message != "" {
		return false, errors.New(res.Error.Message)
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
		return false, errors.New("Document.Exists: missing document id")
	}

	ch := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Index:      d.collection.index,
		Collection: d.collection.collection,
		Controller: "document",
		Action:     "exists",
		Id:         d.Id,
	}

	go d.collection.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Message != "" {
		return false, errors.New(res.Error.Message)
	}

	exists, _ := strconv.ParseBool(string(res.Result))

	return exists, nil
}

/*
  Deletes the document in Kuzzle.
*/
func (d Document) Delete(options types.QueryOptions) (string, error) {
	if d.Id == "" {
		return "", errors.New("Document.Delete: missing document id")
	}

	ch := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Index:      d.collection.index,
		Collection: d.collection.collection,
		Controller: "document",
		Action:     "delete",
		Id:         d.Id,
	}

	go d.collection.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Message != "" {
		return "", errors.New(res.Error.Message)
	}

	document := Document{collection: d.collection}
	json.Unmarshal(res.Result, &document)

	return document.Id, nil
}
