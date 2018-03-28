package index

import (
	"encoding/json"
	"fmt"

	"github.com/kuzzleio/sdk-go/types"
)

// Exists check if the index exists
func (i *Index) Exists(index string, options types.QueryOptions) (bool, error) {
	if index == "" {
		return false, types.NewError("Index.Exists: index required", 400)
	}

	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Index:      index,
		Controller: "index",
		Action:     "exists",
	}

	go i.kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return false, res.Error
	}

	var exists bool

	err := json.Unmarshal(res.Result, &exists)
	if err != nil {
		return false, types.NewError(fmt.Sprintf("Unable to parse response: %s\n%s", err.Error(), res.Result), 500)
	}
	return exists, nil

}
