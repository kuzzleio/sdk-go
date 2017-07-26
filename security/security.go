package security

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/security/profile"
	"github.com/kuzzleio/sdk-go/security/role"
	"github.com/kuzzleio/sdk-go/security/user"
	"github.com/kuzzleio/sdk-go/types"
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
		Profile: profile.SecurityProfile{Kuzzle: *kuzzle},
		Role:    role.SecurityRole{Kuzzle: *kuzzle},
		User:    user.SecurityUser{Kuzzle: *kuzzle},
	}
}

func (s Security) GetAllCredentialFields(options *types.Options) (types.CredentialFields, error) {
	ch := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Controller: "security",
		Action:     "getAllCredentialFields",
	}
	go s.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Message != "" {
		return types.CredentialFields{}, errors.New(res.Error.Message)
	}

	credentialFields := types.CredentialFields{}
	json.Unmarshal(res.Result, &credentialFields)

	return credentialFields, nil
}
