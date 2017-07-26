package kuzzle

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

/*
 * Delete credentials of the specified strategy for the current user.
 */
func (k Kuzzle) DeleteMyCredentials(strategy string, credentials interface{}, options types.QueryOptions) (types.AckResponse, error) {
	type body struct {
		Strategy string      `json:"strategy"`
		Body     interface{} `json:"body"`
	}
	result := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Controller: "auth",
		Action:     "deleteMyCredentials",
		Body:       &body{strategy, credentials},
	}

	go k.Query(query, options, result)

	res := <-result

	if res.Error.Message != "" {
		return types.AckResponse{}, errors.New(res.Error.Message)
	}

	ack := types.AckResponse{}
	json.Unmarshal(res.Result, &ack)

	return ack, nil
}
