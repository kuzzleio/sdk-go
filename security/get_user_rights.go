package security

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

// GetRights returns the rights of an User using its provided unique id.
func (s *Security) GetUserRights(kuid string, options types.QueryOptions) ([]*types.UserRights, error) {
	if kuid == "" {
		return nil, errors.New("Security.User.GetRights: user id required")
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "security",
		Action:     "getUserRights",
		Id:         kuid,
	}
	go s.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return nil, errors.New(res.Error.Message)
	}

	type response struct {
		UserRights []*types.UserRights `json:"hits"`
	}
	userRights := response{}
	json.Unmarshal(res.Result, &userRights)

	return userRights.UserRights, nil
}
