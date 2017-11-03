package security

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

func (s *Security) UpdateSelf(content *types.UserData, options types.QueryOptions) (*User, error) {
	// using a dummy user is marginally helpful
	u := &User{
		Security: s,
	}

	if content != nil {
		u.Content = content.Content
		u.addProfileIds(content.ProfileIds...)
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
