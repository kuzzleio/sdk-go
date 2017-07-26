package security

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
	"reflect"
)

/*
 * Create credentials of the specified strategy for the given user.
 */
func (s Security) CreateCredentials(strategy string, kuid string, credentials interface{}, options *types.Options) (map[string]interface{}, error) {
	if strategy == "" {
		return nil, errors.New("Security.CreateCredentials: strategy is required")
	}

	if kuid == "" {
		return nil, errors.New("Security.CreateCredentials: kuid is required")
	}

	result := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Controller: "security",
		Action:     "createCredentials",
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
