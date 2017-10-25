package collection

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// DeleteDocument deletes the Document using its provided unique id.
func (dc *Collection) DeleteDocument(id string, options types.QueryOptions) (string, error) {
	if id == "" {
		return "", types.NewError("Collection.DeleteDocument: document id required", 400)
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Collection: dc.collection,
		Index:      dc.index,
		Controller: "document",
		Action:     "delete",
		Id:         id,
	}
	go dc.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return "", res.Error
	}

	document := &Document{collection: dc}
	json.Unmarshal(res.Result, document)

	return document.Id, nil
}

// MDeleteDocument deletes specific documents according to given IDs.
func (dc *Collection) MDeleteDocument(ids []string, options types.QueryOptions) ([]string, error) {
	if len(ids) == 0 {
		return nil, types.NewError("Collection.MDeleteDocument: please provide at least one id of document to delete", 400)
	}

	ch := make(chan *types.KuzzleResponse)

	type body struct {
		Ids []string `json:"ids"`
	}

	query := &types.KuzzleRequest{
		Collection: dc.collection,
		Index:      dc.index,
		Controller: "document",
		Action:     "mDelete",
		Body:       &body{Ids: ids},
	}
	go dc.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return nil, res.Error
	}

	result := []string{}
	json.Unmarshal(res.Result, &result)

	return result, nil
}
