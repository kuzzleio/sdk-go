package kuzzle

import (
	"encoding/json"
	"fmt"
	"github.com/kuzzleio/sdk-go/types"
)

// DeleteMyCredentials delete credentials of the specified strategy for the current user.
func (k Kuzzle) DeleteMyCredentials(strategy string, options types.QueryOptions) (bool, error) {
	if strategy == "" {
		return false, types.NewError("Kuzzle.DeleteMyCredentials: strategy is required", 400)
	}

	type body struct {
		Strategy string `json:"strategy"`
	}
	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "auth",
		Action:     "deleteMyCredentials",
		Strategy:   strategy,
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
