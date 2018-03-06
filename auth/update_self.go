package auth

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/security"
	"github.com/kuzzleio/sdk-go/types"
)

//UpdateSelf updates the current User object in Kuzzle's database layer.
func (a *Auth) UpdateSelf(data json.RawMessage, options types.QueryOptions) (*security.User, error) {
	if options == nil {
		options = types.NewQueryOptions()
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "auth",
		Action:     "updateSelf",
		Body:       data,
	}
	go a.k.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return nil, res.Error
	}

	type jsonUser struct {
		Id         string          `json:"_id"`
		Source     json.RawMessage `json:"_source"`
		ProfileIds []string        `json:"profileIds"`
	}
	u := &jsonUser{}
	json.Unmarshal(res.Result, u)

	var content map[string]interface{}
	json.Unmarshal(u.Source, &content)
	user := &security.User{
		Id:         u.Id,
		Content:    content,
		ProfileIds: u.ProfileIds,
	}
	return user, nil
}
