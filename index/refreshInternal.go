package index

import "github.com/kuzzleio/sdk-go/types"

func (i *Index) RefreshInternal(options types.QueryOptions) error {

	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "index",
		Action:     "refreshInternal",
	}
	go i.kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return res.Error
	}

	return nil
}
