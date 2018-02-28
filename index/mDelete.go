package index

import "github.com/kuzzleio/sdk-go/types"

// Delete delete the index
func (i *Index) MDelete(indexes []string) error {
	if len(indexes) == 0 {
		return types.NewError("Index.MDelete: at least one index required", 400)
	}

	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "index",
		Action:     "mDelete",
		Body:       indexes,
	}
	go i.k.Query(query, nil, result)

	res := <-result

	if res.Error != nil {
		return res.Error
	}

	return nil
}
