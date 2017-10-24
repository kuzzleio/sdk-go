package security

import (
	"encoding/json"
	"fmt"

	"github.com/kuzzleio/sdk-go/types"
)

type Role struct {
	Id          string `json:"_id"`
	Controllers map[string]*types.Controller
	Security    *Security
}

type RoleSearchResult struct {
	Hits  []*Role
	Total int
}

// Delete a role from Kuzzle
func (r *Role) Delete(options types.QueryOptions) (string, error) {
	return r.Security.rawDelete("deleteRole", r.Id, options)
}

// Save creates or replaces the role in Kuzzle
func (r *Role) Save(options types.QueryOptions) (*Role, error) {
	action := "createOrReplaceRole"

	if options == nil && r.Id == "" {
		action = "createRole"
	}

	if options != nil {
		if options.GetIfExist() == "error" {
			action = "createRole"
		} else if options.GetIfExist() != "replace" {
			return nil, types.NewError(fmt.Sprintf("Invalid value for 'ifExist' option: '%s'", options.GetIfExist()), 400)
		}
	}

	return r.persist(action, options)
}

// Update sets the role controllers and persists it in Kuzzle
// NB: The role must exist in Kuzzle.
func (r *Role) Update(controllers *types.Controllers, options types.QueryOptions) (*Role, error) {
	r.Controllers = controllers.Controllers
	return r.persist("updateRole", options)
}

func (r *Role) persist(action string, options types.QueryOptions) (*Role, error) {
	if options == nil {
		options = types.NewQueryOptions()
	}

	if action != "createRole" && r.Id == "" {
		return nil, types.NewError("Role.Save: role id is required", 400)
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "security",
		Action:     action,
		Body: types.Controllers{
			Controllers: r.Controllers,
		},
		Id: r.Id,
	}
	go r.Security.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return nil, res.Error
	}

	jsonRole := &jsonRole{}
	json.Unmarshal(res.Result, jsonRole)

	r.Controllers = jsonRole.Source.Controllers
	r.Id = jsonRole.Id

	return r, nil
}
