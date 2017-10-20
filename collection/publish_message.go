package collection

import (
	"github.com/kuzzleio/sdk-go/types"
)

// PublishMessage publishes a realtime message
// Takes an optional argument object with the following properties:
//   - volatile (object, default: null):
//     Additional information passed to notifications to other users
func (dc *Collection) PublishMessage(message map[string]interface{}, options types.QueryOptions) (*Collection, error) {
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

	return dc, res.Error
}
