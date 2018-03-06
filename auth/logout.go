package auth

import (
	"github.com/kuzzleio/sdk-go/types"
)

// Logout logs the user out.
func (a *Auth) Logout() error {
	q := &types.KuzzleRequest{
		Controller: "auth",
		Action:     "logout",
	}
	result := make(chan *types.KuzzleResponse)

	go a.k.Query(q, nil, result)

	res := <-result

	if res.Error != nil {
		return res.Error
	}

	a.k.SetJwt("")

	return nil
}
