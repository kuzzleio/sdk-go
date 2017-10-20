package kuzzle

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
	"reflect"
)

// UpdateMyCredentials update credentials of the specified strategy for the current user.
func (k Kuzzle) UpdateMyCredentials(strategy string, credentials interface{}, options types.QueryOptions) (map[string]interface{}, error) {
	if strategy == "" {
		return nil, types.NewError("Kuzzle.UpdateMyCredentials: strategy is required", 400)
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

	if res.Error != nil {
		return nil, res.Error
	}

	ref := reflect.New(reflect.TypeOf(credentials)).Elem().Interface()
	json.Unmarshal(res.Result, &ref)

	return ref.(map[string]interface{}), nil
}
