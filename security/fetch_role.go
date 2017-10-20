package security

import (
	"github.com/kuzzleio/sdk-go/types"
	"encoding/json"
)

func (s *Security) FetchRole(id string, options types.QueryOptions) (*Role, error) {
	res, err := s.rawFetch("getRole", id, options)

	if err != nil {
		return nil, err
	}

	jsonRole := &jsonRole{}
	json.Unmarshal(res, jsonRole)

	return s.jsonRoleToRole(jsonRole), nil
}
