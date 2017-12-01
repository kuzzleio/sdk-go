package kuzzle

import (
	"github.com/kuzzleio/sdk-go/security"
)

// WhoAmI gets the rights array for the currently logged user.
func (k *Kuzzle) WhoAmI() (*security.User, error) {
	return k.Security.WhoAmI()
}
