package kuzzle

import (
	"github.com/kuzzleio/sdk-go/security"
	"github.com/kuzzleio/sdk-go/types"
)

// UpdateSelf update the currently authenticated user information.
func (k *Kuzzle) UpdateSelf(content *types.UserData, options types.QueryOptions) (*security.User, error) {
	return k.Security.UpdateSelf(content, options)
}
