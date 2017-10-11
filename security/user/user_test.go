package user_test

import (
	"encoding/json"
	"fmt"
	"github.com/kuzzleio/sdk-go/connection/websocket"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/security"
	"github.com/kuzzleio/sdk-go/security/profile"
	"github.com/kuzzleio/sdk-go/security/user"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserSetContent(t *testing.T) {
	id := "userId"

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "getUser", parsedQuery.Action)
			assert.Equal(t, id, parsedQuery.Id)

			res := user.User{Id: id, Source: []byte(`{"profileIds":["admin","other"],"name":"Luke","function":"Jedi"}`)}
			r, _ := json.Marshal(res)
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	u, _ := security.NewSecurity(k).User.Fetch(id, nil)

	newContent := types.UserData{
		ProfileIds: []string{"adminNew", "otherNew"},
	}

	u.SetContent(&newContent)

	assert.Equal(t, &newContent, u.UserData())
}

func ExampleUser_SetContent() {
	id := "userId"
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	u, _ := security.NewSecurity(k).User.Fetch(id, nil)

	newContent := types.UserData{
		ProfileIds: []string{"adminNew", "otherNew"},
	}

	u.SetContent(&newContent)

	fmt.Println(u.UserData())
}

func TestUserSetCredentials(t *testing.T) {
	id := "userId"

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "getUser", parsedQuery.Action)
			assert.Equal(t, id, parsedQuery.Id)

			res := user.User{Id: id, Source: []byte(`{"profileIds":["admin","other"],"name":"Luke","function":"Jedi"}`)}
			r, _ := json.Marshal(res)
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	u, _ := security.NewSecurity(k).User.Fetch(id, nil)

	cred := types.UserCredentials{}
	json.Unmarshal([]byte(`{"local": {"Username": "username", "Password": "password"}}`), cred)

	u.SetCredentials(cred)

	assert.Equal(t, cred, u.UserData().Credentials)
}

func ExampleUser_SetCredentials() {
	id := "userId"
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	u, _ := security.NewSecurity(k).User.Fetch(id, nil)

	cred := types.UserCredentials{}
	json.Unmarshal([]byte(`{"local": {"Username": "username", "Password": "password"}}`), cred)

	res := u.SetCredentials(cred)

	fmt.Println(res.Id, res.UserData().Credentials)
}

func TestUserAddProfile(t *testing.T) {
	id := "userId"

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "getUser", parsedQuery.Action)
			assert.Equal(t, id, parsedQuery.Id)

			res := user.User{Id: id, Source: []byte(`{"profileIds":["admin","other"],"name":"Luke","function":"Jedi"}`)}
			r, _ := json.Marshal(res)
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	u, _ := security.NewSecurity(k).User.Fetch(id, nil)

	u.AddProfile(&profile.Profile{Id: "adminNew"})

	assert.Equal(t, []string{"admin", "other", "adminNew"}, u.UserData().ProfileIds)
}

func ExampleUser_AddProfile() {
	id := "userId"
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	u, _ := security.NewSecurity(k).User.Fetch(id, nil)

	res := u.AddProfile(&profile.Profile{Id: "adminNew"})

	fmt.Println(res.Id, res.UserData().ProfileIds)
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

			res := user.User{Id: id, Source: []byte(`{"profileIds":[],"name":"Luke","function":"Jedi"}`)}
			r, _ := json.Marshal(res)
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	u, _ := security.NewSecurity(k).User.Fetch(id, nil)

	profiles, _ := u.GetProfiles(nil)

	assert.Equal(t, []*profile.Profile{}, profiles)
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

				k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)

				res := user.User{Id: id, Source: []byte(`{"profileIds":["admin","other"],"name":"Luke","function":"Jedi"}`), SU: &user.SecurityUser{Kuzzle: k}}
				r, _ := json.Marshal(res)
				return &types.KuzzleResponse{Result: r}
			}

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "getProfile", parsedQuery.Action)
			assert.Equal(t, "admin", parsedQuery.Id)

			return &types.KuzzleResponse{Error: &types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	u, _ := security.NewSecurity(k).User.Fetch(id, nil)

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

				k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)

				res := user.User{Id: id, Source: []byte(`{"profileIds":["admin","other"],"name":"Luke","function":"Jedi"}`), SU: &user.SecurityUser{Kuzzle: k}}
				r, _ := json.Marshal(res)
				return &types.KuzzleResponse{Result: r}
			}
			if callCount == 1 {
				callCount++
				assert.Equal(t, "security", parsedQuery.Controller)
				assert.Equal(t, "getProfile", parsedQuery.Action)
				assert.Equal(t, "admin", parsedQuery.Id)

				k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)

				res := profile.Profile{Id: "admin", Source: []byte(`{"policies":[{"roleId":"admin"},{"roleId":"other"}]}`), SP: &profile.SecurityProfile{Kuzzle: k}}
				r, _ := json.Marshal(res)
				return &types.KuzzleResponse{Result: r}
			}
			if callCount == 2 {
				callCount++
				assert.Equal(t, "security", parsedQuery.Controller)
				assert.Equal(t, "getProfile", parsedQuery.Action)
				assert.Equal(t, "other", parsedQuery.Id)

				k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)

				res := profile.Profile{Id: "other", Source: []byte(`{"policies":[{"roleId":"other"},{"roleId":"default"}]}`), SP: &profile.SecurityProfile{Kuzzle: k}}
				r, _ := json.Marshal(res)
				return &types.KuzzleResponse{Result: r}
			}

			return &types.KuzzleResponse{Result: nil}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	u, _ := security.NewSecurity(k).User.Fetch(id, nil)

	profiles, _ := u.GetProfiles(nil)

	assert.Equal(t, []*profile.Profile{
		{Id: "admin", Source: []byte(`{"policies":[{"roleId":"admin"},{"roleId":"other"}]}`), SP: &profile.SecurityProfile{Kuzzle: k}},
		{Id: "other", Source: []byte(`{"policies":[{"roleId":"other"},{"roleId":"default"}]}`), SP: &profile.SecurityProfile{Kuzzle: k}},
	}, profiles)
}

func ExampleUser_GetProfiles() {
	id := "userId"
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	u, _ := security.NewSecurity(k).User.Fetch(id, nil)
	res, err := u.GetProfiles(nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res[0].Id, res[0].GetPolicies())
}

func TestUserGetProfileIds(t *testing.T) {
	id := "userId"

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "getUser", parsedQuery.Action)
			assert.Equal(t, id, parsedQuery.Id)

			res := user.User{Id: id, Source: []byte(`{"profileIds":["admin","other"],"name":"Luke","function":"Jedi"}`)}
			r, _ := json.Marshal(res)
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	u, _ := security.NewSecurity(k).User.Fetch(id, nil)

	profileIds := u.GetProfileIds()

	assert.Equal(t, []string{"admin", "other"}, profileIds)
}

func ExampleUser_GetProfileIds() {
	id := "userId"
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	u, _ := security.NewSecurity(k).User.Fetch(id, nil)

	res := u.GetProfileIds()

	fmt.Println(res)
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

			res := user.User{Id: id, Source: []byte(`{"profileIds":["admin","other"],"name":"Luke","function":"Jedi"}`)}
			r, _ := json.Marshal(res)
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	u, _ := security.NewSecurity(k).User.Fetch(id, nil)

	u.SetProfiles([]*profile.Profile{
		{Id: "adminNew"},
		{Id: "otherNew"},
	})

	assert.Equal(t, []string{"adminNew", "otherNew"}, u.UserData().ProfileIds)
}

func ExampleUser_SetProfiles() {
	id := "userId"
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	u, _ := security.NewSecurity(k).User.Fetch(id, nil)

	res := u.SetProfiles([]*profile.Profile{
		{Id: "adminNew"},
		{Id: "otherNew"},
	})

	fmt.Println(res.UserData().ProfileIds)
}

func TestUserCreate(t *testing.T) {
	id := "userId"
	callCount := 0

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)

			if callCount == 0 {
				callCount++
				assert.Equal(t, "security", parsedQuery.Controller)
				assert.Equal(t, "getUser", parsedQuery.Action)
				assert.Equal(t, id, parsedQuery.Id)

				res := user.User{Id: id, Source: []byte(`{"profileIds":["admin","other"],"name":"Luke","function":"Jedi"}`), SU: &user.SecurityUser{Kuzzle: k}}
				r, _ := json.Marshal(res)
				return &types.KuzzleResponse{Result: r}
			}
			if callCount == 1 {
				callCount++
				assert.Equal(t, "security", parsedQuery.Controller)
				assert.Equal(t, "createUser", parsedQuery.Action)
				assert.Equal(t, id, parsedQuery.Id)

				res := user.User{Id: id, Source: []byte(`{"profileIds":["admin","other"],"name":"Luke","function":"Master Jedi"}`), SU: &user.SecurityUser{Kuzzle: k}}
				r, _ := json.Marshal(res)
				return &types.KuzzleResponse{Result: r}
			}

			return &types.KuzzleResponse{Result: nil}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	u, _ := security.NewSecurity(k).User.Fetch(id, nil)

	u.SetContent(&types.UserData{Content: map[string]interface{}{"function": "Master Jedi"}})
	createdUser, _ := u.Create(nil)

	assert.Equal(t, "Master Jedi", createdUser.Content("function"))
}

func ExampleUser_Create() {
	id := "userId"
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	u, _ := security.NewSecurity(k).User.Fetch(id, nil)

	u.SetContent(&types.UserData{Content: map[string]interface{}{"function": "Master Jedi"}})
	res, err := u.Create(nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res.Id, res.UserData().Content)
}

func TestUserSaveRestricted(t *testing.T) {
	id := "userId"
	callCount := 0

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)

			if callCount == 0 {
				callCount++
				assert.Equal(t, "security", parsedQuery.Controller)
				assert.Equal(t, "getUser", parsedQuery.Action)
				assert.Equal(t, id, parsedQuery.Id)

				res := user.User{Id: id, Source: []byte(`{"profileIds":["admin","other"],"name":"Luke","function":"Jedi"}`), SU: &user.SecurityUser{Kuzzle: k}}
				r, _ := json.Marshal(res)
				return &types.KuzzleResponse{Result: r}
			}
			if callCount == 1 {
				callCount++
				assert.Equal(t, "security", parsedQuery.Controller)
				assert.Equal(t, "createRestrictedUser", parsedQuery.Action)
				assert.Equal(t, id, parsedQuery.Id)

				res := user.User{Id: id, Source: []byte(`{"profileIds":["admin","other"],"name":"Luke","function":"Master Jedi"}`), SU: &user.SecurityUser{Kuzzle: k}}
				r, _ := json.Marshal(res)
				return &types.KuzzleResponse{Result: r}
			}

			return &types.KuzzleResponse{Result: nil}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	u, _ := security.NewSecurity(k).User.Fetch(id, nil)

	u.SetContent(&types.UserData{Content: map[string]interface{}{"function": "Master Jedi"}})
	createdUser, _ := u.SaveRestricted(nil)

	assert.Equal(t, "Master Jedi", createdUser.Content("function"))
}

func ExampleUser_SaveRestricted() {
	id := "userId"
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	u, _ := security.NewSecurity(k).User.Fetch(id, nil)

	u.SetContent(&types.UserData{Content: map[string]interface{}{"function": "Master Jedi"}})
	res, err := u.SaveRestricted(nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res.Id, res.UserData().Content)
}

func TestUserReplace(t *testing.T) {
	id := "userId"
	callCount := 0

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)

			if callCount == 0 {
				callCount++
				assert.Equal(t, "security", parsedQuery.Controller)
				assert.Equal(t, "getUser", parsedQuery.Action)
				assert.Equal(t, id, parsedQuery.Id)

				res := user.User{Id: id, Source: []byte(`{"profileIds":["admin","other"],"name":"Luke","function":"Jedi"}`), SU: &user.SecurityUser{Kuzzle: k}}
				r, _ := json.Marshal(res)
				return &types.KuzzleResponse{Result: r}
			}
			if callCount == 1 {
				callCount++
				assert.Equal(t, "security", parsedQuery.Controller)
				assert.Equal(t, "replaceUser", parsedQuery.Action)
				assert.Equal(t, id, parsedQuery.Id)

				res := user.User{Id: id, Source: []byte(`{"profileIds":["admin","other"],"name":"Luke","function":"Master Jedi"}`), SU: &user.SecurityUser{Kuzzle: k}}
				r, _ := json.Marshal(res)
				return &types.KuzzleResponse{Result: r}
			}

			return &types.KuzzleResponse{Result: nil}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	u, _ := security.NewSecurity(k).User.Fetch(id, nil)

	u.SetContent(&types.UserData{Content: map[string]interface{}{"function": "Master Jedi"}})
	createdUser, _ := u.Replace(nil)

	assert.Equal(t, "Master Jedi", createdUser.Content("function"))
}

func ExampleUser_Replace() {
	id := "userId"
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	u, _ := security.NewSecurity(k).User.Fetch(id, nil)

	u.SetContent(&types.UserData{Content: map[string]interface{}{"function": "Master Jedi"}})
	res, err := u.Replace(nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res.Id, res.UserData().Content)
}

func TestUserUpdate(t *testing.T) {
	id := "userId"
	callCount := 0

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)

			if callCount == 0 {
				callCount++
				assert.Equal(t, "security", parsedQuery.Controller)
				assert.Equal(t, "getUser", parsedQuery.Action)
				assert.Equal(t, id, parsedQuery.Id)

				res := user.User{Id: id, Source: []byte(`{"profileIds":["admin","other"],"name":"Luke","function":"Jedi"}`), SU: &user.SecurityUser{Kuzzle: k}}
				r, _ := json.Marshal(res)
				return &types.KuzzleResponse{Result: r}
			}
			if callCount == 1 {
				callCount++
				assert.Equal(t, "security", parsedQuery.Controller)
				assert.Equal(t, "updateUser", parsedQuery.Action)
				assert.Equal(t, id, parsedQuery.Id)

				res := user.User{Id: id, Source: []byte(`{"profileIds":["adminNew","otherNew"]}`), SU: &user.SecurityUser{Kuzzle: k}}
				r, _ := json.Marshal(res)
				return &types.KuzzleResponse{Result: r}
			}

			return &types.KuzzleResponse{Result: nil}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	u, _ := security.NewSecurity(k).User.Fetch(id, nil)

	newContent := types.UserData{ProfileIds: []string{"adminNew", "otherNew"}}
	updatedUser, _ := u.Update(&newContent, nil)

	assert.Equal(t, &newContent, updatedUser.UserData())
}

func ExampleUser_Update() {
	id := "userId"
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	u, _ := security.NewSecurity(k).User.Fetch(id, nil)

	newContent := types.UserData{ProfileIds: []string{"adminNew", "otherNew"}}
	res, err := u.Update(&newContent, nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res.Id, res.UserData().Content)
}

func TestUserDelete(t *testing.T) {
	id := "SomeMenJustWantToWatchTheWorldBurn"
	callCount := 0

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)

			if callCount == 0 {
				callCount++
				assert.Equal(t, "security", parsedQuery.Controller)
				assert.Equal(t, "getUser", parsedQuery.Action)
				assert.Equal(t, id, parsedQuery.Id)

				res := user.User{Id: id, Source: []byte(`{"profileIds":["admin","other"],"name":"Luke","function":"Jedi"}`), SU: &user.SecurityUser{Kuzzle: k}}
				r, _ := json.Marshal(res)
				return &types.KuzzleResponse{Result: r}
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

	u, _ := security.NewSecurity(k).User.Fetch(id, nil)

	inTheEnd, _ := u.Delete(nil)

	assert.Equal(t, id, inTheEnd)
}

func ExampleUser_Delete() {
	id := "SomeMenJustWantToWatchTheWorldBurn"
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	u, _ := security.NewSecurity(k).User.Fetch(id, nil)

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

			k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)

			res := user.User{Id: id, Source: []byte(`{"profileIds":["admin","other"],"name":"Luke","function":"Jedi"}`), SU: &user.SecurityUser{Kuzzle: k}}
			r, _ := json.Marshal(res)
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	u, _ := security.NewSecurity(k).User.Fetch(id, nil)

	assert.Equal(t, nil, u.Content(""))
	assert.Equal(t, "Jedi", u.Content("function"))
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

			k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)

			res := user.User{Id: id, Source: []byte(`{"profileIds":["admin","other"],"name":"Luke","function":"Jedi"}`), SU: &user.SecurityUser{Kuzzle: k}}
			r, _ := json.Marshal(res)
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	u, _ := security.NewSecurity(k).User.Fetch(id, nil)

	assert.Nil(t, u.Content("galaxy"))
	assert.Equal(t, "Jedi", u.Content("function"))
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

			k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)

			res := user.User{Id: id, Source: []byte(`{"profileIds":["admin","other"],"name":"Luke","function":"Jedi"}`), SU: &user.SecurityUser{Kuzzle: k}}
			r, _ := json.Marshal(res)
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	u, _ := security.NewSecurity(k).User.Fetch(id, nil)

	assert.Equal(t, map[string]interface{}{}, u.ContentMap())
}

func TestFetchEmptyKuid(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.MessageError{Message: "Security.User.Fetch: user id required"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := security.NewSecurity(k).User.Fetch("", nil)
	assert.NotNil(t, err)
}

func TestFetchError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := security.NewSecurity(k).User.Fetch("userId", nil)
	assert.NotNil(t, err)
}

func TestFetch(t *testing.T) {
	id := "userId"

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "getUser", parsedQuery.Action)
			assert.Equal(t, id, parsedQuery.Id)

			k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)

			res := user.User{Id: id, Source: []byte(`{"profileIds":["admin","other"],"name":"Luke","function":"Jedi"}`), SU: &user.SecurityUser{Kuzzle: k}}
			r, _ := json.Marshal(res)
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := security.NewSecurity(k).User.Fetch(id, nil)

	assert.Equal(t, id, res.Id)

	assert.Equal(t, []string{"admin", "other"}, res.UserData().ProfileIds)

	assert.Equal(t, "Luke", res.Content("name"))
	assert.Equal(t, "Jedi", res.Content("function"))

	contentAsMap := make(map[string]interface{})
	contentAsMap["name"] = "Luke"
	contentAsMap["function"] = "Jedi"

	assert.Equal(t, contentAsMap, res.ContentMap("name", "function"))
}

func ExampleSecurityUser_Fetch() {
	id := "userId"
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, err := security.NewSecurity(k).User.Fetch(id, nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res.Id, res.UserData().Content)
}

func TestSearchError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := security.NewSecurity(k).User.Search(nil, nil)
	assert.NotNil(t, err)
}

func TestSearch(t *testing.T) {
	mockedK, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	hits := []*user.User{
		{Id: "user42", Source: json.RawMessage(`{"profileIds":["admin","other"],"foo":"bar"}`), SU: &user.SecurityUser{Kuzzle: mockedK}},
	}
	results := user.UserSearchResult{Total: 42, Hits: hits}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "searchUsers", parsedQuery.Action)

			res := user.UserSearchResult{Total: results.Total, Hits: results.Hits}
			r, _ := json.Marshal(res)
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := security.NewSecurity(k).User.Search(nil, nil)
	assert.Equal(t, results.Total, res.Total)
	assert.Equal(t, len(hits), len(res.Hits))
	assert.Equal(t, res.Hits[0].Id, "user42")
	assert.Equal(t, res.Hits[0].Source, json.RawMessage(`{"profileIds":["admin","other"],"foo":"bar"}`))
	assert.Equal(t, res.Hits[0].UserData().ProfileIds, []string{"admin", "other"})
	assert.Equal(t, res.Hits[0].Content("foo"), "bar")
}

func ExampleSecurityUser_Search() {
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, err := security.NewSecurity(k).User.Search(nil, nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res.Hits[0].Id, res.Hits[0].UserData().Content)
}

func TestSearchWithScroll(t *testing.T) {
	mockedK, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	hits := []*user.User{
		{Id: "user42", Source: json.RawMessage(`{"profileIds":["admin","other"],"foo":"bar"}`), SU: &user.SecurityUser{Kuzzle: mockedK}},
	}
	results := user.UserSearchResult{Total: 42, Hits: hits}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "searchUsers", parsedQuery.Action)

			res := user.UserSearchResult{Total: results.Total, Hits: results.Hits, ScrollId: "f00b4r"}
			r, _ := json.Marshal(res)
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	opts := types.NewQueryOptions()
	opts.SetFrom(2)
	opts.SetSize(4)
	opts.SetScroll("1m")

	res, _ := security.NewSecurity(k).User.Search(nil, opts)
	assert.Equal(t, results.Total, res.Total)
	assert.Equal(t, len(hits), len(res.Hits))
	assert.Equal(t, "f00b4r", res.ScrollId)
	assert.Equal(t, res.Hits[0].Id, "user42")
	assert.Equal(t, res.Hits[0].Source, json.RawMessage(`{"profileIds":["admin","other"],"foo":"bar"}`))
	assert.Equal(t, res.Hits[0].UserData().ProfileIds, []string{"admin", "other"})
	assert.Equal(t, res.Hits[0].Content("foo"), "bar")
}

func TestScrollEmptyScrollId(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.MessageError{Message: "Security.User.Scroll: scroll id required"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := security.NewSecurity(k).User.Scroll("", nil)
	assert.NotNil(t, err)
}

func TestScrollError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := security.NewSecurity(k).User.Scroll("f00b4r", nil)
	assert.NotNil(t, err)
}

func TestScroll(t *testing.T) {
	type response struct {
		Total int          `json:"total"`
		Hits  []*user.User `json:"hits"`
	}

	hits := []*user.User{
		{Id: "user42", Source: json.RawMessage(`{"profileIds":["admin","other"],"foo":"bar"}`)},
	}
	results := user.UserSearchResult{Total: 42, Hits: hits}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "scrollUsers", parsedQuery.Action)

			res := response{Total: results.Total, Hits: results.Hits}
			r, _ := json.Marshal(res)
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := security.NewSecurity(k).User.Scroll("f00b4r", nil)
	assert.Equal(t, results.Total, res.Total)
	assert.Equal(t, hits, res.Hits)
	assert.Equal(t, res.Hits[0].Id, "user42")
	assert.Equal(t, res.Hits[0].Source, json.RawMessage(`{"profileIds":["admin","other"],"foo":"bar"}`))
	assert.Equal(t, res.Hits[0].UserData().ProfileIds, []string{"admin", "other"})
	assert.Equal(t, res.Hits[0].Content("foo"), "bar")
}

func ExampleSecurityUser_Scroll() {
	type response struct {
		Total int          `json:"total"`
		Hits  []*user.User `json:"hits"`
	}

	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, err := security.NewSecurity(k).User.Scroll("f00b4r", nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res.Hits[0].Id, res.Hits[0].UserData().Content)
}

func TestCreateEmptyId(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.MessageError{Message: "Security.User.Create: user id required"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := security.NewSecurity(k).User.Create("", &types.UserData{}, nil)
	assert.NotNil(t, err)
}

func TestCreateError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := security.NewSecurity(k).User.Create("userId", &types.UserData{}, nil)
	assert.NotNil(t, err)
}

func TestCreate(t *testing.T) {
	id := "userId"

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "createUser", parsedQuery.Action)
			assert.Equal(t, id, parsedQuery.Id)

			res := user.User{
				Id:     id,
				Source: []byte(`{"profileIds":["admin","other"],"name":"Luke","function":"Jedi"}`),
			}
			r, _ := json.Marshal(res)
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	type UserContent map[string]interface{}
	cred := types.UserCredentials{}
	json.Unmarshal([]byte(`{"local": {"Username": "username", "Password": "password"}}`), cred)
	ud := types.UserData{ProfileIds: []string{"default", "anonymous"}, Content: UserContent{"foo": "bar"}, Credentials: cred}

	res, _ := security.NewSecurity(k).User.Create(id, &ud, nil)

	assert.Equal(t, id, res.Id)

	assert.Equal(t, []string{"admin", "other"}, res.UserData().ProfileIds)

	assert.Equal(t, "Luke", res.Content("name"))
	assert.Equal(t, "Jedi", res.Content("function"))

	contentAsMap := make(map[string]interface{})
	contentAsMap["name"] = "Luke"
	contentAsMap["function"] = "Jedi"

	assert.Equal(t, contentAsMap, res.ContentMap("name", "function"))
}

func ExampleSecurityUser_Create() {
	id := "userId"
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)

	type UserContent map[string]interface{}
	cred := types.UserCredentials{}
	json.Unmarshal([]byte(`{"local": {"Username": "username", "Password": "password"}}`), cred)
	ud := types.UserData{ProfileIds: []string{"default", "anonymous"}, Content: UserContent{"foo": "bar"}, Credentials: cred}

	res, err := security.NewSecurity(k).User.Create(id, &ud, nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res.Id, res.UserData().Content)
}

func TestCreateRestrictedEmptyId(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.MessageError{Message: "Security.User.CreateRestrictedUser: user id required"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := security.NewSecurity(k).User.CreateRestrictedUser("", &types.UserData{}, nil)
	assert.NotNil(t, err)
}

func TestCreateRestrictedError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := security.NewSecurity(k).User.CreateRestrictedUser("userId", &types.UserData{}, nil)
	assert.NotNil(t, err)
}

func TestCreateRestricted(t *testing.T) {
	id := "userId"

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "createRestrictedUser", parsedQuery.Action)
			assert.Equal(t, id, parsedQuery.Id)

			res := user.User{
				Id:     id,
				Source: []byte(`{"profileIds":["admin","other"],"name":"Luke","function":"Jedi"}`),
			}
			r, _ := json.Marshal(res)
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	type UserContent map[string]interface{}
	cred := types.UserCredentials{}
	json.Unmarshal([]byte(`{"local": {"Username": "username", "Password": "password"}}`), cred)
	ud := types.UserData{Content: UserContent{"foo": "bar"}, Credentials: cred}

	res, _ := security.NewSecurity(k).User.CreateRestrictedUser(id, &ud, nil)

	assert.Equal(t, id, res.Id)

	assert.Equal(t, []string{"admin", "other"}, res.UserData().ProfileIds)

	assert.Equal(t, "Luke", res.Content("name"))
	assert.Equal(t, "Jedi", res.Content("function"))

	contentAsMap := make(map[string]interface{})
	contentAsMap["name"] = "Luke"
	contentAsMap["function"] = "Jedi"

	assert.Equal(t, contentAsMap, res.ContentMap("name", "function"))
}

func ExampleSecurityUser_CreateRestrictedUser() {
	id := "userId"
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)

	type UserContent map[string]interface{}
	cred := types.UserCredentials{}
	json.Unmarshal([]byte(`{"local": {"Username": "username", "Password": "password"}}`), cred)
	ud := types.UserData{Content: UserContent{"foo": "bar"}, Credentials: cred}

	res, err := security.NewSecurity(k).User.CreateRestrictedUser(id, &ud, nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res.Id, res.UserData().Content)
}

func TestReplaceEmptyId(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.MessageError{Message: "Security.User.Replace: user id required"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := security.NewSecurity(k).User.Replace("", &types.UserData{}, nil)
	assert.NotNil(t, err)
}

func TestReplaceError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := security.NewSecurity(k).User.Replace("userId", &types.UserData{}, nil)
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

			res := user.User{
				Id:     id,
				Source: []byte(`{"profileIds":["admin","other"],"name":"Luke","function":"Jedi"}`),
			}
			r, _ := json.Marshal(res)
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	type UserContent map[string]interface{}
	cred := types.UserCredentials{}
	json.Unmarshal([]byte(`{"local": {"Username": "username", "Password": "password"}}`), cred)
	ud := types.UserData{ProfileIds: []string{"default", "anonymous"}, Content: UserContent{"foo": "bar"}, Credentials: cred}

	res, _ := security.NewSecurity(k).User.Replace(id, &ud, nil)

	assert.Equal(t, id, res.Id)

	assert.Equal(t, []string{"admin", "other"}, res.UserData().ProfileIds)

	assert.Equal(t, "Luke", res.Content("name"))
	assert.Equal(t, "Jedi", res.Content("function"))

	contentAsMap := make(map[string]interface{})
	contentAsMap["name"] = "Luke"
	contentAsMap["function"] = "Jedi"

	assert.Equal(t, contentAsMap, res.ContentMap("name", "function"))
}

func ExampleSecurityUser_Replace() {
	id := "userId"
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)

	type UserContent map[string]interface{}
	cred := types.UserCredentials{}
	json.Unmarshal([]byte(`{"local": {"Username": "username", "Password": "password"}}`), cred)
	ud := types.UserData{ProfileIds: []string{"default", "anonymous"}, Content: UserContent{"foo": "bar"}, Credentials: cred}

	res, err := security.NewSecurity(k).User.Replace(id, &ud, nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res.Id, res.UserData().Content)
}

func TestUpdateEmptyId(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.MessageError{Message: "Security.User.Update: user id required"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := security.NewSecurity(k).User.Update("", &types.UserData{}, nil)
	assert.NotNil(t, err)
}

func TestUpdateError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := security.NewSecurity(k).User.Update("userId", &types.UserData{}, nil)
	assert.NotNil(t, err)
}

func TestUpdate(t *testing.T) {
	id := "userId"

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "updateUser", parsedQuery.Action)
			assert.Equal(t, id, parsedQuery.Id)

			res := user.User{
				Id:     id,
				Source: []byte(`{"profileIds":["admin","other"],"name":"Luke","function":"Jedi"}`),
			}
			r, _ := json.Marshal(res)
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	type UserContent map[string]interface{}
	cred := types.UserCredentials{}
	json.Unmarshal([]byte(`{"local": {"Username": "username", "Password": "password"}}`), cred)
	ud := types.UserData{ProfileIds: []string{"default", "anonymous"}, Content: UserContent{"foo": "bar"}, Credentials: cred}

	res, _ := security.NewSecurity(k).User.Update(id, &ud, nil)

	assert.Equal(t, id, res.Id)

	assert.Equal(t, []string{"admin", "other"}, res.UserData().ProfileIds)

	assert.Equal(t, "Luke", res.Content("name"))
	assert.Equal(t, "Jedi", res.Content("function"))

	contentAsMap := make(map[string]interface{})
	contentAsMap["name"] = "Luke"
	contentAsMap["function"] = "Jedi"

	assert.Equal(t, contentAsMap, res.ContentMap("name", "function"))
}

func ExampleSecurityUser_Update() {
	id := "userId"
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)

	type UserContent map[string]interface{}
	cred := types.UserCredentials{}
	json.Unmarshal([]byte(`{"local": {"Username": "username", "Password": "password"}}`), cred)
	ud := types.UserData{ProfileIds: []string{"default", "anonymous"}, Content: UserContent{"foo": "bar"}, Credentials: cred}

	res, err := security.NewSecurity(k).User.Update(id, &ud, nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res.Id, res.UserData().Content)
}

func TestDeleteEmptyId(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.MessageError{Message: "Security.User.Delete: user id required"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := security.NewSecurity(k).User.Delete("", nil)
	assert.NotNil(t, err)
}

func TestDeleteError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := security.NewSecurity(k).User.Delete("userId", nil)
	assert.NotNil(t, err)
}

func TestDelete(t *testing.T) {
	id := "userId"

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "deleteUser", parsedQuery.Action)
			assert.Equal(t, id, parsedQuery.Id)

			res := types.ShardResponse{Id: id}
			r, _ := json.Marshal(res)
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := security.NewSecurity(k).User.Delete(id, nil)

	assert.Equal(t, id, res)
}

func ExampleSecurityUser_Delete() {
	id := "userId"
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, err := security.NewSecurity(k).User.Delete(id, nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res)
}

func TestGetRightsEmptyId(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.MessageError{Message: "Security.User.GetRights: user id required"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := security.NewSecurity(k).User.GetRights("", nil)
	assert.NotNil(t, err)
}

func TestGetRightsError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := security.NewSecurity(k).User.GetRights("userId", nil)
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

	res, _ := security.NewSecurity(k).User.GetRights(id, nil)

	expectedRights := []*types.UserRights{
		{Controller: "wow-controller", Action: "such-action", Index: "much indexes", Collection: "very collection", Value: "wow"},
	}

	assert.Equal(t, expectedRights, res)
}

func ExampleSecurityUser_GetRights() {
	id := "userId"
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, err := security.NewSecurity(k).User.GetRights(id, nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res[0].Index, res[0].Collection, res[0].Controller, res[0].Action, res[0].Value)
}

func TestIsActionAllowedNilRights(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)

	_, err := security.NewSecurity(k).User.IsActionAllowed(nil, "wow-controller", "such-action", "", "")
	assert.NotNil(t, err)
}

func TestIsActionAllowedEmptyController(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)

	_, err := security.NewSecurity(k).User.IsActionAllowed([]*types.UserRights{}, "", "such-action", "", "")
	assert.NotNil(t, err)
}

func TestIsActionAllowedEmptyAction(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)

	_, err := security.NewSecurity(k).User.IsActionAllowed([]*types.UserRights{}, "wow-controller", "", "", "")
	assert.NotNil(t, err)
}

func TestIsActionAllowedEmptyRights(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)

	res, _ := security.NewSecurity(k).User.IsActionAllowed([]*types.UserRights{}, "wow-controller", "such-action", "much-index", "very-collection")

	assert.Equal(t, "denied", res)
}

func TestIsActionAllowedResultAllowed(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)

	userRights := []*types.UserRights{
		{Controller: "wow-controller", Action: "*", Index: "much-index", Collection: "very-collection", Value: "allowed"},
	}

	res, _ := security.NewSecurity(k).User.IsActionAllowed(userRights, "wow-controller", "such-action", "much-index", "very-collection")

	assert.Equal(t, "allowed", res)
}

func TestIsActionAllowedResultConditional(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)

	userRights := []*types.UserRights{
		{Controller: "wow-controller", Action: "*", Index: "much-index", Collection: "very-collection", Value: "conditional"},
	}

	res, _ := security.NewSecurity(k).User.IsActionAllowed(userRights, "wow-controller", "action", "much-index", "very-collection")

	assert.Equal(t, "conditional", res)
}

func TestIsActionAllowedResultDenied(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)

	userRights := []*types.UserRights{
		{Controller: "wow-controller.", Action: "action-such", Index: "much-index", Collection: "very-collection", Value: "allowed"},
	}

	res, _ := security.NewSecurity(k).User.IsActionAllowed(userRights, "wow-controller", "action", "much-index", "very-collection")

	assert.Equal(t, "denied", res)
}

func ExampleSecurityUser_IsActionAllowed() {
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)

	userRights := []*types.UserRights{
		{Controller: "wow-controller", Action: "*", Index: "much-index", Collection: "very-collection", Value: "allowed"},
	}

	res, err := security.NewSecurity(k).User.IsActionAllowed(userRights, "wow-controller", "such-action", "much-index", "very-collection")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res)
}
