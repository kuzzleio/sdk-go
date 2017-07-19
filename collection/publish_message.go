package collection

import (
	"github.com/kuzzleio/sdk-go/types"
	"errors"
	"encoding/json"
)

/*
  Publish a realtime message

  Takes an optional argument object with the following properties:
    - volatile (object, default: null):
      Additional information passed to notifications to other users
*/
func (dc Collection) PublishMessage(document interface{}, options *types.Options) (*types.RealtimeResponse, error) {
	ch := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Collection: dc.collection,
		Index:      dc.index,
		Controller: "realtime",
		Action:     "publish",
		Body:       document,
	}
	go dc.kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Message != "" {
		return nil, errors.New(res.Error.Message)
	}

	response := &types.RealtimeResponse{}
	json.Unmarshal(res.Result, response)

	return response, nil
}
