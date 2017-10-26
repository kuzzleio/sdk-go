package security_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/kuzzleio/sdk-go/connection/websocket"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/security"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func TestProfileAddPolicy(t *testing.T) {
	id := "profileId"

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "getProfile", parsedQuery.Action)
			assert.Equal(t, id, parsedQuery.Id)

			p := &security.Profile{
				Id: id,
				Policies: []*types.Policy{
					{RoleId: "admin"},
					{RoleId: "other"},
				},
			}
			r, _ := security.ProfileToJson(p)
			return &types.KuzzleResponse{Result: r}
		},
	}

	k, _ := kuzzle.NewKuzzle(c, nil)
	p, _ := k.Security.FetchProfile(id, nil)

	policy := &types.Policy{
		RoleId: "roleId",
		RestrictedTo: []*types.PolicyRestriction{
			{Index: "index"},
			{Index: "other-index", Collections: []string{"foo", "bar"}},
		},
	}

	p.AddPolicy(policy)

	assert.Equal(t, []*types.Policy{
		{RoleId: "admin"},
		{RoleId: "other"},
		{
			RoleId: "roleId",
			RestrictedTo: []*types.PolicyRestriction{
				{Index: "index"},
				{Index: "other-index", Collections: []string{"foo", "bar"}},
			},
		},
	}, p.Policies)
}

func ExampleProfile_AddPolicy() {
	id := "profileId"
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	p, _ := k.Security.FetchProfile(id, nil)

	policy := types.Policy{
		RoleId: "roleId",
		RestrictedTo: []*types.PolicyRestriction{
			{Index: "index"},
			{Index: "other-index", Collections: []string{"foo", "bar"}},
		},
	}

	p.AddPolicy(&policy)

	fmt.Println(p.Policies)
}

func TestProfileSaveEmptyId(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "createProfile", parsedQuery.Action)
			return &types.KuzzleResponse{Result: []byte{}}
		},
	}

	k, _ := kuzzle.NewKuzzle(c, nil)
	profile := k.Security.NewProfile("", nil)

	_, err := profile.Save(nil)
	assert.Nil(t, err)
}

func TestProfileSaveErrorOption(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "createProfile", parsedQuery.Action)

			return &types.KuzzleResponse{Result: []byte{}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	profile := k.Security.NewProfile("profileid", nil)
	options := types.NewQueryOptions()
	options.SetIfExist("error")

	_, err := profile.Save(options)
	assert.Nil(t, err)

}

func TestProfileSaveInvalidOption(t *testing.T) {
	c := &internal.MockedConnection{}
	k, _ := kuzzle.NewKuzzle(c, nil)
	profile := k.Security.NewProfile("profileid", nil)
	options := types.NewQueryOptions()
	options.SetIfExist("invalid")

	_, err := profile.Save(options)
	assert.NotNil(t, err)
	assert.Equal(t, 400, err.(*types.KuzzleError).Status)
	assert.Equal(t, "Invalid value for 'ifExist' option: 'invalid'", err.(*types.KuzzleError).Message)
}

func TestProfileSaveEmptyIdWithReplaceOption(t *testing.T) {
	c := &internal.MockedConnection{}
	k, _ := kuzzle.NewKuzzle(c, nil)
	profile := k.Security.NewProfile("", nil)
	options := types.NewQueryOptions()
	options.SetIfExist("replace")

	_, err := profile.Save(options)
	assert.NotNil(t, err)
	assert.Equal(t, 400, err.(*types.KuzzleError).Status)
	assert.Equal(t, "Profile.createOrReplaceProfile: id is required", err.(*types.KuzzleError).Message)
}

func TestProfileSave(t *testing.T) {
	id := "profileId"
	expectedNewProfile := &security.Profile{
		Id: id,
		Policies: []*types.Policy{
			{RoleId: "newRoleId"},
			{RoleId: "newRoleId", RestrictedTo: []*types.PolicyRestriction{
				{Index: "index", Collections: []string{"foo", "bar"}},
			}},
		},
	}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "createOrReplaceProfile", parsedQuery.Action)
			assert.Equal(t, id, parsedQuery.Id)
			assert.Equal(t, map[string]interface{}{
				"policies": []interface{}{
					map[string]interface{}{"roleId": "newRoleId"},
					map[string]interface{}{
						"roleId": "otherRoleId",
						"restrictedTo": []interface{}{
							map[string]interface{}{
								"index": "index",
								"collections": []interface{}{
									"foo",
									"bar",
								},
							},
						},
					},
				},
			}, parsedQuery.Body)

			r, _ := security.ProfileToJson(expectedNewProfile)
			return &types.KuzzleResponse{Result: r}
		},
	}

	k, _ := kuzzle.NewKuzzle(c, nil)
	p := k.Security.NewProfile(id, []*types.Policy{
		{RoleId: "newRoleId"},
		{RoleId: "otherRoleId", RestrictedTo: []*types.PolicyRestriction{{Index: "index", Collections: []string{"foo", "bar"}}}},
	})
	_, err := p.Save(nil)
	assert.Nil(t, err)
}

func ExampleProfile_Save() {
	id := "profileId"
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	p, _ := k.Security.FetchProfile(id, nil)

	newPolicies := []types.Policy{
		{RoleId: "newRoleId"},
		{RoleId: "otherRoleId", RestrictedTo: []*types.PolicyRestriction{{Index: "index", Collections: []string{"foo", "bar"}}}},
	}

	for _, policy := range newPolicies {
		p.AddPolicy(&policy)
	}
	res, err := p.Save(nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res.Id, res.Policies)
}

func TestProfileUpdate(t *testing.T) {
	id := "profileId"
	expectedUpdatedProfile := &security.Profile{
		Id: id,
		Policies: []*types.Policy{
			{RoleId: "boringNewRoleId"},
		},
	}
	callCount := 0

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			if callCount == 0 {
				callCount++
				assert.Equal(t, "security", parsedQuery.Controller)
				assert.Equal(t, "getProfile", parsedQuery.Action)
				assert.Equal(t, id, parsedQuery.Id)

				p := &security.Profile{
					Id: id,
					Policies: []*types.Policy{
						{RoleId: "admin"},
						{RoleId: "other"},
					},
				}
				r, _ := security.ProfileToJson(p)
				return &types.KuzzleResponse{Result: r}
			}
			if callCount == 1 {
				callCount++
				assert.Equal(t, "security", parsedQuery.Controller)
				assert.Equal(t, "updateProfile", parsedQuery.Action)
				assert.Equal(t, id, parsedQuery.Id)
				assert.Equal(t, map[string]interface{}{
					"policies": []interface{}{
						map[string]interface{}{"roleId": "boringNewRoleId"},
					},
				}, parsedQuery.Body)

				r, _ := security.ProfileToJson(expectedUpdatedProfile)
				return &types.KuzzleResponse{Result: r}
			}

			return &types.KuzzleResponse{Result: nil}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	p, _ := k.Security.FetchProfile(id, nil)

	newPolicies := []*types.Policy{
		{RoleId: "boringNewRoleId"},
	}

	updatedProfile, _ := p.Update(newPolicies, nil)

	assert.Equal(t, expectedUpdatedProfile.Id, updatedProfile.Id)
	assert.Equal(t, newPolicies, updatedProfile.Policies)
}

func ExampleProfile_Update() {
	id := "profileId"
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	p, _ := k.Security.FetchProfile(id, nil)

	newPolicies := []*types.Policy{
		{RoleId: "boringNewRoleId"},
	}

	res, err := p.Update(newPolicies, nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res.Id, res.Policies)
}

func TestProfileDelete(t *testing.T) {
	id := "SomeMenJustWantToWatchTheWorldBurn"
	callCount := 0

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			callCount++
			if callCount == 1 {
				assert.Equal(t, "security", parsedQuery.Controller)
				assert.Equal(t, "getProfile", parsedQuery.Action)
				assert.Equal(t, id, parsedQuery.Id)

				p := &security.Profile{
					Id: id,
					Policies: []*types.Policy{
						{RoleId: "admin"},
						{RoleId: "other"},
					},
				}
				r, _ := security.ProfileToJson(p)
				return &types.KuzzleResponse{Result: r}
			}
			if callCount == 2 {
				assert.Equal(t, "security", parsedQuery.Controller)
				assert.Equal(t, "deleteProfile", parsedQuery.Action)
				assert.Equal(t, id, parsedQuery.Id)

				res := types.ShardResponse{Id: id}
				r, _ := json.Marshal(res)
				return &types.KuzzleResponse{Result: r}
			}

			return &types.KuzzleResponse{Result: nil}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	p, _ := k.Security.FetchProfile(id, nil)

	inTheEnd, _ := p.Delete(nil)

	assert.Equal(t, id, inTheEnd)
}

func ExampleProfile_Delete() {
	id := "SomeMenJustWantToWatchTheWorldBurn"
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	p, _ := k.Security.FetchProfile(id, nil)
	res, err := p.Delete(nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res)
}
