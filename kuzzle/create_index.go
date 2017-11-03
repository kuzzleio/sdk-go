package kuzzle

import (
	"github.com/kuzzleio/sdk-go/types"
	"encoding/json"
	"fmt"
)

// CreateIndex create a new empty data index, with no associated mapping.
func (k Kuzzle) CreateIndex(index string, options types.QueryOptions) (bool, error) {
	if index == "" {
		return false, types.NewError("Kuzzle.createIndex: index required", 400)
	}

	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Index:      index,
		Controller: "index",
		Action:     "create",
	}
	go k.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return false, res.Error
	}

	ack := &struct {
		Acknowledged bool `json:"acknowledged"`
	}{}
	err := json.Unmarshal(res.Result, ack)
	if err != nil {
		return false, types.NewError(fmt.Sprintf("Unable to parse response: %s\n%s", err.Error(), res.Result), 500)
	}
	return ack.Acknowledged, nil
}
