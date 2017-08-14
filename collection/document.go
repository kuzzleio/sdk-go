package collection

import (
	"encoding/json"
	"errors"
	"fmt"
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

type CollectionDocument struct {
	Collection Collection `json:"-"`
	Document   types.Document
}

type DocumentContent map[string]interface{}

func (documentContent DocumentContent) ToString() string {
	s, _ := json.Marshal(documentContent)

	return string(s)
}

/*
  Helper function to initialize a document into CollectionDocument using fetch query.
 */
func (cd CollectionDocument) Fetch(id string) (CollectionDocument, error) {
	if id == "" {
		return cd, errors.New("CollectionDocument.Fetch: missing document id")
	}

	doc, err := cd.Collection.FetchDocument(id, nil)

	if err != nil {
		return cd, errors.New("CollectionDocument.Fetch: an error occurred: " + fmt.Sprint(err))
	}

	cd.Document = doc

	return cd, nil
}

/*
  Listens to events concerning this document. Has no effect if the document does not have an ID
  (i.e. if the document has not yet been created as a persisted document).
 */
func (cd CollectionDocument) Subscribe(options types.RoomOptions, ch chan<- types.KuzzleNotification) chan types.SubscribeResponse {
	if cd.Document.Id == "" {
		errorResponse := make(chan types.SubscribeResponse, 1)
		errorResponse <- types.SubscribeResponse{Error: errors.New("CollectionDocument.Subscribe: cannot subscribe to a document if no ID has been provided")}

		return errorResponse
	}

	filters := map[string]map[string][]string{
		"ids": {
			"values": []string{cd.Document.Id},
		},
	}

	return cd.Collection.Subscribe(filters, options, ch)
}

/*
  Saves the document into Kuzzle.

  If this is a new document, will create it in Kuzzle and the id property will be made available.
  Otherwise, will replace the latest version of the document in Kuzzle by the current content of this object.
*/
func (cd CollectionDocument) Save(options types.QueryOptions) (CollectionDocument, error) {
	if cd.Document.Id == "" {
		return cd, errors.New("CollectionDocument.Save: missing document id")
	}

	ch := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Index:      cd.Collection.index,
		Collection: cd.Collection.collection,
		Controller: "document",
		Action:     "createOrReplace",
		Id:         cd.Document.Id,
		Body:       cd.Document.Source,
	}

	go cd.Collection.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Message != "" {
		return cd, errors.New(res.Error.Message)
	}

	return cd, nil
}

/*
  Replaces the document with the latest version stored in Kuzzle.
*/
func (cd CollectionDocument) Refresh(options types.QueryOptions) (CollectionDocument, error) {
	if cd.Document.Id == "" {
		return cd, errors.New("CollectionDocument.Refresh: missing document id")
	}

	ch := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Index:      cd.Collection.index,
		Collection: cd.Collection.collection,
		Controller: "document",
		Action:     "get",
		Id:         cd.Document.Id,
	}

	go cd.Collection.Kuzzle.Query(query, options, ch)

	res := <-ch
	if res.Error.Message != "" {
		return cd, errors.New(res.Error.Message)
	}

	document := types.Document{Id: cd.Document.Id}
	json.Unmarshal(res.Result, &document)

	cd.Document = document

	return cd, nil
}

/*
  Sets the document id.
*/
func (cd CollectionDocument) SetDocumentId(id string) CollectionDocument {
	if id != "" {
		cd.Document.Id = id
	}

	return cd
}

/*
  Replaces the current document content with provided data.
  Changes made by this function won’t be applied until the save method is called.
  If replace is set to true, the entire content will be replaced, otherwise, only existing and new fields will be impacted.
*/
func (cd CollectionDocument) SetContent(content DocumentContent, replace bool) CollectionDocument {
	if replace {
		cd.Document.Source, _ = json.Marshal(content)
	} else {
		source := DocumentContent{}
		json.Unmarshal(cd.Document.Source, &source)

		for attr, value := range content {
			source[attr] = value
		}

		cd.Document.Source, _ = json.Marshal(source)
	}

	return cd
}

/*
  Helper function allowing to set headers while chaining calls.

  If the replace argument is set to true, replaces the current headers with the provided ones.
  Otherwise, appends the content to the current headers, only replacing already existing values.
*/
func (cd *CollectionDocument) SetHeaders(content map[string]interface{}, replace bool) {
	cd.Collection.Kuzzle.SetHeaders(content, replace)
}

/*
  Sends the content of the document as a realtime message.
*/
func (cd CollectionDocument) Publish(options types.QueryOptions) (bool, error) {
	ch := make(chan types.KuzzleResponse)

	type message struct {
		Id      string           `json:"_id,omitempty"`
		Version int              `json:"_version,omitempty"`
		Body    json.RawMessage  `json:"body"`
		Meta    types.KuzzleMeta `json:"meta"`
	}

	query := types.KuzzleRequest{
		Index:      cd.Collection.index,
		Collection: cd.Collection.collection,
		Controller: "realtime",
		Action:     "publish",
		Body: message{
			Id:      cd.Document.Id,
			Version: cd.Document.Version,
			Body:    cd.Document.Source,
			Meta:    cd.Document.Meta,
		},
	}

	go cd.Collection.Kuzzle.Query(query, options, ch)

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
func (cd CollectionDocument) Exists(options types.QueryOptions) (bool, error) {
	if cd.Document.Id == "" {
		return false, errors.New("CollectionDocument.Exists: missing document id")
	}

	ch := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Index:      cd.Collection.index,
		Collection: cd.Collection.collection,
		Controller: "document",
		Action:     "exists",
		Id:         cd.Document.Id,
	}

	go cd.Collection.Kuzzle.Query(query, options, ch)

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
func (cd CollectionDocument) Delete(options types.QueryOptions) (string, error) {
	if cd.Document.Id == "" {
		return "", errors.New("CollectionDocument.Delete: missing document id")
	}

	ch := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Index:      cd.Collection.index,
		Collection: cd.Collection.collection,
		Controller: "document",
		Action:     "delete",
		Id:         cd.Document.Id,
	}

	go cd.Collection.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Message != "" {
		return "", errors.New(res.Error.Message)
	}

	document := types.Document{}
	json.Unmarshal(res.Result, &document)

	return document.Id, nil
}
