package profile_test

import (
	"encoding/json"
	"fmt"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/security"
	"github.com/kuzzleio/sdk-go/security/profile"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
	"github.com/kuzzleio/sdk-go/connection/websocket"
)

func TestProfileAddPolicy(t *testing.T) {
	id := "profileId"

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "getProfile", parsedQuery.Action)
			assert.Equal(t, id, parsedQuery.Id)

			res := profile.Profile{Id: id, Source: []byte(`{"policies":[{"roleId":"admin"},{"roleId":"other"}]}`)}
			r, _ := json.Marshal(res)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	p, _ := security.NewSecurity(k).Profile.Fetch(id, nil)

	policy := types.Policy{
		RoleId:             "roleId",
		AllowInternalIndex: true,
		RestrictedTo:       []types.PolicyRestriction{{Index: "index"}, {Index: "other-index", Collections: []string{"foo", "bar"}}},
	}

	p.AddPolicy(policy)

	assert.Equal(t, []types.Policy{
		{RoleId: "admin"},
		{RoleId: "other"},
		{RoleId: "roleId", AllowInternalIndex: true, RestrictedTo: []types.PolicyRestriction{{Index: "index"}, {Index: "other-index", Collections: []string{"foo", "bar"}}}},
	}, p.GetPolicies())
}

func ExampleProfile_AddPolicy() {
	id := "profileId"
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	p, _ := security.NewSecurity(k).Profile.Fetch(id, nil)

	policy := types.Policy{
		RoleId:             "roleId",
		AllowInternalIndex: true,
		RestrictedTo:       []types.PolicyRestriction{{Index: "index"}, {Index: "other-index", Collections: []string{"foo", "bar"}}},
	}

	p.AddPolicy(policy)

	fmt.Println(p.GetPolicies())
}

func TestProfileGetPolicies(t *testing.T) {
	id := "profileId"

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "getProfile", parsedQuery.Action)
			assert.Equal(t, id, parsedQuery.Id)

			res := profile.Profile{Id: id, Source: []byte(`{"policies":[{"roleId":"admin"},{"roleId":"other"}]}`)}
			r, _ := json.Marshal(res)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	p, _ := security.NewSecurity(k).Profile.Fetch(id, nil)

	assert.Equal(t, []types.Policy{
		{RoleId: "admin"},
		{RoleId: "other"}}, p.GetPolicies())
}

func ExampleProfile_GetPolicies() {
	id := "profileId"
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)

	p, _ := security.NewSecurity(k).Profile.Fetch(id, nil)

	fmt.Println(p.GetPolicies())
}

func TestProfileSetPolicies(t *testing.T) {
	id := "profileId"

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "getProfile", parsedQuery.Action)
			assert.Equal(t, id, parsedQuery.Id)

			res := profile.Profile{Id: id, Source: []byte(`{"policies":[{"roleId":"admin"},{"roleId":"other"}]}`)}
			r, _ := json.Marshal(res)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	p, _ := security.NewSecurity(k).Profile.Fetch(id, nil)

	newPolicies := []types.Policy{
		{RoleId: "newRoleId", AllowInternalIndex: true},
		{RoleId: "otherRoleId", RestrictedTo: []types.PolicyRestriction{{Index: "index", Collections: []string{"foo", "bar"}}}},
	}

	p.SetPolicies(newPolicies)

	assert.Equal(t, newPolicies, p.GetPolicies())
}

func ExampleProfile_SetPolicies() {
	id := "profileId"
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)

	p, _ := security.NewSecurity(k).Profile.Fetch(id, nil)

	newPolicies := []types.Policy{
		{RoleId: "newRoleId", AllowInternalIndex: true},
		{RoleId: "otherRoleId", RestrictedTo: []types.PolicyRestriction{{Index: "index", Collections: []string{"foo", "bar"}}}},
	}

	p.SetPolicies(newPolicies)

	fmt.Println(p.GetPolicies())
}

func TestProfileSetContent(t *testing.T) {
	id := "profileId"

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "getProfile", parsedQuery.Action)
			assert.Equal(t, id, parsedQuery.Id)

			res := profile.Profile{Id: id, Source: []byte(`{"policies":[{"roleId":"admin"},{"roleId":"other"}]}`)}
			r, _ := json.Marshal(res)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	p, _ := security.NewSecurity(k).Profile.Fetch(id, nil)

	newContent := []byte(`{"policies":[{"roleId":"newRoleId","allowInternalIndex":true},{"roleId":"otherRoleId","restrictedTo":[{"index":"index","collections":["foo","bar"]}]}]}`)
	expectedPolicies := []types.Policy{
		{RoleId: "newRoleId", AllowInternalIndex: true},
		{RoleId: "otherRoleId", RestrictedTo: []types.PolicyRestriction{{Index: "index", Collections: []string{"foo", "bar"}}}},
	}

	p.SetContent(newContent)

	assert.Equal(t, json.RawMessage(newContent), p.Source)
	assert.Equal(t, expectedPolicies, p.GetPolicies())
}

func ExampleProfile_SetContent() {
	id := "profileId"
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)

	p, _ := security.NewSecurity(k).Profile.Fetch(id, nil)

	newContent := []byte(`{"policies":[{"roleId":"newRoleId","allowInternalIndex":true},{"roleId":"otherRoleId","restrictedTo":[{"index":"index","collections":["foo","bar"]}]}]}`)

	p.SetContent(newContent)

	fmt.Println(p.Source, p.GetPolicies())
}

func TestProfileSave(t *testing.T) {
	id := "profileId"
	expectedNewProfile := profile.Profile{Id: id, Source: []byte(`{"im":"emptyInside","policies":[{"roleId":"newRoleId","allowInternalIndex":true},{"roleId":"otherRoleId","restrictedTo":[{"index":"index","collections":["foo","bar"]}]}]}`)}
	callCount := 0

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			if callCount == 0 {
				callCount++
				assert.Equal(t, "security", parsedQuery.Controller)
				assert.Equal(t, "getProfile", parsedQuery.Action)
				assert.Equal(t, id, parsedQuery.Id)

				res := profile.Profile{Id: id, Source: []byte(`{"policies":[{"roleId":"admin"},{"roleId":"other"}]}`)}
				r, _ := json.Marshal(res)
				return types.KuzzleResponse{Result: r}
			}
			if callCount == 1 {
				callCount++
				assert.Equal(t, "security", parsedQuery.Controller)
				assert.Equal(t, "createOrReplaceProfile", parsedQuery.Action)
				assert.Equal(t, "replace", options.GetIfExist())
				assert.Equal(t, id, parsedQuery.Id)

				r, _ := json.Marshal(expectedNewProfile)
				return types.KuzzleResponse{Result: r}
			}

			return types.KuzzleResponse{Result: nil}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	p, _ := security.NewSecurity(k).Profile.Fetch(id, nil)

	newPolicies := []types.Policy{
		{RoleId: "newRoleId", AllowInternalIndex: true},
		{RoleId: "otherRoleId", RestrictedTo: []types.PolicyRestriction{{Index: "index", Collections: []string{"foo", "bar"}}}},
	}

	p.SetContent([]byte(`{"im":"emptyInside"}`))
	for _, policy := range newPolicies {
		p.AddPolicy(policy)
	}
	newProfile, _ := p.Save(nil)

	assert.Equal(t, expectedNewProfile.Id, newProfile.Id)
	assert.Equal(t, expectedNewProfile.Source, newProfile.Source)
	assert.Equal(t, expectedNewProfile.GetPolicies(), newProfile.GetPolicies())
}

func ExampleProfile_Save() {
	id := "profileId"
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)

	p, _ := security.NewSecurity(k).Profile.Fetch(id, nil)

	newPolicies := []types.Policy{
		{RoleId: "newRoleId", AllowInternalIndex: true},
		{RoleId: "otherRoleId", RestrictedTo: []types.PolicyRestriction{{Index: "index", Collections: []string{"foo", "bar"}}}},
	}

	p.SetContent([]byte(`{"im":"emptyInside"}`))
	for _, policy := range newPolicies {
		p.AddPolicy(policy)
	}
	res, err := p.Save(nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res.Id, res.GetPolicies())
}

func TestProfileUpdate(t *testing.T) {
	id := "profileId"
	expectedUpdatedProfile := profile.Profile{Id: id, Source: []byte(`{"you":"completeMe","policies":[{"roleId":"boringNewRoleId"}]}`)}
	callCount := 0

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			if callCount == 0 {
				callCount++
				assert.Equal(t, "security", parsedQuery.Controller)
				assert.Equal(t, "getProfile", parsedQuery.Action)
				assert.Equal(t, id, parsedQuery.Id)

				res := profile.Profile{Id: id, Source: []byte(`{"im":"emptyInside","policies":[{"roleId":"admin"},{"roleId":"other"}]}`)}
				r, _ := json.Marshal(res)
				return types.KuzzleResponse{Result: r}
			}
			if callCount == 1 {
				callCount++
				assert.Equal(t, "security", parsedQuery.Controller)
				assert.Equal(t, "updateProfile", parsedQuery.Action)
				assert.Equal(t, id, parsedQuery.Id)

				r, _ := json.Marshal(expectedUpdatedProfile)
				return types.KuzzleResponse{Result: r}
			}

			return types.KuzzleResponse{Result: nil}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	p, _ := security.NewSecurity(k).Profile.Fetch(id, nil)

	newPolicies := []types.Policy{
		{RoleId: "boringNewRoleId"},
	}

	p.SetContent([]byte(`{"you":"completeMe"}`))
	updatedProfile, _ := p.Update(newPolicies, nil)

	assert.Equal(t, expectedUpdatedProfile.Id, updatedProfile.Id)
	assert.Equal(t, expectedUpdatedProfile.Source, updatedProfile.Source)
	assert.Equal(t, newPolicies, updatedProfile.GetPolicies())
}

func ExampleProfile_Update() {
	id := "profileId"
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)

	p, _ := security.NewSecurity(k).Profile.Fetch(id, nil)

	newPolicies := []types.Policy{
		{RoleId: "boringNewRoleId"},
	}

	p.SetContent([]byte(`{"you":"completeMe"}`))
	res, err := p.Update(newPolicies, nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res.Id, res.GetPolicies())
}

func TestProfileDelete(t *testing.T) {
	id := "SomeMenJustWantToWatchTheWorldBurn"
	callCount := 0

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			if callCount == 0 {
				callCount++
				assert.Equal(t, "security", parsedQuery.Controller)
				assert.Equal(t, "getProfile", parsedQuery.Action)
				assert.Equal(t, id, parsedQuery.Id)

				res := profile.Profile{Id: id, Source: []byte(`{"policies":[{"roleId":"admin"},{"roleId":"other"}]}`)}
				r, _ := json.Marshal(res)
				return types.KuzzleResponse{Result: r}
			}
			if callCount == 1 {
				callCount++
				assert.Equal(t, "security", parsedQuery.Controller)
				assert.Equal(t, "deleteProfile", parsedQuery.Action)
				assert.Equal(t, id, parsedQuery.Id)

				res := types.ShardResponse{Id: id}
				r, _ := json.Marshal(res)
				return types.KuzzleResponse{Result: r}
			}

			return types.KuzzleResponse{Result: nil}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	p, _ := security.NewSecurity(k).Profile.Fetch(id, nil)

	inTheEnd, _ := p.Delete(nil)

	assert.Equal(t, id, inTheEnd)
}

func ExampleProfile_Delete() {
	id := "SomeMenJustWantToWatchTheWorldBurn"
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	p, _ := security.NewSecurity(k).Profile.Fetch(id, nil)
	res, err := p.Delete(nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res)
}

func TestFetchEmptyId(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Security.Profile.Fetch: profile id required"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := security.NewSecurity(k).Profile.Fetch("", nil)
	assert.NotNil(t, err)
}

func TestFetchError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := security.NewSecurity(k).Profile.Fetch("profileId", nil)
	assert.NotNil(t, err)
}

func TestFetch(t *testing.T) {
	id := "profileId"

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "getProfile", parsedQuery.Action)
			assert.Equal(t, id, parsedQuery.Id)

			res := profile.Profile{Id: id, Source: []byte(`{"policies":[{"roleId":"admin"},{"roleId":"other"}]}`)}
			r, _ := json.Marshal(res)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := security.NewSecurity(k).Profile.Fetch(id, nil)

	assert.Equal(t, id, res.Id)
	assert.Equal(t, []types.Policy{{RoleId: "admin"}, {RoleId: "other"}}, res.GetPolicies())
}

func ExampleSecurityProfile_Fetch() {
	id := "profileId"
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, err := security.NewSecurity(k).Profile.Fetch(id, nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res.Id, res.GetPolicies())
}

func TestSearchError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := security.NewSecurity(k).Profile.Search(nil, nil)
	assert.NotNil(t, err)
}

func TestSearch(t *testing.T) {
	hits := make([]profile.Profile, 1)
	hits[0] = profile.Profile{Id: "profile42", Source: json.RawMessage(`{"policies":[{"roleId":"admin"}]}`)}
	var results = profile.ProfileSearchResult{Total: 42, Hits: hits}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "searchProfiles", parsedQuery.Action)

			res := profile.ProfileSearchResult{Total: results.Total, Hits: results.Hits}
			r, _ := json.Marshal(res)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := security.NewSecurity(k).Profile.Search(nil, nil)
	assert.Equal(t, results.Total, res.Total)
	assert.Equal(t, hits, res.Hits)
	assert.Equal(t, "profile42", res.Hits[0].Id)
	assert.Equal(t, json.RawMessage(`{"policies":[{"roleId":"admin"}]}`), res.Hits[0].Source)
	assert.Equal(t, "admin", res.Hits[0].GetPolicies()[0].RoleId)
}

func ExampleSecurityProfile_Search() {
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, err := security.NewSecurity(k).Profile.Search(nil, nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res.Hits[0].Id, res.Hits[0].GetPolicies())
}

func TestSearchWithScroll(t *testing.T) {
	hits := make([]profile.Profile, 1)
	hits[0] = profile.Profile{Id: "profile42", Source: json.RawMessage(`{"policies":[{"roleId":"admin"}]}`)}
	var results = profile.ProfileSearchResult{Total: 42, Hits: hits}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "searchProfiles", parsedQuery.Action)

			res := profile.ProfileSearchResult{Total: results.Total, Hits: results.Hits, ScrollId: "f00b4r"}
			r, _ := json.Marshal(res)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	opts := types.NewQueryOptions()
	opts.SetFrom(2)
	opts.SetSize(4)
	opts.SetScroll("1m")

	res, _ := security.NewSecurity(k).Profile.Search(nil, opts)
	assert.Equal(t, results.Total, res.Total)
	assert.Equal(t, hits, res.Hits)
	assert.Equal(t, "f00b4r", res.ScrollId)
	assert.Equal(t, "profile42", res.Hits[0].Id)
	assert.Equal(t, json.RawMessage(`{"policies":[{"roleId":"admin"}]}`), res.Hits[0].Source)
	assert.Equal(t, "admin", res.Hits[0].GetPolicies()[0].RoleId)
}

func TestScrollEmptyScrollId(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Security.Profile.Scroll: scroll id required"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := security.NewSecurity(k).Profile.Scroll("", nil)
	assert.NotNil(t, err)
}

func TestScrollError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := security.NewSecurity(k).Profile.Scroll("f00b4r", nil)
	assert.NotNil(t, err)
}

func TestScroll(t *testing.T) {
	type response struct {
		Total int               `json:"total"`
		Hits  []profile.Profile `json:"hits"`
	}

	hits := make([]profile.Profile, 1)
	hits[0] = profile.Profile{Id: "profile42", Source: json.RawMessage(`{"policies":[{"roleId":"admin"}]}`)}
	var results = profile.ProfileSearchResult{Total: 42, Hits: hits}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "scrollProfiles", parsedQuery.Action)

			res := response{Total: results.Total, Hits: results.Hits}
			r, _ := json.Marshal(res)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := security.NewSecurity(k).Profile.Scroll("f00b4r", nil)
	assert.Equal(t, results.Total, res.Total)
	assert.Equal(t, hits, res.Hits)
	assert.Equal(t, "profile42", res.Hits[0].Id)
	assert.Equal(t, json.RawMessage(`{"policies":[{"roleId":"admin"}]}`), res.Hits[0].Source)
	assert.Equal(t, "admin", res.Hits[0].GetPolicies()[0].RoleId)
}

func ExampleSecurityProfile_Scroll() {
	type response struct {
		Total int               `json:"total"`
		Hits  []profile.Profile `json:"hits"`
	}
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, err := security.NewSecurity(k).Profile.Scroll("f00b4r", nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res.Hits[0].Id, res.Hits[0].GetPolicies())
}

func TestCreateEmptyId(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Security.Profile.Create: profile id required"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := security.NewSecurity(k).Profile.Create("", types.Policies{}, nil)
	assert.NotNil(t, err)
}

func TestCreateError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := security.NewSecurity(k).Profile.Create("profileId", types.Policies{}, nil)
	assert.NotNil(t, err)
}

func TestCreate(t *testing.T) {
	id := "profileId"

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "createProfile", parsedQuery.Action)
			assert.Equal(t, id, parsedQuery.Id)

			res := profile.Profile{
				Id:     id,
				Source: []byte(`{"policies":[{"roleId":"admin"},{"roleId":"other"}]}`),
			}
			r, _ := json.Marshal(res)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	policies := []types.Policy{}
	policies = append(policies, types.Policy{RoleId: "admin"})
	policies = append(policies, types.Policy{RoleId: "other"})
	res, _ := security.NewSecurity(k).Profile.Create(id, types.Policies{Policies: policies}, nil)

	assert.Equal(t, id, res.Id)
	assert.Equal(t, []types.Policy{{RoleId: "admin"}, {RoleId: "other"}}, res.GetPolicies())
}

func ExampleSecurityProfile_Create() {
	id := "profileId"
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)

	policies := []types.Policy{}
	policies = append(policies, types.Policy{RoleId: "admin"})
	policies = append(policies, types.Policy{RoleId: "other"})
	res, err := security.NewSecurity(k).Profile.Create(id, types.Policies{Policies: policies}, nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res.Id, res.GetPolicies())
}

func TestCreateIfExists(t *testing.T) {
	id := "profileId"

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "createOrReplaceProfile", parsedQuery.Action)
			assert.Equal(t, id, parsedQuery.Id)

			res := profile.Profile{
				Id:     id,
				Source: []byte(`{"policies":[{"roleId":"admin"},{"roleId":"other"}]}`),
			}
			r, _ := json.Marshal(res)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	policies := []types.Policy{}
	policies = append(policies, types.Policy{RoleId: "admin"})
	policies = append(policies, types.Policy{RoleId: "other"})

	opts := types.NewQueryOptions()
	opts.SetIfExist("replace")

	res, _ := security.NewSecurity(k).Profile.Create(id, types.Policies{Policies: policies}, opts)

	assert.Equal(t, id, res.Id)
	assert.Equal(t, []types.Policy{{RoleId: "admin"}, {RoleId: "other"}}, res.GetPolicies())
}

func TestCreateWithStrictOption(t *testing.T) {
	id := "profileId"

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "createProfile", parsedQuery.Action)
			assert.Equal(t, id, parsedQuery.Id)

			res := profile.Profile{
				Id:     id,
				Source: []byte(`{"policies":[{"roleId":"admin"},{"roleId":"other"}]}`),
			}
			r, _ := json.Marshal(res)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	policies := []types.Policy{}
	policies = append(policies, types.Policy{RoleId: "admin"})
	policies = append(policies, types.Policy{RoleId: "other"})

	opts := types.NewQueryOptions()
	opts.SetIfExist("error")

	res, _ := security.NewSecurity(k).Profile.Create(id, types.Policies{Policies: policies}, opts)

	assert.Equal(t, id, res.Id)
	assert.Equal(t, []types.Policy{{RoleId: "admin"}, {RoleId: "other"}}, res.GetPolicies())
}

func TestCreateWithWrongOption(t *testing.T) {
	id := "profileId"

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	policies := []types.Policy{}
	policies = append(policies, types.Policy{RoleId: "admin"})
	policies = append(policies, types.Policy{RoleId: "other"})

	opts := types.NewQueryOptions()
	opts.SetIfExist("unknown")

	_, err := security.NewSecurity(k).Profile.Create(id, types.Policies{Policies: policies}, opts)

	assert.Equal(t, "Invalid value for the 'ifExist' option: 'unknown'", fmt.Sprint(err))
}

func TestUpdateEmptyId(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Security.Profile.Update: profile id required"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := security.NewSecurity(k).Profile.Update("", types.Policies{}, nil)
	assert.NotNil(t, err)
}

func TestUpdateError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := security.NewSecurity(k).Profile.Update("profileId", types.Policies{}, nil)
	assert.NotNil(t, err)
}

func TestUpdate(t *testing.T) {
	id := "profileId"

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "updateProfile", parsedQuery.Action)
			assert.Equal(t, id, parsedQuery.Id)

			res := profile.Profile{
				Id:     id,
				Source: []byte(`{"policies":[{"roleId":"admin"},{"roleId":"other"}]}`),
			}
			r, _ := json.Marshal(res)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	policies := []types.Policy{}
	policies = append(policies, types.Policy{RoleId: "admin"})
	policies = append(policies, types.Policy{RoleId: "other"})
	res, _ := security.NewSecurity(k).Profile.Update(id, types.Policies{Policies: policies}, nil)

	assert.Equal(t, id, res.Id)
	assert.Equal(t, []types.Policy{{RoleId: "admin"}, {RoleId: "other"}}, res.GetPolicies())
}

func ExampleSecurityProfile_Update() {
	id := "profileId"
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)

	policies := []types.Policy{}
	policies = append(policies, types.Policy{RoleId: "admin"})
	policies = append(policies, types.Policy{RoleId: "other"})
	res, err := security.NewSecurity(k).Profile.Update(id, types.Policies{Policies: policies}, nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res.Id, res.GetPolicies())
}

func TestDeleteEmptyId(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Security.Profile.Delete: profile id required"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := security.NewSecurity(k).Profile.Delete("", nil)
	assert.NotNil(t, err)
}

func TestDeleteError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := security.NewSecurity(k).Profile.Delete("profileId", nil)
	assert.NotNil(t, err)
}

func TestDelete(t *testing.T) {
	id := "profileId"

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "deleteProfile", parsedQuery.Action)
			assert.Equal(t, id, parsedQuery.Id)

			res := types.ShardResponse{Id: id}
			r, _ := json.Marshal(res)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := security.NewSecurity(k).Profile.Delete(id, nil)

	assert.Equal(t, id, res)
}

func ExampleSecurityProfile_Delete() {
	id := "profileId"
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, err := security.NewSecurity(k).Profile.Delete(id, nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res)
}
