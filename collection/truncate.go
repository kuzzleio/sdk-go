package collection

import (
	"github.com/kuzzleio/sdk-go/types"
	"encoding/json"
	"fmt"
)

// Truncate delete every Documents from the provided Collection.
func (dc *Collection) Truncate(options types.QueryOptions) (bool, error) {
	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Collection: dc.collection,
		Index:      dc.index,
		Controller: "collection",
		Action:     "truncate",
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
