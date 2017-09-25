package role

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
)

type SecurityRole struct {
	Kuzzle *kuzzle.Kuzzle
}

type Role struct {
	Id     string           `json:"_id"`
	Source json.RawMessage  `json:"_source"`
	Meta   types.KuzzleMeta `json:"_meta"`
	SR     SecurityRole
}

type RoleSearchResult struct {
	Hits  []Role `json:"hits"`
	Total int    `json:"total"`
}

// SetContent replaces the content of the Role object.
func (r *Role) SetContent(controllers types.Controllers) Role {
	r.Source, _ = json.Marshal(controllers)

	return *r
}

// Save creates or replaces the role in Kuzzle's database layer.
func (r Role) Save(options types.QueryOptions) (Role, error) {
	if options == nil {
		options = types.NewQueryOptions()
	}

	return r.SR.Create(r.Id, types.Controllers{Controllers: r.Controllers()}, options.SetIfExist("replace"))
}

// Update performs a partial content update on this object.
func (r Role) Update(content types.Controllers, options types.QueryOptions) (Role, error) {
	return r.SR.Update(r.Id, content, options)
}

// Delete this profile from Kuzzle.
func (r Role) Delete(options types.QueryOptions) (string, error) {
	return r.SR.Delete(r.Id, options)
}

// Controllers returns the role policies
func (r Role) Controllers() map[string]types.Controller {
	var controllers = types.Controllers{}
	json.Unmarshal(r.Source, &controllers)

	return controllers.Controllers
}

// Fetch retrieves a Role using its provided unique id.
func (sr SecurityRole) Fetch(id string, options types.QueryOptions) (Role, error) {
	if id == "" {
		return Role{}, errors.New("Security.Role.Fetch: role id required")
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
		return Role{}, errors.New(res.Error.Message)
	}

	role := Role{}
	json.Unmarshal(res.Result, &role)
	role.SR = sr

	return role, nil
}

// Search executes a search on Roles according to filters.
func (sr SecurityRole) Search(filters interface{}, options types.QueryOptions) (RoleSearchResult, error) {
	ch := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Controller: "security",
		Action:     "searchRoles",
		Body:       filters,
	}

	if options != nil {
		query.From = options.GetFrom()
		query.Size = options.GetSize()

		scroll := options.GetScroll()
		if scroll != "" {
			query.Scroll = scroll
		}
	}

	go sr.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Message != "" {
		return RoleSearchResult{}, errors.New(res.Error.Message)
	}

	searchResult := RoleSearchResult{}
	json.Unmarshal(res.Result, &searchResult)

	return searchResult, nil
}

// Create a new Role in Kuzzle.
func (sr SecurityRole) Create(id string, controllers types.Controllers, options types.QueryOptions) (Role, error) {
	if id == "" {
		return Role{}, errors.New("Security.Role.Create: role id required")
	}

	action := "createRole"

	if options != nil {
		if options.GetIfExist() == "replace" {
			action = "createOrReplaceRole"
		} else if options.GetIfExist() != "error" {
			return Role{}, errors.New(fmt.Sprintf("Invalid value for the 'ifExist' option: '%s'", options.GetIfExist()))
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
		return Role{}, errors.New(res.Error.Message)
	}

	role := Role{}
	json.Unmarshal(res.Result, &role)
	role.SR = sr

	return role, nil
}

// Update a Role in Kuzzle.
func (sr SecurityRole) Update(id string, controllers types.Controllers, options types.QueryOptions) (Role, error) {
	if id == "" {
		return Role{}, errors.New("Security.Role.Update: role id required")
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
		return Role{}, errors.New(res.Error.Message)
	}

	role := Role{}
	json.Unmarshal(res.Result, &role)
	role.SR = sr

	return role, nil
}

// Delete a Role in Kuzzle.
// There is a small delay between role deletion and their deletion in our advanced search layer, usually a couple of seconds.
// This means that a role that has just been deleted will still be returned by this function.
func (sr SecurityRole) Delete(id string, options types.QueryOptions) (string, error) {
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
