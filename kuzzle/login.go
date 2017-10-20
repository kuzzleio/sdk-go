package kuzzle

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/event"
	"github.com/kuzzleio/sdk-go/types"
)

// Login send login request to kuzzle with credentials.
// If login success, store the jwtToken into kuzzle object.
func (k *Kuzzle) Login(strategy string, credentials interface{}, expiresIn *int) (string, error) {
	if strategy == "" {
		return "", types.NewError("Kuzzle.Login: cannot authenticate to Kuzzle without an authentication strategy", 400)
	}

	type loginResult struct {
		Jwt string `json:"jwt"`
	}

	var token loginResult
	var body interface{}

	if credentials != nil {
		body = credentials
	}

	q := &types.KuzzleRequest{
		Controller: "auth",
		Action:     "login",
		Body:       body,
		Strategy:   strategy,
	}

	if expiresIn != nil {
		q.ExpiresIn = *expiresIn
	}

	result := make(chan *types.KuzzleResponse)

	go k.Query(q, nil, result)

	res := <-result

	json.Unmarshal(res.Result, &token)

	if res.Error != nil {
		k.socket.EmitEvent(event.LoginAttempt, &types.LoginAttempt{Success: false, Error: res.Error})
		return "", res.Error
	}

	k.SetJwt(token.Jwt)

	return token.Jwt, nil
}
