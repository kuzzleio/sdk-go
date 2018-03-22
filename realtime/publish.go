package realtime

import "github.com/kuzzleio/sdk-go/types"

// Publish sends a real-time message to Kuzzle
func (r *Realtime) Publish(index string, collection string, body string) error {
	if (index == "" || collection == "") || body == "" {
		return types.NewError("Realtime.Publish: index, collection and body required", 400)
	}

	query := &types.KuzzleRequest{
		Controller: "realtime",
		Action:     "publish",
		Index:      index,
		Collection: collection,
		Body:       body,
	}

	result := make(chan *types.KuzzleResponse)

	go r.k.Query(query, nil, result)

	res := <-result

	if res.Error != nil {
		return res.Error
	}

	return nil
}
