package security

import (
	"github.com/kuzzleio/sdk-go/types"
)

func (s *Security) NewUser(id string, content *types.UserData) *User {
	u := &User{
		Id:       id,
		Security: s,
	}

	if content != nil {
		u.Content = content.Content
		u.ProfileIds = content.ProfileIds
	}

	return u
}
