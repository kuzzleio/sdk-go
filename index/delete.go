package index

import "github.com/kuzzleio/sdk-go/types"

// Delete the given index
func (i *Index) Delete(index string) error {
	if index == "" {
		return types.NewError("Index.Delete: index required", 400)
	}

	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Index:      index,
		Controller: "index",
		Action:     "delete",
	}
	go i.kuzzle.Query(query, nil, result)

	res := <-result

	if res.Error != nil {
		return res.Error
	}

	return nil
}
