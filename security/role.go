package security

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
	"fmt"
)

type Role struct {
	Id	string						`json:"_id"`
	Controllers	map[string]types.Controller
	Kuzzle types.IKuzzle
}

type RoleSearchResult struct {
	Hits  []*Role
	Total int
}

func (r *Role) Delete(options types.QueryOptions) (string, error) {
	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "security",
		Action:     "deleteRole",
		Id:         r.Id,
	}
	go r.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return "", errors.New(res.Error.Message)
	}

	shardResponse := types.ShardResponse{}
	json.Unmarshal(res.Result, &shardResponse)

	return shardResponse.Id, nil
}

func (r *Role) Save(options types.QueryOptions) (*Role, error) {
	action := "createOrReplaceRole"

	if options == nil && r.Id == "" {
		action = "createRole"
	}

	if options != nil {
		if options.GetIfExist() == "error" {
			action = "createRole"
		} else if options.GetIfExist() != "replace" {
			return nil, errors.New(fmt.Sprintf("Invalid value for 'ifExist' option: '%s'", options.GetIfExist()))
		}
	}

	return r.persist(action, options)
}

func (r *Role) SetContent(controllers *types.Controllers) *Role {
	r.Controllers = controllers.Controllers
	return r
}

func (r *Role) Update(controllers *types.Controllers, options types.QueryOptions) (*Role, error) {
	r.SetContent(controllers)
	return r.persist("updateRole", options)
}

func (r *Role) persist(action string, options types.QueryOptions) (*Role, error) {
	if options == nil {
		options = types.NewQueryOptions()
	}

	if action != "createRole" && r.Id == "" {
		return nil, errors.New("Role.Save: role id is required")
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "security",
		Action: action,
		Body: types.Controllers{
			Controllers: r.Controllers,
		},
		Id: r.Id,
	}
	go r.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return nil, errors.New(res.Error.Message)
	}

	jsonRole := &jsonRole{}
	json.Unmarshal(res.Result, jsonRole)

	r.Controllers = jsonRole.Source.Controllers

	return r, nil
}



