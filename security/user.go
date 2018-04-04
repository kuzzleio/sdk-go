package security

import (
	"github.com/kuzzleio/sdk-go/types"
)

type User struct {
	Id         string                 `json:"_id"`
	Content    map[string]interface{} `json:"_source"`
	ProfileIds []string
}

type UserSearchResult struct {
	Hits     []*User
	Total    int
	ScrollId string
}

func NewUser(id string, content *types.UserData) *User {
	u := &User{
		Id: id,
	}

	if content != nil {
		u.Content = content.Content
		u.ProfileIds = content.ProfileIds
	}

	return u
}
