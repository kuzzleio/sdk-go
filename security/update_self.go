package security

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

func (s *Security) UpdateSelf(data *types.UserData, options types.QueryOptions) (*User, error) {
	// using a dummy user is marginally helpful
	u := &User{
		Security: s,
	}

	if data != nil {
		u.Content = data.Content
		u.ProfileIds = data.ProfileIds
	}

	body := u.getFlatBody()

	if options == nil {
		options = types.NewQueryOptions()
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "auth",
		Action:     "updateSelf",
		Body:       body,
	}
	go s.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return nil, res.Error
	}

	jsonUser := &jsonUser{}
	json.Unmarshal(res.Result, jsonUser)

	updatedUser := u.Security.jsonUserToUser(jsonUser)

	u.Id = updatedUser.Id
	u.Content = updatedUser.Content
	u.ProfileIds = updatedUser.ProfileIds

	return u, nil
}
