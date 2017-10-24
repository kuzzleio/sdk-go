package security

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
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
