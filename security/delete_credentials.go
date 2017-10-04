package security

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

// DeleteCredentials deletes credentials of the specified strategy for the given user.
func (s Security) DeleteCredentials(strategy string, kuid string, options types.QueryOptions) (*types.AckResponse, error) {
	if strategy == "" {
		return &types.AckResponse{}, errors.New("Security.DeleteCredentials: strategy is required")
	}

	if kuid == "" {
		return &types.AckResponse{}, errors.New("Security.DeleteCredentials: kuid is required")
	}

	type body struct {
		Strategy string `json:"strategy"`
	}
	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "security",
		Action:     "deleteCredentials",
		Body:       &body{strategy},
		Id:         kuid,
	}

	go s.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error.Message != "" {
		return &types.AckResponse{}, errors.New(res.Error.Message)
	}

	ack := &types.AckResponse{}
	json.Unmarshal(res.Result, ack)

	return ack, nil
}
