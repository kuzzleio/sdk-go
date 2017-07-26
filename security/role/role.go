package role

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
)

type SecurityRole struct {
	Kuzzle kuzzle.Kuzzle
}

/*
  Retrieves a Role using its provided unique id.
*/
func (sr SecurityRole) Fetch(id string, options *types.Options) (types.Role, error) {
	if id == "" {
		return types.Role{}, errors.New("Security.Role.Fetch: role id required")
	}

	ch := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Controller: "security",
		Action:     "getRole",
		Id:         id,
	}
	go sr.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Message != "" {
		return types.Role{}, errors.New(res.Error.Message)
	}

	role := types.Role{}
	json.Unmarshal(res.Result, &role)

	return role, nil
}

/*
  Create a new Role in Kuzzle.
*/
func (sr SecurityRole) Create(id string, controllers types.Controllers, options *types.Options) (types.Role, error) {
	if id == "" {
		return types.Role{}, errors.New("Security.Role.Create: role id required")
	}

	action := "createRole"

	if options != nil {
		if options.IfExist == "replace" {
			action = "createOrReplaceRole"
		} else if options.IfExist != "error" {
			return types.Role{}, errors.New(fmt.Sprintf("Invalid value for the 'ifExist' option: '%s'", options.IfExist))
		}
	}

	ch := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Controller: "security",
		Action:     action,
		Body:       controllers,
		Id:         id,
	}
	go sr.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Message != "" {
		return types.Role{}, errors.New(res.Error.Message)
	}

	role := types.Role{}
	json.Unmarshal(res.Result, &role)

	return role, nil
}

/*
  Update a Role in Kuzzle.
*/
func (sr SecurityRole) Update(id string, controllers types.Controllers, options *types.Options) (types.Role, error) {
	if id == "" {
		return types.Role{}, errors.New("Security.Role.Update: role id required")
	}

	ch := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Controller: "security",
		Action:     "updateRole",
		Body:       controllers,
		Id:         id,
	}
	go sr.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Message != "" {
		return types.Role{}, errors.New(res.Error.Message)
	}

	role := types.Role{}
	json.Unmarshal(res.Result, &role)

	return role, nil
}

/*
 * Delete a Role in Kuzzle.
 *
 * There is a small delay between role deletion and their deletion in our advanced search layer, usually a couple of seconds.
 * This means that a role that has just been deleted will still be returned by this function.
 */
func (sr SecurityRole) Delete(id string, options *types.Options) (string, error) {
	if id == "" {
		return "", errors.New("Security.Role.Delete: role id required")
	}

	ch := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Controller: "security",
		Action:     "deleteRole",
		Id:         id,
	}
	go sr.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Message != "" {
		return "", errors.New(res.Error.Message)
	}

	shardResponse := types.ShardResponse{}
	json.Unmarshal(res.Result, &shardResponse)

	return shardResponse.Id, nil
}
