package security_test

import (
	"encoding/json"
	"fmt"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
	"github.com/kuzzleio/sdk-go/connection/websocket"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/security"
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
				Policies: []*types.Policy {
					{ RoleId: "admin" },
					{ RoleId: "other"},
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
		AllowInternalIndex: true,
		RestrictedTo: []*types.PolicyRestriction {
			{ Index: "index" },
			{ Index: "other-index", Collections: []string{ "foo", "bar" } },
		},
	}

	p.AddPolicy(policy)

	assert.Equal(t, []*types.Policy{
		{RoleId: "admin"},
		{RoleId: "other"},
		{
			RoleId: "roleId",
			AllowInternalIndex: true,
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
		RoleId:             "roleId",
		AllowInternalIndex: true,
		RestrictedTo:       []*types.PolicyRestriction{
			{Index: "index"},
			{Index: "other-index", Collections: []string{"foo", "bar"}},
		},
	}

	p.AddPolicy(&policy)

	fmt.Println(p.Policies)
}

func TestProfileSave(t *testing.T) {
	id := "profileId"
	expectedNewProfile := &security.Profile{
		Id: id,
		Policies: []*types.Policy {
			{ RoleId: "newRoleId", AllowInternalIndex: true },
			{ RoleId: "newRoleId", RestrictedTo: []*types.PolicyRestriction {
				{ Index: "index", Collections: []string{ "foo", "bar" } },
			} },
		},
	}
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
					Policies: []*types.Policy {
						{ RoleId: "admin"},
						{ RoleId: "other"},
					},
				}
				r, _ := security.ProfileToJson(p)
				return &types.KuzzleResponse{Result: r}
			}
			if callCount == 2 {
				assert.Equal(t, "security", parsedQuery.Controller)
				assert.Equal(t, "createOrReplaceProfile", parsedQuery.Action)
				assert.Equal(t, id, parsedQuery.Id)

				r, _ := security.ProfileToJson(expectedNewProfile)
				return &types.KuzzleResponse{Result: r}
			}

			return &types.KuzzleResponse{Result: nil}
		},
	}

	k, _ := kuzzle.NewKuzzle(c, nil)
	p, _ := k.Security.FetchProfile(id, nil)
	newPolicies := []types.Policy{
		{RoleId: "newRoleId", AllowInternalIndex: true},
		{RoleId: "otherRoleId", RestrictedTo: []*types.PolicyRestriction{{Index: "index", Collections: []string{"foo", "bar"}}}},
	}

	for _, policy := range newPolicies {
		p.AddPolicy(&policy)
	}
	newProfile, _ := p.Save(nil)

	assert.Equal(t, expectedNewProfile.Id, newProfile.Id)
	assert.Equal(t, expectedNewProfile.Policies, newProfile.Policies)
}

func ExampleProfile_Save() {
	id := "profileId"
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	p, _ := k.Security.FetchProfile(id, nil)

	newPolicies := []types.Policy{
		{RoleId: "newRoleId", AllowInternalIndex: true},
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
		Policies: []*types.Policy {
			{ RoleId: "boringNewRoleId" },
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
					Policies: []*types.Policy {
						{ RoleId: "admin" },
						{ RoleId: "other" },
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

func TestFetchProfileEmptyId(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.MessageError{Message: "Security.Profile.Fetch: profile id required"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	_, err := k.Security.FetchProfile("", nil)

	assert.NotNil(t, err)
}

func TestFetchProfileError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	_, err := k.Security.FetchProfile("profileId", nil)
	assert.NotNil(t, err)
}

func TestFetchProfile(t *testing.T) {
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
				Policies: []*types.Policy {
					{ RoleId: "admin" },
					{ RoleId: "other" },
				},
			}
			r, _ := security.ProfileToJson(p)
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	res, _ := k.Security.FetchProfile(id, nil)

	assert.Equal(t, id, res.Id)
	assert.Equal(t, []*types.Policy{{RoleId: "admin"}, {RoleId: "other"}}, res.Policies)
}

func ExampleFetchProfile() {
	id := "profileId"
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	res, err := k.Security.FetchProfile(id, nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res.Id, res.Policies)
}

func TestSearchProfileError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	_, err := k.Security.SearchProfiles(nil, nil)
	assert.NotNil(t, err)
}

func TestSearchProfile(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "searchProfiles", parsedQuery.Action)

			return &types.KuzzleResponse{Result: []byte(`{
				"total": 42,
				"hits": [
					{ "_id": "profile42", "_source": { "policies": [{"roleId": "admin"}] } }
				]
			}`)}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	res, _ := k.Security.SearchProfiles(nil, nil)

	assert.Equal(t, 42, res.Total)
	assert.Equal(t, "profile42", res.Hits[0].Id)
	assert.Equal(t, []*types.Policy {
		{ RoleId: "admin" },
	}, res.Hits[0].Policies)
}

func ExampleSearchProfiles() {
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	res, err := k.Security.SearchProfiles(nil, nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res.Hits[0].Id, res.Hits[0].Policies)
}

func TestSearchProfileWithScroll(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "searchProfiles", parsedQuery.Action)

			return &types.KuzzleResponse{Result: []byte(`{
				"total": 42,
				"hits": [
					{"_id": "profile42", "_source": {"policies": [{"roleId": "admin"}]}}
				],
				"scrollId": "f00b4r"
			}`)}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	opts := types.NewQueryOptions()
	opts.SetFrom(2)
	opts.SetSize(4)
	opts.SetScroll("1m")

	res, _ := k.Security.SearchProfiles(nil, opts)
	assert.Equal(t, 42, res.Total)
	assert.Equal(t, 1, len(res.Hits))
	assert.Equal(t, []*types.Policy{
		{RoleId: "admin"},
	}, res.Hits[0].Policies)
	assert.Equal(t, "f00b4r", res.ScrollId)
	assert.Equal(t, "profile42", res.Hits[0].Id)
}

func TestScrollProfileEmptyScrollId(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.MessageError{Message: "Security.Profile.Scroll: scroll id required"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	_, err := k.Security.ScrollProfiles("", nil)
	assert.NotNil(t, err)
}

func TestScrollProfileError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	_, err := k.Security.ScrollProfiles("f00b4r", nil)
	assert.NotNil(t, err)
}

func TestScrollProfile(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "scrollProfiles", parsedQuery.Action)

			return &types.KuzzleResponse{Result: []byte(`{
				"total": 42,
				"hits": [
					{"_id": "profile42", "_source": {"policies": [{"roleId": "admin"}]}}
				]
			}`)}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := k.Security.ScrollProfiles("f00b4r", nil)
	assert.Equal(t, 42, res.Total)
	assert.Equal(t, 1, len(res.Hits))
	assert.Equal(t, []*types.Policy{
		{RoleId: "admin"},
	}, res.Hits[0].Policies)
	assert.Equal(t, "profile42", res.Hits[0].Id)
}

