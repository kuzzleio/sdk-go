package profile_test

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/security"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
	"fmt"
)

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

			res := types.Profile{Id: id, Source: []byte(`{"policies":[{"roleId":"admin"},{"roleId":"other"}]}`)}
			r, _ := json.Marshal(res)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := security.NewSecurity(k).Profile.Fetch(id, nil)

	assert.Equal(t, id, res.Id)
	assert.Equal(t, []string{"admin", "other"}, res.Policies())
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
	hits := make([]types.Profile, 1)
	hits[0] = types.Profile{Id: "profile42", Source: json.RawMessage(`{"policies":[{"roleId":"admin"}]}`)}
	var results = types.KuzzleSearchProfilesResult{Total: 42, Hits: hits}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "searchProfiles", parsedQuery.Action)

			res := types.KuzzleSearchProfilesResult{Total: results.Total, Hits: results.Hits}
			r, _ := json.Marshal(res)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := security.NewSecurity(k).Profile.Search(nil, nil)
	assert.Equal(t, results.Total, res.Total)
	assert.Equal(t, hits, res.Hits)
	assert.Equal(t, res.Hits[0].Id, "profile42")
	assert.Equal(t, res.Hits[0].Source, json.RawMessage(`{"policies":[{"roleId":"admin"}]}`))
	assert.Equal(t, res.Hits[0].Policies()[0], "admin")
}

func TestSearchWithScroll(t *testing.T) {
	hits := make([]types.Profile, 1)
	hits[0] = types.Profile{Id: "profile42", Source: json.RawMessage(`{"policies":[{"roleId":"admin"}]}`)}
	var results = types.KuzzleSearchProfilesResult{Total: 42, Hits: hits}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "searchProfiles", parsedQuery.Action)

			res := types.KuzzleSearchProfilesResult{Total: results.Total, Hits: results.Hits, ScrollId: "f00b4r"}
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
	assert.Equal(t, res.Hits[0].Id, "profile42")
	assert.Equal(t, res.Hits[0].Source, json.RawMessage(`{"policies":[{"roleId":"admin"}]}`))
	assert.Equal(t, res.Hits[0].Policies()[0], "admin")
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
		Total int             `json:"total"`
		Hits  []types.Profile `json:"hits"`
	}

	hits := make([]types.Profile, 1)
	hits[0] = types.Profile{Id: "profile42", Source: json.RawMessage(`{"policies":[{"roleId":"admin"}]}`)}
	var results = types.KuzzleSearchProfilesResult{Total: 42, Hits: hits}

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
	assert.Equal(t, res.Hits[0].Id, "profile42")
	assert.Equal(t, res.Hits[0].Source, json.RawMessage(`{"policies":[{"roleId":"admin"}]}`))
	assert.Equal(t, res.Hits[0].Policies()[0], "admin")
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

			res := types.Profile{
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
	assert.Equal(t, []string{"admin", "other"}, res.Policies())
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

			res := types.Profile{
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
	assert.Equal(t, []string{"admin", "other"}, res.Policies())
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

			res := types.Profile{
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
	assert.Equal(t, []string{"admin", "other"}, res.Policies())
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

			res := types.Profile{
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
	assert.Equal(t, []string{"admin", "other"}, res.Policies())
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

			res := types.ShardResponse{Id: id, }
			r, _ := json.Marshal(res)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := security.NewSecurity(k).Profile.Delete(id, nil)

	assert.Equal(t, id, res)
}

