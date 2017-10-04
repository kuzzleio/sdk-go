package kuzzle

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
	"reflect"
)

// UpdateMyCredentials update credentials of the specified strategy for the current user.
func (k Kuzzle) UpdateMyCredentials(strategy string, credentials interface{}, options types.QueryOptions) (map[string]interface{}, error) {
	if strategy == "" {
		return nil, errors.New("Kuzzle.UpdateMyCredentials: strategy is required")
	}

	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "auth",
		Action:     "updateMyCredentials",
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
