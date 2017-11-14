package ms

import (
	"github.com/kuzzleio/sdk-go/types"
)

// SetEx sets a key with the provided value, and an expiration delay expressed in seconds.
// If the key does not exist, it is created beforehand.
func (ms Ms) Setex(key string, value interface{}, ttl int, options types.QueryOptions) error {
	result := make(chan *types.KuzzleResponse)

	type body struct {
		Value interface{} `json:"value"`
		Ttk   int         `json:"seconds"`
	}

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "setex",
		Id:         key,
		Body:       &body{Value: value, Ttk: ttl},
	}

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	return res.Error
}
