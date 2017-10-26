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

			return &types.KuzzleResponse{
				Error: types.NewError("Test error"),
			}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	u, _ := k.Security.FetchUser(id, nil)

	_, err := u.GetProfiles(nil)
	assert.NotNil(t, err)
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
			assert.Equal(t, map[string]interface{}{
				"content": map[string]interface{}{
					"function": "Master Jedi",
				},
				"credentials": map[string]interface{}{
					"local": "credentials",
				},
				"profileIds": []interface{}{"admin"},
			}, parsedQuery.Body)

			return &types.KuzzleResponse{Result: []byte{}}
		},
	}

	k, _ := kuzzle.NewKuzzle(c, nil)
	u := k.Security.NewUser("", &types.UserData{
		Content: map[string]interface{}{
			"function": "Master Jedi",
		},
		ProfileIds: []string{"admin"},
	})
	u.SetCredentials("local", "credentials")

	_, err := u.Create(nil)
	assert.Nil(t, err)
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

func TestUserSaveRestrictedEmptyId(t *testing.T) {
	c := &internal.MockedConnection{}
	k, _ := kuzzle.NewKuzzle(c, nil)
	user := k.Security.NewUser("", nil)

	_, err := user.SaveRestricted(nil)
	assert.NotNil(t, err)
	assert.Equal(t, 400, err.(*types.KuzzleError).Status)
	assert.Equal(t, "User.SaveRestricted: id is required", err.(*types.KuzzleError).Message)
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
				assert.Equal(t, map[string]interface{}{
					"content": map[string]interface{}{
						"name":     "Luke",
						"function": "Master Jedi",
					},
					"credentials": map[string]interface{}{"local": "credentials"},
				}, parsedQuery.Body)

				return &types.KuzzleResponse{Result: []byte{}}
			}

			return &types.KuzzleResponse{Result: nil}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	u, _ := k.Security.FetchUser(id, nil)

	u.Content["function"] = "Master Jedi"
	u.SetCredentials("local", "credentials")
	_, err := u.SaveRestricted(nil)
	assert.Nil(t, err)
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
				assert.Equal(t, map[string]interface{}{
					"profileIds": []interface{}{"admin", "other"},
					"name":       "Luke",
					"function":   "Master Jedi",
				}, parsedQuery.Body)

				return &types.KuzzleResponse{Result: []byte{}}
			}

			return &types.KuzzleResponse{Result: nil}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	u, _ := k.Security.FetchUser(id, nil)

	u.Content["function"] = "Master Jedi"
	_, err := u.Replace(nil)
	assert.Nil(t, err)
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
				// nb: profile ids order is not guaranteed
				assert.Equal(t, map[string]interface{}{
					"profileIds": []interface{}{"admin", "other", "adminNew", "otherNew"},
					"name":       "Luke",
					"function":   "Jedi",
					"weapon":     "lightsaber",
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
		ProfileIds: []string{"adminNew", "other", "otherNew"},
		Content: map[string]interface{}{
			"weapon": "lightsaber",
		},
	}
	updatedUser, _ := u.Update(&newContent, nil)

	assert.Equal(t, []string{"admin", "other", "adminNew", "otherNew"}, updatedUser.ProfileIds)
	assert.Equal(t, map[string]interface{}{
		"name":     "Luke",
		"function": "Jedi",
		"weapon":   "lightsaber",
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
					"_id": "` + id + `",
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

func TestUserCreateError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{
				Error: types.NewError("Test error"),
			}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	user := k.Security.NewUser("userId", nil)

	_, err := user.Create(nil)
	assert.NotNil(t, err)
}

func TestReplaceEmptyId(t *testing.T) {
	c := &internal.MockedConnection{}
	k, _ := kuzzle.NewKuzzle(c, nil)
	user := k.Security.NewUser("", nil)

	_, err := user.Replace(nil)
	assert.NotNil(t, err)
}

func TestReplaceError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{
				Error: types.NewError("Test error"),
			}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	user := k.Security.NewUser("", nil)

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

	user := k.Security.NewUser(id, &types.UserData{
		Content: map[string]interface{}{
			"foo": "bar",
		},
		ProfileIds: []string{"default", "anonymous"},
	})
	user.SetCredentials("local", map[string]string{
		"Username": "username",
		"Password": "password",
	})

	user.Replace(nil)
}

func TestUserUpdateEmptyId(t *testing.T) {
	c := &internal.MockedConnection{}
	k, _ := kuzzle.NewKuzzle(c, nil)
	user := k.Security.NewUser("", nil)

	_, err := user.Update(&types.UserData{}, nil)
	assert.NotNil(t, err)
}

func TestUpdateError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{
				Error: types.NewError("Test error"),
			}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	user := k.Security.NewUser("userId", nil)

	_, err := user.Update(&types.UserData{}, nil)
	assert.NotNil(t, err)
}

func TestUserDeleteEmptyId(t *testing.T) {
	c := &internal.MockedConnection{}
	k, _ := kuzzle.NewKuzzle(c, nil)
	user := k.Security.NewUser("", nil)

	_, err := user.Delete(nil)
	assert.NotNil(t, err)
}

func TestDeleteError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{
				Error: types.NewError("Test error"),
			}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	user := k.Security.NewUser("userId", nil)

	_, err := user.Delete(nil)
	assert.NotNil(t, err)
}

func TestGetRightsEmptyId(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{
				Error: types.NewError("Test error"),
			}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	_, err := k.Security.GetUserRights("", nil)
	assert.NotNil(t, err)
}

func TestGetRightsError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{
				Error: types.NewError("Test error"),
			}
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
