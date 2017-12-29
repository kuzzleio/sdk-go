package collection

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// ReplaceDocument replaces a document in Kuzzle.
func (dc *Collection) ReplaceDocument(id string, document *Document, options types.QueryOptions) (*Document, error) {
	if id == "" {
		return nil, types.NewError("Collection.ReplaceDocument: document id required", 400)
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Collection: dc.collection,
		Index:      dc.index,
		Controller: "document",
		Action:     "replace",
		Body:       document.Content,
		Id:         id,
	}
	go dc.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return nil, res.Error
	}

	d := &Document{collection: dc}
	json.Unmarshal(res.Result, d)

	return d, nil
}
