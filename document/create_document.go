package collection

import (
	"encoding/json"
	"fmt"
	"github.com/kuzzleio/sdk-go/types"
)

// Create a new document in Kuzzle.
// Takes an optional argument object with the following properties:
//   - volatile (object, default: null):
//       Additional information passed to notifications to other users
//   - ifExist (string, allowed values: "error" (default), "replace"):
//       If the same document already exists:
//         - resolves with an error if set to "error".
//         - replaces the existing document if set to "replace"
func (dc *Collection) CreateDocument(id string, document *Document, options types.QueryOptions) (*Document, error) {
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
		Collection: dc.collection,
		Index:      dc.index,
		Controller: "document",
		Action:     action,
		Body:       document.Content,
		Id:         id,
	}
	go dc.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return nil, res.Error
	}

	documentResponse := &Document{collection: dc}
	json.Unmarshal(res.Result, documentResponse)

	return documentResponse, nil
}
