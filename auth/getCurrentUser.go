package auth

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/security"
	"github.com/kuzzleio/sdk-go/types"
)

func (a *Auth) GetCurrentUser() (*security.User, error) {
	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "auth",
		Action:     "getCurrentUser",
	}

	go a.k.Query(query, nil, result)

	res := <-result

	if res.Error != nil {
		return nil, res.Error
	}

	type jsonUser struct {
		Id     string          `json:"_id"`
		Source json.RawMessage `json:"_source"`
	}
	ju := &jsonUser{}
	json.Unmarshal(res.Result, ju)

	var unmarsh map[string]interface{}
	json.Unmarshal(ju.Source, &unmarsh)
	u := &security.User{
		Id:      ju.Id,
		Content: unmarsh,
	}
	return u, nil
}
