package security

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
	"reflect"
)

// UpdateCredentials updates credentials of the specified strategy for the given user.
func (s Security) UpdateCredentials(strategy string, kuid string, credentials interface{}, options types.QueryOptions) (map[string]interface{}, error) {
	if strategy == "" {
		return nil, errors.New("Security.UpdateCredentials: strategy is required")
	}

	if kuid == "" {
		return nil, errors.New("Security.UpdateCredentials: kuid is required")
	}

	result := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Controller: "security",
		Action:     "updateCredentials",
		Body:       credentials,
		Strategy:   strategy,
		Id:         kuid,
	}
	go s.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error.Message != "" {
		return nil, errors.New(res.Error.Message)
	}

	ref := reflect.New(reflect.TypeOf(credentials)).Elem().Interface()
	json.Unmarshal(res.Result, &ref)

	return ref.(map[string]interface{}), nil
}
