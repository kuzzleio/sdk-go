package kuzzle

import (
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

// Logs the user out.
func (k *Kuzzle) Logout() error {
	q := types.KuzzleRequest{
		Controller: "auth",
		Action:     "logout",
	}
	result := make(chan types.KuzzleResponse)

	go k.Query(q, nil, result)

	res := <-result

	if res.Error.Message != "" {
		return errors.New(res.Error.Message)
	}

	k.jwt = ""

	return nil
}
