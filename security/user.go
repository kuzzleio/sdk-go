package security

import (
	"github.com/kuzzleio/sdk-go/types"
)

type User struct {
	Id         string                 `json:"_id"`
	Content    map[string]interface{} `json:"_source"`
	ProfileIds []string
	Security   *Security
}

type UserSearchResult struct {
	Hits     []*User
	Total    int
	ScrollId string
}

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
