package index

import (
	"encoding/json"
	"fmt"

	"github.com/kuzzleio/sdk-go/types"
)

func (i *Index) GetAutoRefresh(index string) (bool, error) {
	if index == "" {
		return false, types.NewError("Index.GetAutoRefresh: index required", 400)
	}

	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Index:      index,
		Controller: "index",
		Action:     "getAutoRefresh",
	}

	go i.kuzzle.Query(query, nil, result)

	res := <-result

	if res.Error != nil {
		return false, res.Error
	}

	var autoR bool

	err := json.Unmarshal(res.Result, &autoR)
	if err != nil {
		return false, types.NewError(fmt.Sprintf("Unable to parse response: %s\n%s", err.Error(), res.Result), 500)
	}
	return autoR, nil

}
