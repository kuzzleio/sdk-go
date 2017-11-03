package collection

import (
	"github.com/kuzzleio/sdk-go/types"
	"encoding/json"
	"fmt"
)

// Create creates a new empty data collection, with no associated mapping.
func (dc *Collection) Create(options types.QueryOptions) (bool, error) {
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
		return false, res.Error
	}

	ack := struct {
		Acknowledged bool `json:"acknowledged"`
	}{}
	json.Unmarshal(res.Result, ack)

	return ack.Acknowledged, nil
}
