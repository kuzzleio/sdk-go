package security

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// DeleteCredentials deletes credentials of the specified strategy for the given user.
func (s Security) DeleteCredentials(strategy string, kuid string, options types.QueryOptions) (*types.AckResponse, error) {
	if strategy == "" {
		return nil, types.NewError("Security.DeleteCredentials: strategy is required", 400)
	}

	if kuid == "" {
		return nil, types.NewError("Security.DeleteCredentials: kuid is required", 400)
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

	if res.Error != nil {
		return nil, res.Error
	}

	ack := &types.AckResponse{}
	json.Unmarshal(res.Result, ack)

	return ack, nil
}
