package index

import "github.com/kuzzleio/sdk-go/types"

func (i *Index) RefreshInternal() error {

	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "index",
		Action:     "refreshInternal",
	}
	go i.k.Query(query, nil, result)

	res := <-result

	if res.Error != nil {
		return res.Error
	}

	return nil
}
