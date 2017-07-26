package kuzzle

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
	"reflect"
)

/*
 * Create credentials of the specified strategy for the current user.
 */
func (k Kuzzle) CreateMyCredentials(strategy string, credentials interface{}, options *types.Options) (map[string]interface{}, error) {
	if strategy == "" {
		return nil, errors.New("Kuzzle.CreateMyCredentials: strategy is required")
	}

	result := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Controller: "auth",
		Action:     "createMyCredentials",
		Body:       credentials,
		Strategy:   strategy,
	}
	go k.Query(query, options, result)

	res := <-result

	if res.Error.Message != "" {
		return nil, errors.New(res.Error.Message)
	}

	ref := reflect.New(reflect.TypeOf(credentials)).Elem().Interface()
	json.Unmarshal(res.Result, &ref)

	return ref.(map[string]interface{}), nil
}
