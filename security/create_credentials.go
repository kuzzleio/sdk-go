package security

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
	"reflect"
)

// CreateCredentials creates credential of the specified strategy for the given user.
func (s Security) CreateCredentials(strategy string, kuid string, credentials interface{}, options types.QueryOptions) (map[string]interface{}, error) {
	if strategy == "" {
		return nil, types.NewError("Security.CreateCredentials: strategy is required")
	}

	if kuid == "" {
		return nil, types.NewError("Security.CreateCredentials: kuid is required")
	}

	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "security",
		Action:     "createCredentials",
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
