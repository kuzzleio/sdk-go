package security_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/kuzzleio/sdk-go/connection/websocket"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/security"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"os/user"
)

func TestUserAddProfile(t *testing.T) {
	id := "userId"

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "getUser", parsedQuery.Action)
			assert.Equal(t, id, parsedQuery.Id)

			return &types.KuzzleResponse{Result: []byte(`{
				"_id": "userId",
				"_source": {
					"profileIds": ["admin", "other"],
					"name": "Luke",
					"function": "Jedi"
				}
			}`)}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	u, _ := k.Security.FetchUser(id, nil)

	u.AddProfile(&security.Profile{Id: "adminNew"})

	assert.Equal(t, []string{"admin", "other", "adminNew"}, u.ProfileIds)
}

func ExampleUser_AddProfile() {
	id := "userId"
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	u, _ := k.Security.FetchUser(id, nil)

	res := u.AddProfile(&security.Profile{Id: "adminNew"})

	fmt.Println(res.Id, res.ProfileIds)
}

func TestUserGetProfilesEmptyProfileIds(t *testing.T) {
	id := "userId"

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "getUser", parsedQuery.Action)
			assert.Equal(t, id, parsedQuery.Id)

			return &types.KuzzleResponse{Result: []byte(`{
				"_id": "userId",
				"_source": {
					"profileIds": [],
					"name": "Luke",
					"function": "Jedi"
				}
			}`)}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	u, _ := k.Security.FetchUser(id, nil)

	profiles, _ := u.GetProfiles(nil)

	assert.Equal(t, []*security.Profile{}, profiles)
}

func TestUserGetProfilesError(t *testing.T) {
	id := "userId"
	callCount := 0

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			if callCount == 0 {
				callCount++
				assert.Equal(t, "security", parsedQuery.Controller)
				assert.Equal(t, "getUser", parsedQuery.Action)
				assert.Equal(t, id, parsedQuery.Id)

				return &types.KuzzleResponse{Result: []byte(`{
					"_id": "userId",
					"_source": {
						"profileIds": ["admin", "other"],
						"name": "Luke",
						"function": "Jedi"
					}
				}`)}
			}

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "getProfile", parsedQuery.Action)
			assert.Equal(t, "admin", parsedQuery.Id)

			return &types.KuzzleResponse{Error: &types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	u, _ := k.Security.FetchUser(id, nil)

	_, err := u.GetProfiles(nil)

	assert.Equal(t, "Unit test error", fmt.Sprint(err))
}

func TestUserGetProfiles(t *testing.T) {
	id := "userId"
	callCount := 0

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			if callCount == 0 {
				callCount++
				assert.Equal(t, "security", parsedQuery.Controller)
				assert.Equal(t, "getUser", parsedQuery.Action)
				assert.Equal(t, id, parsedQuery.Id)

				return &types.KuzzleResponse{Result: []byte(`{
					"_id": "userId",
					"_source": {
						"profileIds": ["admin", "other"],
						"name": "Luke",
						"function": "Jedi"
					}
				}`)}
			}
			if callCount == 1 {
				callCount++
				assert.Equal(t, "security", parsedQuery.Controller)
				assert.Equal(t, "getProfile", parsedQuery.Action)
				assert.Equal(t, "admin", parsedQuery.Id)

				return &types.KuzzleResponse{Result: []byte(`{
					"_id": "admin",
					"_source": {
						"policies": [
							{"roleId": "admin"},
							{"roleId": "other"}
						]
					}
				}`)}
			}
			if callCount == 2 {
				callCount++
				assert.Equal(t, "security", parsedQuery.Controller)
				assert.Equal(t, "getProfile", parsedQuery.Action)
				assert.Equal(t, "other", parsedQuery.Id)

				return &types.KuzzleResponse{Result: []byte(`{
					"_id": "other",
					"_source": {
						"policies": [
							{"roleId": "other"},
							{"roleId": "default"}
						]
					}
				}`)}
			}

			return &types.KuzzleResponse{Result: nil}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	u, _ := k.Security.FetchUser(id, nil)

	profiles, _ := u.GetProfiles(nil)

	assert.Equal(t, 2, len(profiles))
	assert.Equal(t, "admin", profiles[0].Id)
	assert.Equal(t, []*types.Policy{
		{RoleId: "admin"},
		{RoleId: "other"},
	}, profiles[0].Policies)
	assert.Equal(t, "other", profiles[1].Id)
	assert.Equal(t, []*types.Policy{
		{RoleId: "other"},
		{RoleId: "default"},
	}, profiles[1].Policies)
}

func ExampleUser_GetProfiles() {
	id := "userId"
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	u, _ := k.Security.FetchUser(id, nil)
	res, err := u.GetProfiles(nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res[0].Id, res[0].Policies)
}

func TestUserSetProfiles(t *testing.T) {
	id := "userId"

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "getUser", parsedQuery.Action)
			assert.Equal(t, id, parsedQuery.Id)

			return &types.KuzzleResponse{Result: []byte(`{
				"_id": "userId",
				"_source": {
					"profileIds": ["admin", "other"],
					"name": "Luke",
					"function": "Jedi"
				}
			}`)}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	u, _ := k.Security.FetchUser(id, nil)

	u.SetProfiles([]*security.Profile{
		{Id: "adminNew"},
		{Id: "otherNew"},
	})

	assert.Equal(t, []string{"adminNew", "otherNew"}, u.ProfileIds)
}

func ExampleUser_SetProfiles() {
	id := "userId"
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	u, _ := k.Security.FetchUser(id, nil)

	res := u.SetProfiles([]*security.Profile{
		{Id: "adminNew"},
		{Id: "otherNew"},
	})

	fmt.Println(res.ProfileIds)
}

func TestUserCreate(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "createUser", parsedQuery.Action)

			return &types.KuzzleResponse{Result: []byte(`{
				"_id": "userId",
				"_source": {
					"profileIds": ["admin", "other"],
					"name": "Luke",
					"function": "Jedi"
				}
			}`)}
		},
	}

	k, _ := kuzzle.NewKuzzle(c, nil)
	u := k.Security.NewUser()

	u.Content = map[string]interface{}{}
	u.Content["function"] = "Master Jedi"
	createdUser, _ := u.Create(nil)

	assert.Equal(t, "Jedi", createdUser.Content["function"])
}

func ExampleUser_Create() {
	id := "userId"
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	u, _ := k.Security.FetchUser(id, nil)

	if u.Content == nil {
		u.Content = make(map[string]interface{})
	}
	u.Content["function"] = "Master Jedi"
	res, err := u.Create(nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res.Id, res.Content)
}

func TestUserSaveRestricted(t *testing.T) {
	id := "userId"
	callCount := 0

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			if callCount == 0 {
				callCount++
				assert.Equal(t, "security", parsedQuery.Controller)
				assert.Equal(t, "getUser", parsedQuery.Action)
				assert.Equal(t, id, parsedQuery.Id)

				return &types.KuzzleResponse{Result: []byte(`{
					"_id": "userId",
					"_source": {
						"profileIds": ["admin", "other"],
						"name": "Luke",
						"function": "Jedi"
					}
				}`)}
			}
			if callCount == 1 {
				callCount++
				assert.Equal(t, "security", parsedQuery.Controller)
				assert.Equal(t, "createRestrictedUser", parsedQuery.Action)
				assert.Equal(t, id, parsedQuery.Id)

				return &types.KuzzleResponse{Result: []byte(`{
					"_id": "userId",
					"_source": {
						"profileIds": ["admin", "other"],
						"name": "Luke",
						"function": "Master Jedi"
					}
				}`)}
			}

			return &types.KuzzleResponse{Result: nil}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	u, _ := k.Security.FetchUser(id, nil)

	u.Content["function"] = "Master Jedi"
	createdUser, _ := u.SaveRestricted(nil)

	assert.Equal(t, "Master Jedi", createdUser.Content["function"])
}

func ExampleUser_SaveRestricted() {
	id := "userId"
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	u, _ := k.Security.FetchUser(id, nil)

	if u.Content == nil {
		u.Content = make(map[string]interface{})
	}
	u.Content["function"] = "Master Jedi"
	res, err := u.SaveRestricted(nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res.Id, res.Content)
}

func TestUserReplace(t *testing.T) {
	id := "userId"
	callCount := 0

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			if callCount == 0 {
				callCount++
				assert.Equal(t, "security", parsedQuery.Controller)
				assert.Equal(t, "getUser", parsedQuery.Action)
				assert.Equal(t, id, parsedQuery.Id)

				return &types.KuzzleResponse{Result: []byte(`{
					"_id": "userId",
					"_source": {
						"profileIds": ["admin", "other"],
						"name": "Luke",
						"function": "Jedi"
					}
				}`)}
			}
			if callCount == 1 {
				callCount++
				assert.Equal(t, "security", parsedQuery.Controller)
				assert.Equal(t, "replaceUser", parsedQuery.Action)
				assert.Equal(t, id, parsedQuery.Id)

				return &types.KuzzleResponse{Result: []byte(`{
					"_id": "userId",
					"_source": {
						"profileIds": ["admin", "other"],
						"name": "Luke",
						"function": "Master Jedi"
					}
				}`)}
			}

			return &types.KuzzleResponse{Result: nil}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	u, _ := k.Security.FetchUser(id, nil)

	u.Content["function"] = "Master Jedi"
	createdUser, _ := u.Replace(nil)

	assert.Equal(t, "Master Jedi", createdUser.Content["function"])
}

func ExampleUser_Replace() {
	id := "userId"
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	u, _ := k.Security.FetchUser(id, nil)

	if u.Content == nil {
		u.Content = make(map[string]interface{})
	}
	u.Content["function"] = "Master Jedi"
	res, err := u.Replace(nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res.Id, res.Content)
}

func TestUserUpdate(t *testing.T) {
	id := "userId"
	callCount := 0

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			if callCount == 0 {
				callCount++
				assert.Equal(t, "security", parsedQuery.Controller)
				assert.Equal(t, "getUser", parsedQuery.Action)
				assert.Equal(t, id, parsedQuery.Id)

				return &types.KuzzleResponse{Result: []byte(`{
					"_id": "userId",
					"_source": {
						"profileIds": ["admin", "other"],
						"name": "Luke",
						"function": "Jedi"
					}
				}`)}
			}
			if callCount == 1 {
				callCount++
				assert.Equal(t, "security", parsedQuery.Controller)
				assert.Equal(t, "updateUser", parsedQuery.Action)
				assert.Equal(t, id, parsedQuery.Id)
				assert.Equal(t, map[string]interface{}{
					"profileIds": []interface{}{"admin", "other", "adminNew", "otherNew"},
					"name": "Luke",
					"function": "Jedi",
					"weapon": "lightsaber",
				}, parsedQuery.Body.(map[string]interface{}))

				jsonBody, _ := json.Marshal(parsedQuery.Body)
				return &types.KuzzleResponse{Result: []byte(`{
					"_id": "` + fmt.Sprintf("%s", id) + `",
					"_source": ` + fmt.Sprintf("%s", jsonBody) + `
				}`)}
			}

			return &types.KuzzleResponse{Result: nil}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	u, _ := k.Security.FetchUser(id, nil)

	newContent := types.UserData{
		ProfileIds: []string{"adminNew", "otherNew"},
		Content: map[string]interface{}{
			"weapon": "lightsaber",
		},
	}
	updatedUser, _ := u.Update(&newContent, nil)

	assert.Equal(t, newContent.ProfileIds, updatedUser.ProfileIds)
	assert.Equal(t, map[string]interface{}{
		"name": "Luke",
		"function": "Jedi",
		"weapon": "lightsaber",
	}, updatedUser.Content)
}

func ExampleUser_Update() {
	id := "userId"
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	u, _ := k.Security.FetchUser(id, nil)

	newContent := types.UserData{ProfileIds: []string{"adminNew", "otherNew"}}
	res, err := u.Update(&newContent, nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res.Id, res.Content)
}

func TestUserDelete(t *testing.T) {
	id := "SomeMenJustWantToWatchTheWorldBurn"
	callCount := 0

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			if callCount == 0 {
				callCount++
				assert.Equal(t, "security", parsedQuery.Controller)
				assert.Equal(t, "getUser", parsedQuery.Action)
				assert.Equal(t, id, parsedQuery.Id)

				return &types.KuzzleResponse{Result: []byte(`{
					"_id": "userId",
					"_source": {
						"profileIds": ["admin", "other"],
						"name": "Luke",
						"function": "Jedi"
					}
				}`)}
			}
			if callCount == 1 {
				callCount++
				assert.Equal(t, "security", parsedQuery.Controller)
				assert.Equal(t, "deleteUser", parsedQuery.Action)
				assert.Equal(t, id, parsedQuery.Id)

				res := types.ShardResponse{Id: id}
				r, _ := json.Marshal(res)
				return &types.KuzzleResponse{Result: r}
			}

			return &types.KuzzleResponse{Result: nil}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	u, _ := k.Security.FetchUser(id, nil)

	inTheEnd, _ := u.Delete(nil)

	assert.Equal(t, id, inTheEnd)
}

func ExampleUser_Delete() {
	id := "SomeMenJustWantToWatchTheWorldBurn"
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	u, _ := k.Security.FetchUser(id, nil)

	res, err := u.Delete(nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res)
}

func TestUserContentEmptyKey(t *testing.T) {
	id := "userId"

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "getUser", parsedQuery.Action)
			assert.Equal(t, id, parsedQuery.Id)

			return &types.KuzzleResponse{Result: []byte(`{
				"_id": "userId",
				"_source": {
					"profileIds": ["admin", "other"],
					"name": "Luke",
					"function": "Jedi"
				}
			}`)}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	u, _ := k.Security.FetchUser(id, nil)

	assert.Equal(t, nil, u.Content[""])
	assert.Equal(t, "Jedi", u.Content["function"])
}

func TestUserContentNonExistingKey(t *testing.T) {
	id := "userId"

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "getUser", parsedQuery.Action)
			assert.Equal(t, id, parsedQuery.Id)

			return &types.KuzzleResponse{Result: []byte(`{
				"_id": "userId",
				"_source": {
					"profileIds": ["admin", "other"],
					"name": "Luke",
					"function": "Jedi"
				}
			}`)}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	u, _ := k.Security.FetchUser(id, nil)

	assert.Nil(t, u.Content["galaxy"])
	assert.Equal(t, "Jedi", u.Content["function"])
}

func TestUserEmptyContentMap(t *testing.T) {
	id := "userId"

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "getUser", parsedQuery.Action)
			assert.Equal(t, id, parsedQuery.Id)

			return &types.KuzzleResponse{Result: []byte(`{
				"_id": "userId",
				"_source": {
					"profileIds": ["admin", "other"],
					"name": "Luke",
					"function": "Jedi"
				}
			}`)}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	u, _ := k.Security.FetchUser(id, nil)

	assert.Equal(t, map[string]interface{}{}, u.Content)
}

func TestFetchEmptyKuid(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.MessageError{Message: "Security.User.Fetch: user id required"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	_, err := k.Security.FetchUser("", nil)

	assert.NotNil(t, err)
}

func TestFetchUserError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	_, err := k.Security.FetchUser("userId", nil)
	assert.NotNil(t, err)
}

func TestFetchUser(t *testing.T) {
	id := "userId"

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "getUser", parsedQuery.Action)
			assert.Equal(t, id, parsedQuery.Id)

			return &types.KuzzleResponse{Result: []byte(`{
				"_id": "userId",
				"_source": {
					"profileIds": ["admin", "other"],
					"name": "Luke",
					"function": "Jedi"
				}
			}`)}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	res, _ := k.Security.FetchUser(id, nil)

	assert.Equal(t, id, res.Id)

	assert.Equal(t, []string{"admin", "other"}, res.ProfileIds)

	assert.Equal(t, "Luke", res.Content["name"])
	assert.Equal(t, "Jedi", res.Content["function"])

	contentAsMap := make(map[string]interface{})
	contentAsMap["name"] = "Luke"
	contentAsMap["function"] = "Jedi"

	assert.Equal(t, contentAsMap, res.Content)
}

func ExampleSecurityUser_Fetch() {
	id := "userId"
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	res, err := k.Security.FetchUser(id, nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res.Id, res.Content)
}

func TestSearchUsersError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	_, err := k.Security.SearchUsers(nil, nil)
	assert.NotNil(t, err)
}

func TestSearchUsers(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "searchUsers", parsedQuery.Action)

			return &types.KuzzleResponse{Result: []byte(`{
				"total": 42,
				"hits": [
					{"_id": "user42", "_source": {"profileIds": ["admin", "other"], "foo": "bar"}}
				]
			}`)}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	res, _ := k.Security.SearchUsers(nil, nil)

	assert.Equal(t, 42, res.Total)
	assert.Equal(t, 1, len(res.Hits))
	assert.Equal(t, res.Hits[0].Id, "user42")
	assert.Equal(t, []string{"admin", "other"}, res.Hits[0].ProfileIds)
	assert.Equal(t, "foo", res.Hits[0].Content["foo"])
}

func ExampleSearchUsers() {
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	res, err := k.Security.SearchUsers(nil, nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res.Hits[0].Id, res.Hits[0].Content)
}

func TestSearchUsersWithScroll(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "searchUsers", parsedQuery.Action)

			return &types.KuzzleResponse{Result: []byte(`{
				"total": 42,
				"hits": [
					{"_id": "user42", "_source": {"profileIds": ["admin", "other"}, "foo": "bar"}
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

	res, _ := k.Security.SearchUsers(nil, opts)
	assert.Equal(t, 42, res.Total)
	assert.Equal(t, 1, len(res.Hits))
	assert.Equal(t, "f00b4r", res.ScrollId)
	assert.Equal(t, res.Hits[0].Id, "user42")
	assert.Equal(t, []string{"admin", "other"}, res.Hits[0].ProfileIds)
	assert.Equal(t, map[string]interface{}{"foo": "bar"}, res.Hits[0].Content)
}

func TestScrollEmptyScrollId(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.MessageError{Message: "Security.User.Scroll: scroll id required"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	_, err := k.Security.ScrollUsers("", nil)
	assert.NotNil(t, err)
}

func TestScrollError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	_, err := k.Security.ScrollUsers("f00b4r", nil)
	assert.NotNil(t, err)
}

func TestScroll(t *testing.T) {
	type response struct {
		Total int          `json:"total"`
		Hits  []*user.User `json:"hits"`
	}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "scrollUsers", parsedQuery.Action)

			return &types.KuzzleResponse{Result: []byte(`{
				"total": 42,
				"hits": [
					{"_id": "user42", "_source": {"profileIds": ["admin", "other"], "foo": "bar"}}
				]
			}`)}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := k.Security.ScrollUsers("f00b4r", nil)
	assert.Equal(t, 42, res.Total)
	assert.Equal(t, 1, len(res.Hits))
	assert.Equal(t, []string{"admin", "other"}, res.Hits[0].ProfileIds)
	assert.Equal(t, map[string]interface{}{"foo": "bar"}, res.Hits[0].Content)
}

func ExampleScrollUsers() {
	type response struct {
		Total int          `json:"total"`
		Hits  []*user.User `json:"hits"`
	}

	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, err := k.Security.ScrollUsers("f00b4r", nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res.Hits[0].Id, res.Hits[0].Content)
}

func TestUserCreateError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	user := k.Security.NewUser()
	user.Id = "userId"

	_, err := user.Create(nil)
	assert.NotNil(t, err)
}

func TestReplaceEmptyId(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.MessageError{Message: "Security.User.Replace: user id required"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	user := k.Security.NewUser()

	_, err := user.Replace(nil)
	assert.NotNil(t, err)
}

func TestReplaceError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	user := k.Security.NewUser()

	_, err := user.Replace(nil)
	assert.NotNil(t, err)
}

func TestReplace(t *testing.T) {
	id := "userId"

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "replaceUser", parsedQuery.Action)
			assert.Equal(t, id, parsedQuery.Id)

			return &types.KuzzleResponse{Result: []byte(`{
				"_id": "userId",
				"_source": {
					"profileIds": ["admin", "other"],
					"name": "Luke",
					"function": "Jedi"
				}
			}`)}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	user := k.Security.NewUser()
	user.Id = id
	user.ProfileIds = []string{"default", "anonymous"}
	user.SetCredentials(types.UserCredentials{
		"local": map[string]string{
			"Username": "username",
			"Password": "password",
		},
	})
	user.Content = map[string]interface{}{
		"foo": "bar",
	}

	user.Replace(nil)
}

func TestUserUpdateEmptyId(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.MessageError{Message: "Security.User.Update: user id required"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	user := k.Security.NewUser()

	_, err := user.Update(&types.UserData{}, nil)
	assert.NotNil(t, err)
}

func TestUpdateError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	user := k.Security.NewUser()

	_, err := user.Update(&types.UserData{}, nil)
	assert.NotNil(t, err)
}

func TestUserDeleteEmptyId(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.MessageError{Message: "Security.User.Delete: user id required"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	user := k.Security.NewUser()

	_, err := user.Delete(nil)
	assert.NotNil(t, err)
}

func TestDeleteError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	user := k.Security.NewUser()

	_, err := user.Delete(nil)
	assert.NotNil(t, err)
}

func TestGetRightsEmptyId(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.MessageError{Message: "Security.User.GetRights: user id required"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	_, err := k.Security.GetUserRights("", nil)
	assert.NotNil(t, err)
}

func TestGetRightsError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	_, err := k.Security.GetUserRights("userId", nil)
	assert.NotNil(t, err)
}

func TestGetRights(t *testing.T) {
	id := "userId"

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "getUserRights", parsedQuery.Action)
			assert.Equal(t, id, parsedQuery.Id)

			type resultUserRights struct {
				UserRights []*types.UserRights `json:"hits"`
			}
			userRights := []*types.UserRights{
				{Controller: "wow-controller", Action: "such-action", Index: "much indexes", Collection: "very collection", Value: "wow"},
			}
			actualRights := resultUserRights{UserRights: userRights}
			r, _ := json.Marshal(actualRights)
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	res, _ := k.Security.GetUserRights(id, nil)

	expectedRights := []*types.UserRights{
		{Controller: "wow-controller", Action: "such-action", Index: "much indexes", Collection: "very collection", Value: "wow"},
	}

	assert.Equal(t, expectedRights, res)
}

func ExampleSecurityUser_GetRights() {
	id := "userId"
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, err := k.Security.GetUserRights(id, nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res[0].Index, res[0].Collection, res[0].Controller, res[0].Action, res[0].Value)
}

func TestIsActionAllowedNilRights(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)

	_, err := k.Security.IsActionAllowed(nil, "wow-controller", "such-action", "", "")
	assert.NotNil(t, err)
}

func TestIsActionAllowedEmptyController(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)

	_, err := k.Security.IsActionAllowed([]*types.UserRights{}, "", "such-action", "", "")
	assert.NotNil(t, err)
}

func TestIsActionAllowedEmptyAction(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)

	_, err := k.Security.IsActionAllowed([]*types.UserRights{}, "wow-controller", "", "", "")
	assert.NotNil(t, err)
}

func TestIsActionAllowedEmptyRights(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)

	res, _ := k.Security.IsActionAllowed([]*types.UserRights{}, "wow-controller", "such-action", "much-index", "very-collection")

	assert.Equal(t, "denied", res)
}

func TestIsActionAllowedResultAllowed(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)

	userRights := []*types.UserRights{
		{Controller: "wow-controller", Action: "*", Index: "much-index", Collection: "very-collection", Value: "allowed"},
	}

	res, _ := k.Security.IsActionAllowed(userRights, "wow-controller", "such-action", "much-index", "very-collection")

	assert.Equal(t, "allowed", res)
}

func TestIsActionAllowedResultConditional(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)

	userRights := []*types.UserRights{
		{Controller: "wow-controller", Action: "*", Index: "much-index", Collection: "very-collection", Value: "conditional"},
	}

	res, _ := k.Security.IsActionAllowed(userRights, "wow-controller", "action", "much-index", "very-collection")

	assert.Equal(t, "conditional", res)
}

func TestIsActionAllowedResultDenied(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)

	userRights := []*types.UserRights{
		{Controller: "wow-controller.", Action: "action-such", Index: "much-index", Collection: "very-collection", Value: "allowed"},
	}

	res, _ := k.Security.IsActionAllowed(userRights, "wow-controller", "action", "much-index", "very-collection")

	assert.Equal(t, "denied", res)
}

func ExampleSecurityUser_IsActionAllowed() {
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)

	userRights := []*types.UserRights{
		{Controller: "wow-controller", Action: "*", Index: "much-index", Collection: "very-collection", Value: "allowed"},
	}

	res, err := k.Security.IsActionAllowed(userRights, "wow-controller", "such-action", "much-index", "very-collection")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res)
}
