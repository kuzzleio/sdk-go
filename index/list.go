package index

import "github.com/kuzzleio/sdk-go/types"

func (i *Index) List() (string, error) {
	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "index",
		Action:     "list",
	}

	go i.k.Query(query, nil, result)

	res := <-result

	if res.Error != nil {
		return "", res.Error
	}

	var collectionList string
	collectionList = string(res.Result)

	return collectionList, nil
}
