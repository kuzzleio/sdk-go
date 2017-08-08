package user_test

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/security"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFetchEmptyKuid(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Security.User.Fetch: user id required"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := security.NewSecurity(k).User.Fetch("", nil)
	assert.NotNil(t, err)
}

func TestFetchError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := security.NewSecurity(k).User.Fetch("userId", nil)
	assert.NotNil(t, err)
}

func TestFetch(t *testing.T) {
	id := "userId"

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "getUser", parsedQuery.Action)
			assert.Equal(t, id, parsedQuery.Id)

			res := types.User{Id: id, Source: []byte(`{"profileIds":["admin","other"],"name":"Luke","function":"Jedi"}`)}
			r, _ := json.Marshal(res)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := security.NewSecurity(k).User.Fetch(id, nil)

	assert.Equal(t, id, res.Id)

	assert.Equal(t, []string{"admin", "other"}, res.ProfileIDs())

	assert.Equal(t, "Luke", res.Content("name"))
	assert.Equal(t, "Jedi", res.Content("function"))

	contentAsMap := make(map[string]interface{})
	contentAsMap["name"] = "Luke"
	contentAsMap["function"] = "Jedi"

	assert.Equal(t, contentAsMap, res.ContentMap("name", "function"))
}

func TestSearchError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := security.NewSecurity(k).User.Search(nil, nil)
	assert.NotNil(t, err)
}

func TestSearch(t *testing.T) {
	hits := make([]types.User, 1)
	hits[0] = types.User{Id: "user42", Source: json.RawMessage(`{"profileIds":["admin","other"],"foo":"bar"}`)}
	var results = types.KuzzleSearchUsersResult{Total: 42, Hits: hits}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "searchUsers", parsedQuery.Action)

			res := types.KuzzleSearchUsersResult{Total: results.Total, Hits: results.Hits}
			r, _ := json.Marshal(res)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := security.NewSecurity(k).User.Search(nil, nil)
	assert.Equal(t, results.Total, res.Total)
	assert.Equal(t, hits, res.Hits)
	assert.Equal(t, res.Hits[0].Id, "user42")
	assert.Equal(t, res.Hits[0].Source, json.RawMessage(`{"profileIds":["admin","other"],"foo":"bar"}`))
	assert.Equal(t, res.Hits[0].ProfileIDs(), []string{"admin", "other"})
	assert.Equal(t, res.Hits[0].Content("foo"), "bar")
}

func TestSearchWithScroll(t *testing.T) {
	hits := make([]types.User, 1)
	hits[0] = types.User{Id: "user42", Source: json.RawMessage(`{"profileIds":["admin","other"],"foo":"bar"}`)}
	var results = types.KuzzleSearchUsersResult{Total: 42, Hits: hits}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
 			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "searchUsers", parsedQuery.Action)

			res := types.KuzzleSearchUsersResult{Total: results.Total, Hits: results.Hits, ScrollId: "f00b4r"}
			r, _ := json.Marshal(res)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	opts := types.NewQueryOptions()
	opts.SetFrom(2)
	opts.SetSize(4)
	opts.SetScroll("1m")

	res, _ := security.NewSecurity(k).User.Search(nil, opts)
	assert.Equal(t, results.Total, res.Total)
	assert.Equal(t, hits, res.Hits)
	assert.Equal(t, "f00b4r", res.ScrollId)
	assert.Equal(t, res.Hits[0].Id, "user42")
	assert.Equal(t, res.Hits[0].Source, json.RawMessage(`{"profileIds":["admin","other"],"foo":"bar"}`))
	assert.Equal(t, res.Hits[0].ProfileIDs(), []string{"admin", "other"})
	assert.Equal(t, res.Hits[0].Content("foo"), "bar")
}

func TestScrollEmptyScrollId(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Security.User.Scroll: scroll id required"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := security.NewSecurity(k).User.Scroll("", nil)
	assert.NotNil(t, err)
}

func TestScrollError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := security.NewSecurity(k).User.Scroll("f00b4r", nil)
	assert.NotNil(t, err)
}

func TestScroll(t *testing.T) {
	type response struct {
		Total int          `json:"total"`
		Hits  []types.User `json:"hits"`
	}

	hits := make([]types.User, 1)
	hits[0] = types.User{Id: "user42", Source: json.RawMessage(`{"profileIds":["admin","other"],"foo":"bar"}`)}
	var results = types.KuzzleSearchUsersResult{Total: 42, Hits: hits}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "scrollUsers", parsedQuery.Action)

			res := response{Total: results.Total, Hits: results.Hits}
			r, _ := json.Marshal(res)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := security.NewSecurity(k).User.Scroll("f00b4r", nil)
	assert.Equal(t, results.Total, res.Total)
	assert.Equal(t, hits, res.Hits)
	assert.Equal(t, res.Hits[0].Id, "user42")
	assert.Equal(t, res.Hits[0].Source, json.RawMessage(`{"profileIds":["admin","other"],"foo":"bar"}`))
	assert.Equal(t, res.Hits[0].ProfileIDs(), []string{"admin", "other"})
	assert.Equal(t, res.Hits[0].Content("foo"), "bar")
}

func TestCreateEmptyId(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Security.User.Create: user id required"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := security.NewSecurity(k).User.Create("", types.UserData{}, nil)
	assert.NotNil(t, err)
}

func TestCreateError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := security.NewSecurity(k).User.Create("userId", types.UserData{}, nil)
	assert.NotNil(t, err)
}

func TestCreate(t *testing.T) {
	id := "userId"

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "createUser", parsedQuery.Action)
			assert.Equal(t, id, parsedQuery.Id)

			res := types.User{
				Id:     id,
				Source: []byte(`{"profileIds":["admin","other"],"name":"Luke","function":"Jedi"}`),
			}
			r, _ := json.Marshal(res)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	type UserContent map[string]interface{}
	cred := types.UserCredentials{}
	json.Unmarshal([]byte(`{"local": {"Username": "username", "Password": "password"}}`), cred)
	ud := types.UserData{ProfileIds: []string{"default", "anonymous"}, Content: UserContent{"foo": "bar"}, Credentials: cred}

	res, _ := security.NewSecurity(k).User.Create(id, ud, nil)

	assert.Equal(t, id, res.Id)

	assert.Equal(t, []string{"admin", "other"}, res.ProfileIDs())

	assert.Equal(t, "Luke", res.Content("name"))
	assert.Equal(t, "Jedi", res.Content("function"))

	contentAsMap := make(map[string]interface{})
	contentAsMap["name"] = "Luke"
	contentAsMap["function"] = "Jedi"

	assert.Equal(t, contentAsMap, res.ContentMap("name", "function"))
}

func TestCreateRestrictedEmptyId(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Security.User.CreateRestrictedUser: user id required"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := security.NewSecurity(k).User.CreateRestrictedUser("", types.UserData{}, nil)
	assert.NotNil(t, err)
}

func TestCreateRestrictedError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := security.NewSecurity(k).User.CreateRestrictedUser("userId", types.UserData{}, nil)
	assert.NotNil(t, err)
}

func TestCreateRestricted(t *testing.T) {
	id := "userId"

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "createRestrictedUser", parsedQuery.Action)
			assert.Equal(t, id, parsedQuery.Id)

			res := types.User{
				Id:     id,
				Source: []byte(`{"profileIds":["admin","other"],"name":"Luke","function":"Jedi"}`),
			}
			r, _ := json.Marshal(res)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	type UserContent map[string]interface{}
	cred := types.UserCredentials{}
	json.Unmarshal([]byte(`{"local": {"Username": "username", "Password": "password"}}`), cred)
	ud := types.UserData{Content: UserContent{"foo": "bar"}, Credentials: cred}

	res, _ := security.NewSecurity(k).User.CreateRestrictedUser(id, ud, nil)

	assert.Equal(t, id, res.Id)

	assert.Equal(t, []string{"admin", "other"}, res.ProfileIDs())

	assert.Equal(t, "Luke", res.Content("name"))
	assert.Equal(t, "Jedi", res.Content("function"))

	contentAsMap := make(map[string]interface{})
	contentAsMap["name"] = "Luke"
	contentAsMap["function"] = "Jedi"

	assert.Equal(t, contentAsMap, res.ContentMap("name", "function"))
}

func TestReplaceEmptyId(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Security.User.Replace: user id required"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := security.NewSecurity(k).User.Replace("", types.UserData{}, nil)
	assert.NotNil(t, err)
}

func TestReplaceError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := security.NewSecurity(k).User.Replace("userId", types.UserData{}, nil)
	assert.NotNil(t, err)
}

func TestReplace(t *testing.T) {
	id := "userId"

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "replaceUser", parsedQuery.Action)
			assert.Equal(t, id, parsedQuery.Id)

			res := types.User{
				Id:     id,
				Source: []byte(`{"profileIds":["admin","other"],"name":"Luke","function":"Jedi"}`),
			}
			r, _ := json.Marshal(res)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	type UserContent map[string]interface{}
	cred := types.UserCredentials{}
	json.Unmarshal([]byte(`{"local": {"Username": "username", "Password": "password"}}`), cred)
	ud := types.UserData{ProfileIds: []string{"default", "anonymous"}, Content: UserContent{"foo": "bar"}, Credentials: cred}

	res, _ := security.NewSecurity(k).User.Replace(id, ud, nil)

	assert.Equal(t, id, res.Id)

	assert.Equal(t, []string{"admin", "other"}, res.ProfileIDs())

	assert.Equal(t, "Luke", res.Content("name"))
	assert.Equal(t, "Jedi", res.Content("function"))

	contentAsMap := make(map[string]interface{})
	contentAsMap["name"] = "Luke"
	contentAsMap["function"] = "Jedi"

	assert.Equal(t, contentAsMap, res.ContentMap("name", "function"))
}

func TestUpdateEmptyId(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Security.User.Update: user id required"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := security.NewSecurity(k).User.Update("", types.UserData{}, nil)
	assert.NotNil(t, err)
}

func TestUpdateError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := security.NewSecurity(k).User.Update("userId", types.UserData{}, nil)
	assert.NotNil(t, err)
}

func TestUpdate(t *testing.T) {
	id := "userId"

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "updateUser", parsedQuery.Action)
			assert.Equal(t, id, parsedQuery.Id)

			res := types.User{
				Id:     id,
				Source: []byte(`{"profileIds":["admin","other"],"name":"Luke","function":"Jedi"}`),
			}
			r, _ := json.Marshal(res)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	type UserContent map[string]interface{}
	cred := types.UserCredentials{}
	json.Unmarshal([]byte(`{"local": {"Username": "username", "Password": "password"}}`), cred)
	ud := types.UserData{ProfileIds: []string{"default", "anonymous"}, Content: UserContent{"foo": "bar"}, Credentials: cred}

	res, _ := security.NewSecurity(k).User.Update(id, ud, nil)

	assert.Equal(t, id, res.Id)

	assert.Equal(t, []string{"admin", "other"}, res.ProfileIDs())

	assert.Equal(t, "Luke", res.Content("name"))
	assert.Equal(t, "Jedi", res.Content("function"))

	contentAsMap := make(map[string]interface{})
	contentAsMap["name"] = "Luke"
	contentAsMap["function"] = "Jedi"

	assert.Equal(t, contentAsMap, res.ContentMap("name", "function"))
}

func TestDeleteEmptyId(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Security.User.Delete: user id required"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := security.NewSecurity(k).User.Delete("", nil)
	assert.NotNil(t, err)
}

func TestDeleteError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := security.NewSecurity(k).User.Delete("userId", nil)
	assert.NotNil(t, err)
}

func TestDelete(t *testing.T) {
	id := "userId"

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "deleteUser", parsedQuery.Action)
			assert.Equal(t, id, parsedQuery.Id)

			res := types.ShardResponse{Id: id, }
			r, _ := json.Marshal(res)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := security.NewSecurity(k).User.Delete(id, nil)

	assert.Equal(t, id, res)
}

func TestGetRightsEmptyId(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Security.User.GetRights: user id required"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := security.NewSecurity(k).User.GetRights("", nil)
	assert.NotNil(t, err)
}

func TestGetRightsError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := security.NewSecurity(k).User.GetRights("userId", nil)
	assert.NotNil(t, err)
}

func TestGetRights(t *testing.T) {
	id := "userId"

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "getUserRights", parsedQuery.Action)
			assert.Equal(t, id, parsedQuery.Id)

			type resultUserRights struct {
				UserRights []types.UserRights `json:"hits"`
			}
			userRights := []types.UserRights{}
			userRights = append(userRights, types.UserRights{Controller: "wow-controller", Action: "such-action", Index: "much indexes", Collection: "very collection", Value: "wow"})
			actualRights := resultUserRights{UserRights: userRights}
			r, _ := json.Marshal(actualRights)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := security.NewSecurity(k).User.GetRights(id, nil)

	expectedRights := []types.UserRights{}
	expectedRights = append(expectedRights, types.UserRights{Controller: "wow-controller", Action: "such-action", Index: "much indexes", Collection: "very collection", Value: "wow"})

	assert.Equal(t, expectedRights, res)
}

func TestIsActionAllowedNilRights(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)

	_, err := security.NewSecurity(k).User.IsActionAllowed(nil, "wow-controller", "such-action", "", "")
	assert.NotNil(t, err)
}

func TestIsActionAllowedEmptyController(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)

	_, err := security.NewSecurity(k).User.IsActionAllowed([]types.UserRights{}, "", "such-action", "", "")
	assert.NotNil(t, err)
}

func TestIsActionAllowedEmptyAction(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)

	_, err := security.NewSecurity(k).User.IsActionAllowed([]types.UserRights{}, "wow-controller", "", "", "")
	assert.NotNil(t, err)
}

func TestIsActionAllowedEmptyRights(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)

	res, _ := security.NewSecurity(k).User.IsActionAllowed([]types.UserRights{}, "wow-controller", "such-action", "much-index", "very-collection")

	assert.Equal(t, "denied", res)
}

func TestIsActionAllowedResultAllowed(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)

	userRights := []types.UserRights{}
	userRights = append(userRights, types.UserRights{Controller: "wow-controller", Action: "*", Index: "much-index", Collection: "very-collection", Value: "allowed"})

	res, _ := security.NewSecurity(k).User.IsActionAllowed(userRights, "wow-controller", "such-action", "much-index", "very-collection")

	assert.Equal(t, "allowed", res)
}

func TestIsActionAllowedResultConditional(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)

	userRights := []types.UserRights{}
	userRights = append(userRights, types.UserRights{Controller: "wow-controller", Action: "*", Index: "much-index", Collection: "very-collection", Value: "conditional"})

	res, _ := security.NewSecurity(k).User.IsActionAllowed(userRights, "wow-controller", "action", "much-index", "very-collection")

	assert.Equal(t, "conditional", res)
}

func TestIsActionAllowedResultDenied(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)

	userRights := []types.UserRights{}
	userRights = append(userRights, types.UserRights{Controller: "wow-controller.", Action: "action-such", Index: "much-index", Collection: "very-collection", Value: "allowed"})

	res, _ := security.NewSecurity(k).User.IsActionAllowed(userRights, "wow-controller", "action", "much-index", "very-collection")

	assert.Equal(t, "denied", res)
}
