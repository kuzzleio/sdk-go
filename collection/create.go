package collection

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

/*
  Create a new empty data collection, with no associated mapping.
*/
func (dc Collection) Create(options types.QueryOptions) (types.AckResponse, error) {
	ch := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Collection: dc.collection,
		Index:      dc.index,
		Controller: "collection",
		Action:     "create",
	}
	go dc.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Message != "" {
		return types.AckResponse{}, errors.New(res.Error.Message)
	}

	ack := types.AckResponse{}
	json.Unmarshal(res.Result, &ack)

	return ack, nil
}
