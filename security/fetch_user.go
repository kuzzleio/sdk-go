package security

import (
	"github.com/kuzzleio/sdk-go/types"
	"encoding/json"
)

func (s *Security) FetchUser(id string, options types.QueryOptions) (*User, error) {
	res, err := s.rawFetch("getUser", id, options)

	if err != nil {
		return nil, err
	}

	jsonUser := &jsonUser{}
	json.Unmarshal(res, jsonUser)

	return s.jsonUserToUser(jsonUser), nil
}
