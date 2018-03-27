package index

import (
	"encoding/json"
	"fmt"

	"github.com/kuzzleio/sdk-go/types"
)

// List list all index
func (i *Index) List(options types.QueryOptions) ([]string, error) {
	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "index",
		Action:     "list",
	}

	go i.kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return nil, res.Error
	}

	var collectionList struct {
		Total int
		Hits  []string
	}

	err := json.Unmarshal(res.Result, &collectionList)

	if err != nil {
		return nil, types.NewError(fmt.Sprintf("Unable to parse response: %s\n%s", err.Error(), res.Result), 500)
	}

	return collectionList.Hits, nil
}
