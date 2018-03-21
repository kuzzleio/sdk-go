package auth

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

// GetMyRights gets the rights array for the currently logged user.
func (a *Auth) GetMyRights(options types.QueryOptions) ([]*types.UserRights, error) {
	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "auth",
		Action:     "getMyRights",
	}

	type rights struct {
		Hits []*types.UserRights `json:"hits"`
	}

	go a.kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return nil, res.Error
	}

	r := rights{}
	json.Unmarshal(res.Result, &r)

	return r.Hits, nil
}
