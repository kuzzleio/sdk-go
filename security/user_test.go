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

func TestUser_GetProfiles_EmptyProfileIds(t *testing.T) {
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

func TestUser_GetProfiles_Error(t *testing.T) {
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

func TestUser_GetProfiles(t *testing.T) {
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

func TestUser_Create_Error(t *testing.T) {
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

func TestUser_Create(t *testing.T) {
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

func TestUser_CreateCredentials(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "createCredentials", parsedQuery.Action)
			assert.Equal(t, "strategy", parsedQuery.Strategy)
			assert.Equal(t, "userid", parsedQuery.Id)
			assert.Equal(t, "myCredentials", parsedQuery.Body)

			return &types.KuzzleResponse{Result: []byte{}}
		},
	}

	k, err := kuzzle.NewKuzzle(c, nil)
	user := k.Security.NewUser("userid", nil)
	_, err = user.CreateCredentials("strategy", "myCredentials", nil)

	assert.Nil(t, err)
}

func TestUser_CreateWithCredentials(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "createUser", parsedQuery.Action)
			assert.Equal(t, map[string]interface{}{
				"content": map[string]interface{}{
					"function": "Jedi",
				},
				"credentials": map[string]interface{}{
					"strategy": "myCredentials",
				},
				"profileIds": []interface{}{"profile1", "profile2"},
			}, parsedQuery.Body)

			return &types.KuzzleResponse{Result: []byte{}}
		},
	}

	k, err := kuzzle.NewKuzzle(c, nil)
	user := k.Security.NewUser("userid", &types.UserData{
		Content:    map[string]interface{}{"function": "Jedi"},
		ProfileIds: []string{"profile1", "profile2"},
	})
	_, err = user.CreateWithCredentials(types.Credentials{"strategy": "myCredentials"}, nil)

	assert.Nil(t, err)
}

func TestUser_Delete_EmptyId(t *testing.T) {
	c := &internal.MockedConnection{}
	k, _ := kuzzle.NewKuzzle(c, nil)
	user := k.Security.NewUser("", nil)

	_, err := user.Delete(nil)
	assert.NotNil(t, err)
}

func TestUser_Delete_Error(t *testing.T) {
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

func TestUser_Delete(t *testing.T) {
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

func TestUser_DeleteCredentials(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "deleteCredentials", parsedQuery.Action)
			assert.Equal(t, "strategy", parsedQuery.Strategy)
			assert.Equal(t, "userid", parsedQuery.Id)

			return &types.KuzzleResponse{}
		},
	}

	k, err := kuzzle.NewKuzzle(c, nil)
	user := k.Security.NewUser("userid", nil)

	_, err = user.DeleteCredentials("strategy", nil)

	assert.Nil(t, err)
}

func TestUser_GetCredentialsInfo_EmptyStrategy(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	user := k.Security.NewUser("userid", nil)
	_, err := user.GetCredentialsInfo("", nil)

	assert.NotNil(t, err)
	assert.Equal(t, 400, err.(*types.KuzzleError).Status)
	assert.Equal(t, "Security.getCredentials: strategy is required", err.(*types.KuzzleError).Message)
}

func TestUser_GetCredentialsInfo_EmptyUserId(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	user := k.Security.NewUser("", nil)
	_, err := user.GetCredentialsInfo("strategy", nil)

	assert.NotNil(t, err)
	assert.Equal(t, 400, err.(*types.KuzzleError).Status)
	assert.Equal(t, "Security.getCredentials: user id is required", err.(*types.KuzzleError).Message)
}

func TestUser_GetCredentialsInfo(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			q := &types.KuzzleRequest{}
			json.Unmarshal(query, q)

			assert.Equal(t, "security", q.Controller)
			assert.Equal(t, "getCredentials", q.Action)
			assert.Equal(t, "strategy", q.Strategy)
			assert.Equal(t, "userid", q.Id)

			return &types.KuzzleResponse{}
		},
	}

	k, _ := kuzzle.NewKuzzle(c, nil)
	user := k.Security.NewUser("userid", nil)

	_, err := user.GetCredentialsInfo("strategy", nil)
	assert.Nil(t, err)
}

func TestUser_GetRights_EmptyId(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{
				Error: types.NewError("Test error"),
			}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	user := k.Security.NewUser("", nil)
	_, err := user.GetRights(nil)
	assert.NotNil(t, err)
}

func TestUser_GetRights_Error(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{
				Error: types.NewError("Test error"),
			}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	user := k.Security.NewUser("userid", nil)
	_, err := user.GetRights(nil)
	assert.NotNil(t, err)
}

func TestUser_GetRights(t *testing.T) {
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
	user := k.Security.NewUser(id, nil)
	res, _ := user.GetRights(nil)

	expectedRights := []*types.UserRights{
		{Controller: "wow-controller", Action: "such-action", Index: "much indexes", Collection: "very collection", Value: "wow"},
	}

	assert.Equal(t, expectedRights, res)
}

func ExampleUser_GetRights() {
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	user := k.Security.NewUser("userid", nil)

	res, err := user.GetRights(nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res[0].Index, res[0].Collection, res[0].Controller, res[0].Action, res[0].Value)
}

func TestUser_HasCredentials(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "hasCredentials", parsedQuery.Action)
			assert.Equal(t, "strategy", parsedQuery.Strategy)
			assert.Equal(t, "userid", parsedQuery.Id)

			return &types.KuzzleResponse{}
		},
	}

	k, err := kuzzle.NewKuzzle(c, nil)
	user := k.Security.NewUser("userid", nil)

	_, err = user.HasCredentials("strategy", nil)
	assert.Nil(t, err)
}

func TestUserSaveRestrictedEmptyId(t *testing.T) {
	c := &internal.MockedConnection{}
	k, _ := kuzzle.NewKuzzle(c, nil)
	user := k.Security.NewUser("", nil)

	_, err := user.SaveRestricted(nil, nil)
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
				}, parsedQuery.Body)

				return &types.KuzzleResponse{Result: []byte{}}
			}

			return &types.KuzzleResponse{Result: nil}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	u, _ := k.Security.FetchUser(id, nil)

	u.Content["function"] = "Master Jedi"
	_, err := u.SaveRestricted(nil, nil)
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
	res, err := u.SaveRestricted(nil, nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res.Id, res.Content)
}

func TestUser_Replace_EmptyId(t *testing.T) {
	c := &internal.MockedConnection{}
	k, _ := kuzzle.NewKuzzle(c, nil)
	user := k.Security.NewUser("", nil)

	_, err := user.Replace(nil)
	assert.NotNil(t, err)
}

func TestUser_Replace_Error(t *testing.T) {
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

func TestUser_Replace(t *testing.T) {
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

func TestUser_UpdateCredentials(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			q := &types.KuzzleRequest{}
			json.Unmarshal(query, q)

			assert.Equal(t, "security", q.Controller)
			assert.Equal(t, "updateCredentials", q.Action)
			assert.Equal(t, "strategy", q.Strategy)
			assert.Equal(t, "userid", q.Id)
			assert.Equal(t, "myCredentials", q.Body)

			return &types.KuzzleResponse{}
		},
	}

	k, _ := kuzzle.NewKuzzle(c, nil)
	user := k.Security.NewUser("userid", nil)

	_, err := user.UpdateCredentials("strategy", "myCredentials", nil)
	assert.Nil(t, err)
}
