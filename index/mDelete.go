package index

import "github.com/kuzzleio/sdk-go/types"

// Delete delete the index
func (Index) MDelete(indexes []string) error {
	if indexes.len() == 0 {
		return types.NewError("Index.MDelete: at least one index required", 400)
	}

	result := make(chan *types.KuzzleResponse)

	for _, index := range indexes {
		query := &types.KuzzleRequest{
			Index:      index,
			Controller: "index",
			Action:     "delete",
		}
		go k.Query(query, nil, result)

		res := <-result

		if res.Error != nil {
			return res.Error
		}
	}

	return nil
}
