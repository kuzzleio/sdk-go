package collection

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

// PublishMessage publishes a realtime message
// Takes an optional argument object with the following properties:
//   - volatile (object, default: null):
//     Additional information passed to notifications to other users
func (dc Collection) PublishMessage(message map[string]interface{}, options types.QueryOptions) (*types.RealtimeResponse, error) {
	ch := make(chan *types.KuzzleResponse)


	query := &types.KuzzleRequest{
		Collection: dc.collection,
		Index:      dc.index,
		Controller: "realtime",
		Action:     "publish",
		Body:       message,
	}
	go dc.Kuzzle.Query(query, options, ch)

	res := <-ch

	response := &types.RealtimeResponse{}

	if res.Error != nil {
		return response, errors.New(res.Error.Message)
	}

	json.Unmarshal(res.Result, response)

	return response, nil
}
