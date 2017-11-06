package collection

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
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

	ack := &struct {
		Acknowledged bool `json:"acknowledged"`
	}{}
	err := json.Unmarshal(res.Result, ack)
	if err != nil {
		return false, types.NewError(fmt.Sprintf("Unable to parse response: %s\n%s", err.Error(), res.Result), 500)
	}

	return ack.Acknowledged, nil
}
