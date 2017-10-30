package security

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

type jsonUser struct {
	Id     string          `json:"_id"`
	Source json.RawMessage `json:"_source"`
}

type jsonUserSearchResult struct {
	Total    int         `json:"total"`
	Hits     []*jsonUser `json:"hits"`
	ScrollId string      `json:"scrollId"`
}

func (s *Security) jsonUserToUser(j *jsonUser) *User {
	u := &User{
		Id:       j.Id,
		Security: s,
	}

	userData := &types.UserData{}
	json.Unmarshal(j.Source, userData)
	if userData.ProfileIds != nil {
		u.ProfileIds = userData.ProfileIds
	}

	m := map[string]interface{}{}
	json.Unmarshal(j.Source, &m)

	for k, v := range m {
		if k != "profileIds" && k != "credentials" && v != nil {
			if u.Content == nil {
				u.Content = make(map[string]interface{})
			}
			u.Content[k] = v
		}
	}

	return u
}
