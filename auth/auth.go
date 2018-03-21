package auth

import "github.com/kuzzleio/sdk-go/types"

type Auth struct {
	kuzzle types.IKuzzle
}

func NewAuth(k types.IKuzzle) *Auth {
	return &Auth{k}
}
