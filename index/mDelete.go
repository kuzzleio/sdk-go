package index

import (
	"encoding/json"
	"fmt"

	"github.com/kuzzleio/sdk-go/types"
)

// Delete delete the index
func (i *Index) MDelete(indexes []string) ([]string, error) {
	if len(indexes) == 0 {
		return nil, types.NewError("Index.MDelete: at least one index required", 400)
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
		return nil, res.Error
	}

	var deletedIndexes []string

	err := json.Unmarshal(res.Result, &deletedIndexes)
	if err != nil {
		return nil, types.NewError(fmt.Sprintf("Unable to parse response: %s\n%s", err.Error(), res.Result), 500)
	}

	return deletedIndexes, nil
}
