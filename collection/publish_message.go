package collection

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// PublishMessage publishes a realtime message
// Takes an optional argument object with the following properties:
//   - volatile (object, default: null):
//     Additional information passed to notifications to other users
func (dc *Collection) PublishMessage(message interface{}, options types.QueryOptions) (bool, error) {
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

	if res.Error != nil {
		return false, res.Error
	}

	response := types.RealtimeResponse{}

	json.Unmarshal(res.Result, &response)

	return response.Published, nil
}
