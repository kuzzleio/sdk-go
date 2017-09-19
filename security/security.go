package security

import (
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/security/profile"
	"github.com/kuzzleio/sdk-go/security/role"
	"github.com/kuzzleio/sdk-go/security/user"
)

type Security struct {
	Kuzzle  *kuzzle.Kuzzle
	Profile profile.SecurityProfile
	Role    role.SecurityRole
	User    user.SecurityUser
}

func NewSecurity(kuzzle *kuzzle.Kuzzle) *Security {
	return &Security{
		Kuzzle:  kuzzle,
		Profile: profile.SecurityProfile{Kuzzle: kuzzle},
		Role:    role.SecurityRole{Kuzzle: kuzzle},
		User:    user.SecurityUser{Kuzzle: kuzzle},
	}
}
