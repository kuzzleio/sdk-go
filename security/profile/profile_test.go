package profile_test

import (
  "testing"
  "github.com/kuzzleio/sdk-go/internal"
  "github.com/kuzzleio/sdk-go/kuzzle"
  "github.com/stretchr/testify/assert"
  "encoding/json"
  "github.com/kuzzleio/sdk-go/types"
  "github.com/kuzzleio/sdk-go/security"
  "fmt"
)

func TestFetchEmptyId(t *testing.T) {
  c := &internal.MockedConnection{
    MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
      return types.KuzzleResponse{Error: types.MessageError{Message: "Security.Profile.Fetch: profile id required"}}
    },
  }
  k, _ := kuzzle.NewKuzzle(c, nil)

  _, err := security.NewSecurity(k).Profile.Fetch("", nil)
  assert.NotNil(t, err)
}

func TestFetchError(t *testing.T) {
  c := &internal.MockedConnection{
    MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
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
    MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
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

func TestCreateEmptyId(t *testing.T) {
  c := &internal.MockedConnection{
    MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
      return types.KuzzleResponse{Error: types.MessageError{Message: "Security.Profile.Create: profile id required"}}
    },
  }
  k, _ := kuzzle.NewKuzzle(c, nil)

  _, err := security.NewSecurity(k).Profile.Create("", types.Policies{}, nil)
  assert.NotNil(t, err)
}

func TestCreateError(t *testing.T) {
  c := &internal.MockedConnection{
    MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
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
    MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
      parsedQuery := &types.KuzzleRequest{}
      json.Unmarshal(query, parsedQuery)

      assert.Equal(t, "security", parsedQuery.Controller)
      assert.Equal(t, "createProfile", parsedQuery.Action)
      assert.Equal(t, id, parsedQuery.Id)

      res := types.Profile{
        Id: id,
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
    MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
      parsedQuery := &types.KuzzleRequest{}
      json.Unmarshal(query, parsedQuery)

      assert.Equal(t, "security", parsedQuery.Controller)
      assert.Equal(t, "createOrReplaceProfile", parsedQuery.Action)
      assert.Equal(t, id, parsedQuery.Id)

      res := types.Profile{
        Id: id,
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
  options := &types.Options{IfExist: "replace"}
  res, _ := security.NewSecurity(k).Profile.Create(id, types.Policies{Policies: policies}, options)

  assert.Equal(t, id, res.Id)
  assert.Equal(t, []string{"admin", "other"}, res.Policies())
}

func TestCreateWithStrictOption(t *testing.T) {
  id := "profileId"

  c := &internal.MockedConnection{
    MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
      parsedQuery := &types.KuzzleRequest{}
      json.Unmarshal(query, parsedQuery)

      assert.Equal(t, "security", parsedQuery.Controller)
      assert.Equal(t, "createProfile", parsedQuery.Action)
      assert.Equal(t, id, parsedQuery.Id)

      res := types.Profile{
        Id: id,
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
  options := &types.Options{IfExist: "error"}
  res, _ := security.NewSecurity(k).Profile.Create(id, types.Policies{Policies: policies}, options)

  assert.Equal(t, id, res.Id)
  assert.Equal(t, []string{"admin", "other"}, res.Policies())
}

func TestCreateWithWrongOption(t *testing.T) {
  id := "profileId"

  c := &internal.MockedConnection{
    MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
      return types.KuzzleResponse{}
    },
  }
  k, _ := kuzzle.NewKuzzle(c, nil)

  policies := []types.Policy{}
  policies = append(policies, types.Policy{RoleId: "admin"})
  policies = append(policies, types.Policy{RoleId: "other"})
  options := &types.Options{IfExist: "unknown"}
  _, err := security.NewSecurity(k).Profile.Create(id, types.Policies{Policies: policies}, options)

  assert.Equal(t, "Invalid value for the 'ifExist' option: 'unknown'", fmt.Sprint(err))
}

func TestUpdateEmptyId(t *testing.T) {
  c := &internal.MockedConnection{
    MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
      return types.KuzzleResponse{Error: types.MessageError{Message: "Security.Profile.Update: profile id required"}}
    },
  }
  k, _ := kuzzle.NewKuzzle(c, nil)

  _, err := security.NewSecurity(k).Profile.Update("", types.Policies{}, nil)
  assert.NotNil(t, err)
}

func TestUpdateError(t *testing.T) {
  c := &internal.MockedConnection{
    MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
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
    MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
      parsedQuery := &types.KuzzleRequest{}
      json.Unmarshal(query, parsedQuery)

      assert.Equal(t, "security", parsedQuery.Controller)
      assert.Equal(t, "updateProfile", parsedQuery.Action)
      assert.Equal(t, id, parsedQuery.Id)

      res := types.Profile{
        Id: id,
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
    MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
      return types.KuzzleResponse{Error: types.MessageError{Message: "Security.Profile.Delete: profile id required"}}
    },
  }
  k, _ := kuzzle.NewKuzzle(c, nil)

  _, err := security.NewSecurity(k).Profile.Delete("", nil)
  assert.NotNil(t, err)
}

func TestDeleteError(t *testing.T) {
  c := &internal.MockedConnection{
    MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
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
    MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
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
