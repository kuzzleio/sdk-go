package index

import (
	"github.com/kuzzleio/sdk-go/types"
)

// Create a new empty data index, with no associated mapping.
func (i *Index) Create(index string) error {
	if index == "" {
		return types.NewError("Index.Create: index required", 400)
	}

	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Index:      index,
		Controller: "index",
		Action:     "create",
	}
	go i.k.Query(query, nil, result)

	res := <-result

	if res.Error != nil {
		return res.Error
	}

	return nil
}
