package ms

import (
	"github.com/kuzzleio/sdk-go/types"
)

// Ltrim trims an existing list so that it will
// contain only the specified range of elements specified.
func (ms *Ms) Ltrim(key string, start int, stop int, options types.QueryOptions) error {
	result := make(chan *types.KuzzleResponse)

	type body struct {
		Start int `json:"start"`
		Stop  int `json:"stop"`
	}

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "ltrim",
		Id:         key,
		Body:       &body{Start: start, Stop: stop},
	}

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	return res.Error
}
