package kuzzle

import (
	"github.com/kuzzleio/sdk-go/types"
)

// Logout logs the user out.
func (k Kuzzle) Logout() error {
	q := &types.KuzzleRequest{
		Controller: "auth",
		Action:     "logout",
	}
	result := make(chan *types.KuzzleResponse)

	go k.Query(q, nil, result)

	res := <-result

	if res.Error != nil {
		return res.Error
	}

	k.jwt = ""

	return nil
}
