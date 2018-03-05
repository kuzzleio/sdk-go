package collection

import (
	"github.com/kuzzleio/sdk-go/types"
)

// ListCollections retrieves the list of known data collections contained in a specified index.
func (dc *Collection) List(index string, options *ListOptions) (string, error) {
	if index == "" {
		return "", types.NewError("Collection.List: index required", 400)
	}

	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "collection",
		Action:     "list",
		Index:      index,
		Body:       options.Type,
	}

	queryOpts := types.NewQueryOptions()
	queryOpts.SetFrom(options.From)
	queryOpts.SetSize(options.Size)

	go dc.Kuzzle.Query(query, queryOpts, result)

	res := <-result

	if res.Error != nil {
		return "", res.Error
	}

	var collectionList string
	collectionList = string(res.Result)

	return collectionList, nil
}
