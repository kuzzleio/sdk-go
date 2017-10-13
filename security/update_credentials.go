package security

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
	"reflect"
)

// UpdateCredentials updates credentials of the specified strategy for the given user.
func (s Security) UpdateCredentials(strategy string, kuid string, credentials interface{}, options types.QueryOptions) (map[string]interface{}, error) {
	if strategy == "" {
		return nil, types.NewError("Security.UpdateCredentials: strategy is required", 400)
	}

	if kuid == "" {
		return nil, types.NewError("Security.UpdateCredentials: kuid is required", 400)
	}

	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "security",
		Action:     "updateCredentials",
		Body:       credentials,
		Strategy:   strategy,
		Id:         kuid,
	}
	go s.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return nil, res.Error
	}

	ref := reflect.New(reflect.TypeOf(credentials)).Elem().Interface()
	json.Unmarshal(res.Result, &ref)

	return ref.(map[string]interface{}), nil
}
