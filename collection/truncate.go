package collection

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

/*
  Delete every Documents from the provided Collection.
*/
func (dc Collection) Truncate(options types.QueryOptions) (types.AckResponse, error) {
	ch := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Collection: dc.collection,
		Index:      dc.index,
		Controller: "collection",
		Action:     "truncate",
	}
	go dc.kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Message != "" {
		return types.AckResponse{}, errors.New(res.Error.Message)
	}

	ack := types.AckResponse{}
	json.Unmarshal(res.Result, &ack)

	return ack, nil
}
