package ms

import (
	"github.com/kuzzleio/sdk-go/types"
)

// Flushdb delete all keys from the database
func (ms *Ms) Flushdb(options types.QueryOptions) error {
	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "flushdb",
	}
	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	return res.Error
}
