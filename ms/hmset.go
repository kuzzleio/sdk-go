package ms

import (
	"github.com/kuzzleio/sdk-go/types"
)

// Hmset sets multiple fields at once in a hash.
func (ms Ms) Hmset(key string, entries []*types.MsHashField, options types.QueryOptions) error {
	if len(entries) == 0 {
		return types.NewError("Ms.Hmset: at least one entry field to set is required", 400)
	}

	result := make(chan *types.KuzzleResponse)

	type body struct {
		Entries []*types.MsHashField `json:"entries"`
	}

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "hmset",
		Id:         key,
		Body:       &body{Entries: entries},
	}

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	return res.Error
}
