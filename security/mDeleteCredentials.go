package security

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

// MDeleteCredentials deletes all credentials matching with given ids
func (s *Security) MDeleteCredentials(ids []string, options types.QueryOptions) ([]string, error) {
	if len(ids) == 0 {
		return nil, types.NewError("Security.MDeleteCredentials: ids array can't be nil", 400)
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "security",
		Action:     "mDeleteCredentials",
		Body: struct {
			Ids []string `json:"ids"`
		}{ids},
	}
	go s.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return nil, res.Error
	}

	var mDeletedIds []string
	json.Unmarshal(res.Result, &mDeletedIds)

	return mDeletedIds, nil

}
