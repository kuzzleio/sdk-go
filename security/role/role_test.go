package role_test

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
      return types.KuzzleResponse{Error: types.MessageError{Message: "Security.Role.Fetch: role id required"}}
    },
  }
  k, _ := kuzzle.NewKuzzle(c, nil)

  _, err := security.NewSecurity(k).Role.Fetch("", nil)
  assert.NotNil(t, err)
}

func TestFetchError(t *testing.T) {
  c := &internal.MockedConnection{
    MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
      return types.KuzzleResponse{Error: types.MessageError{Message: "Unit test error"}}
    },
  }
  k, _ := kuzzle.NewKuzzle(c, nil)

  _, err := security.NewSecurity(k).Role.Fetch("roleId", nil)
  assert.NotNil(t, err)
}

func TestFetch(t *testing.T) {
  id := "roleId"

  c := &internal.MockedConnection{
    MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
      parsedQuery := &types.KuzzleRequest{}
      json.Unmarshal(query, parsedQuery)

      assert.Equal(t, "security", parsedQuery.Controller)
      assert.Equal(t, "getRole", parsedQuery.Action)
      assert.Equal(t, id, parsedQuery.Id)

      res := types.Role{Id: id, Source: []byte(`{"controllers":{"*":{"actions":{"*":true}}}}`)}
      r, _ := json.Marshal(res)
      return types.KuzzleResponse{Result: r}
    },
  }
  k, _ := kuzzle.NewKuzzle(c, nil)

  res, _ := security.NewSecurity(k).Role.Fetch(id, nil)

  assert.Equal(t, id, res.Id)
  assert.Equal(t, true, res.Controllers()["*"].Actions["*"])
}

func TestCreateEmptyId(t *testing.T) {
  c := &internal.MockedConnection{
    MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
      return types.KuzzleResponse{Error: types.MessageError{Message: "Security.Role.Create: role id required"}}
    },
  }
  k, _ := kuzzle.NewKuzzle(c, nil)

  _, err := security.NewSecurity(k).Role.Create("", types.Controllers{}, nil)
  assert.NotNil(t, err)
}

func TestCreateError(t *testing.T) {
  c := &internal.MockedConnection{
    MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
      return types.KuzzleResponse{Error: types.MessageError{Message: "Unit test error"}}
    },
  }
  k, _ := kuzzle.NewKuzzle(c, nil)

  _, err := security.NewSecurity(k).Role.Create("roleId", types.Controllers{}, nil)
  assert.NotNil(t, err)
}

func TestCreate(t *testing.T) {
  id := "roleId"

  c := &internal.MockedConnection{
    MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
      parsedQuery := &types.KuzzleRequest{}
      json.Unmarshal(query, parsedQuery)

      assert.Equal(t, "security", parsedQuery.Controller)
      assert.Equal(t, "createRole", parsedQuery.Action)
      assert.Equal(t, id, parsedQuery.Id)

      res := types.Role{
        Id: id,
        Source: []byte(`{"controllers":{"*":{"actions":{"*":true}}}}`),
      }
      r, _ := json.Marshal(res)
      return types.KuzzleResponse{Result: r}
    },
  }
  k, _ := kuzzle.NewKuzzle(c, nil)

  res, _ := security.NewSecurity(k).Role.Create(id, types.Controllers{map[string]types.Controller{"*":{map[string]bool{"*": true}}}}, nil)

  assert.Equal(t, id, res.Id)
  assert.Equal(t, true, res.Controllers()["*"].Actions["*"])
}

func TestCreateIfExists(t *testing.T) {
  id := "roleId"

  c := &internal.MockedConnection{
    MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
      parsedQuery := &types.KuzzleRequest{}
      json.Unmarshal(query, parsedQuery)

      assert.Equal(t, "security", parsedQuery.Controller)
      assert.Equal(t, "createOrReplaceRole", parsedQuery.Action)
      assert.Equal(t, id, parsedQuery.Id)

      res := types.Role{
        Id: id,
        Source: []byte(`{"controllers":{"*":{"actions":{"*":true}}}}`),
      }
      r, _ := json.Marshal(res)
      return types.KuzzleResponse{Result: r}
    },
  }
  k, _ := kuzzle.NewKuzzle(c, nil)

  options := &types.Options{IfExist: "replace"}
  res, _ := security.NewSecurity(k).Role.Create(id, types.Controllers{map[string]types.Controller{"*":{map[string]bool{"*": true}}}}, options)

  assert.Equal(t, id, res.Id)
  assert.Equal(t, true, res.Controllers()["*"].Actions["*"])
}

func TestCreateWithStrictOption(t *testing.T) {
  id := "roleId"

  c := &internal.MockedConnection{
    MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
      parsedQuery := &types.KuzzleRequest{}
      json.Unmarshal(query, parsedQuery)

      assert.Equal(t, "security", parsedQuery.Controller)
      assert.Equal(t, "createRole", parsedQuery.Action)
      assert.Equal(t, id, parsedQuery.Id)

      res := types.Role{
        Id: id,
        Source: []byte(`{"controllers":{"*":{"actions":{"*":true}}}}`),
      }
      r, _ := json.Marshal(res)
      return types.KuzzleResponse{Result: r}
    },
  }
  k, _ := kuzzle.NewKuzzle(c, nil)

  options := &types.Options{IfExist: "error"}
  res, _ := security.NewSecurity(k).Role.Create(id, types.Controllers{map[string]types.Controller{"*":{map[string]bool{"*": true}}}}, options)

  assert.Equal(t, id, res.Id)
  assert.Equal(t, true, res.Controllers()["*"].Actions["*"])
}

func TestCreateWithWrongOption(t *testing.T) {
  id := "roleId"

  c := &internal.MockedConnection{
    MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
      return types.KuzzleResponse{}
    },
  }
  k, _ := kuzzle.NewKuzzle(c, nil)

  options := &types.Options{IfExist: "unknown"}
  _, err := security.NewSecurity(k).Role.Create(id, types.Controllers{map[string]types.Controller{"*":{map[string]bool{"*": true}}}}, options)

  assert.Equal(t, "Invalid value for the 'ifExist' option: 'unknown'", fmt.Sprint(err))
}

func TestUpdateEmptyId(t *testing.T) {
  c := &internal.MockedConnection{
    MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
      return types.KuzzleResponse{Error: types.MessageError{Message: "Security.Role.Update: role id required"}}
    },
  }
  k, _ := kuzzle.NewKuzzle(c, nil)

  _, err := security.NewSecurity(k).Role.Update("", types.Controllers{}, nil)
  assert.NotNil(t, err)
}

func TestUpdateError(t *testing.T) {
  c := &internal.MockedConnection{
    MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
      return types.KuzzleResponse{Error: types.MessageError{Message: "Unit test error"}}
    },
  }
  k, _ := kuzzle.NewKuzzle(c, nil)

  _, err := security.NewSecurity(k).Role.Update("roleId", types.Controllers{}, nil)
  assert.NotNil(t, err)
}

func TestUpdate(t *testing.T) {
  id := "roleId"

  c := &internal.MockedConnection{
    MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
      parsedQuery := &types.KuzzleRequest{}
      json.Unmarshal(query, parsedQuery)

      assert.Equal(t, "security", parsedQuery.Controller)
      assert.Equal(t, "updateRole", parsedQuery.Action)
      assert.Equal(t, id, parsedQuery.Id)

      res := types.Role{
        Id: id,
        Source: []byte(`{"controllers":{"*":{"actions":{"*":true}}}}`),
      }
      r, _ := json.Marshal(res)
      return types.KuzzleResponse{Result: r}
    },
  }
  k, _ := kuzzle.NewKuzzle(c, nil)

  res, _ := security.NewSecurity(k).Role.Update(id, types.Controllers{map[string]types.Controller{"*":{map[string]bool{"*": true}}}}, nil)

  assert.Equal(t, id, res.Id)
  assert.Equal(t, true, res.Controllers()["*"].Actions["*"])
}

func TestDeleteEmptyId(t *testing.T) {
  c := &internal.MockedConnection{
    MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
      return types.KuzzleResponse{Error: types.MessageError{Message: "Security.Role.Delete: role id required"}}
    },
  }
  k, _ := kuzzle.NewKuzzle(c, nil)

  _, err := security.NewSecurity(k).Role.Delete("", nil)
  assert.NotNil(t, err)
}

func TestDeleteError(t *testing.T) {
  c := &internal.MockedConnection{
    MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
      return types.KuzzleResponse{Error: types.MessageError{Message: "Unit test error"}}
    },
  }
  k, _ := kuzzle.NewKuzzle(c, nil)

  _, err := security.NewSecurity(k).Role.Delete("roleId", nil)
  assert.NotNil(t, err)
}

func TestDelete(t *testing.T) {
  id := "roleId"

  c := &internal.MockedConnection{
    MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
      parsedQuery := &types.KuzzleRequest{}
      json.Unmarshal(query, parsedQuery)

      assert.Equal(t, "security", parsedQuery.Controller)
      assert.Equal(t, "deleteRole", parsedQuery.Action)
      assert.Equal(t, id, parsedQuery.Id)

      res := types.ShardResponse{Id: id, }
      r, _ := json.Marshal(res)
      return types.KuzzleResponse{Result: r}
    },
  }
  k, _ := kuzzle.NewKuzzle(c, nil)

  res, _ := security.NewSecurity(k).Role.Delete(id, nil)

  assert.Equal(t, id, res)
}
