package collection

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Create creates a new empty data collection, with no associated mapping.
func (dc *Collection) Create(options types.QueryOptions) (*types.AckResponse, error) {
	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Collection: dc.collection,
		Index:      dc.index,
		Controller: "collection",
		Action:     "create",
	}
	go dc.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return nil, res.Error
	}

	ack := &types.AckResponse{}
	json.Unmarshal(res.Result, ack)

	return ack, nil
}
