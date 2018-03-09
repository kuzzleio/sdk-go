package index

import (
	"github.com/kuzzleio/sdk-go/types"
)

func (i *Index) Refresh(index string) error {
	if index == "" {
		return types.NewError("Index.Refresh: index required", 400)
	}

	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "index",
		Action:     "refresh",
		Index:      index,
	}
	go i.kuzzle.Query(query, nil, result)

	res := <-result

	return res.Error
}
