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

func TestFetchUserEmptyId(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Security.User.Fetch: user id required"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := security.NewSecurity(k).User.Fetch("", nil)
	assert.NotNil(t, err)
}

func TestFetchUserError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := security.NewSecurity(k).User.Fetch("userId", nil)
	assert.NotNil(t, err)
}

func TestFetchUser(t *testing.T) {
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

func TestGetRightsEmptyId(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Security.User.GetRights: user id required"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := security.NewSecurity(k).User.GetRights("", nil)
	assert.NotNil(t, err)
}

func TestGetRightsError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
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
		MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "getUserRights", parsedQuery.Action)
			assert.Equal(t, id, parsedQuery.Id)

			type resultUserRights struct {
				UserRights []types.UserRights `json:"hits"`
			}
			userRights := []types.UserRights{}
			userRights = append(userRights, types.UserRights{Controller:"wow-controller", Action: "such-action", Index: "much-index", Collection: "very-collection", Value: "wow"})
			actualRights := resultUserRights{UserRights: userRights}
			r, _ := json.Marshal(actualRights)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := security.NewSecurity(k).User.GetRights(id, nil)

	expectedRights := []types.UserRights{}
	expectedRights = append(expectedRights, types.UserRights{Controller:"wow-controller", Action: "such-action", Index: "much-index", Collection: "very-collection", Value: "wow"})

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
	userRights = append(userRights, types.UserRights{Controller:"wow-controller", Action: "*", Index: "much-index", Collection: "very-collection", Value: "allowed"})

	res, _ := security.NewSecurity(k).User.IsActionAllowed(userRights, "wow-controller", "such-action", "much-index", "very-collection")

	assert.Equal(t, "allowed", res)
}

func TestIsActionAllowedResultConditional(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)

	userRights := []types.UserRights{}
	userRights = append(userRights, types.UserRights{Controller:"wow-controller", Action: "*", Index: "much-index", Collection: "very-collection", Value: "conditional"})

	res, _ := security.NewSecurity(k).User.IsActionAllowed(userRights, "wow-controller", "action", "much-index", "very-collection")

	assert.Equal(t, "conditional", res)
}

func TestIsActionAllowedResultDenied(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)

	userRights := []types.UserRights{}
	userRights = append(userRights, types.UserRights{Controller:"wow-controller.", Action: "action-such", Index: "much-index", Collection: "very-collection", Value: "allowed"})

	res, _ := security.NewSecurity(k).User.IsActionAllowed(userRights, "wow-controller", "action", "much-index", "very-collection")

	assert.Equal(t, "denied", res)
}
