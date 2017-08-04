package kuzzle

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/event"
	"github.com/kuzzleio/sdk-go/types"
)

func (k *Kuzzle) Login(strategy string, credentials interface{}, expiresIn *int) (string, error) {
	if strategy == "" {
		return "", errors.New("Kuzzle.Login: cannot authenticate to Kuzzle without an authentication strategy")
	}

	type loginResult struct {
		Jwt string `json:"jwt"`
	}

	var token loginResult
	var body interface{}

	if credentials != nil {
		body = credentials
	}

	q := types.KuzzleRequest{
		Controller: "auth",
		Action:     "login",
		Body:       body,
		Strategy:   strategy,
	}

	if expiresIn != nil {
		q.ExpiresIn = *expiresIn
	}

	result := make(chan types.KuzzleResponse)

	go k.Query(q, nil, result)

	res := <-result

	json.Unmarshal(res.Result, &token)

	if res.Error.Message != "" {
		err := errors.New(res.Error.Message)
		k.socket.EmitEvent(event.LoginAttempt, types.LoginAttempt{Success: false, Error: err})
		return "", err
	}

	k.jwt = token.Jwt

	if token.Jwt != "" {
		k.socket.RenewSubscriptions()
		k.socket.EmitEvent(event.LoginAttempt, types.LoginAttempt{Success: true})
	}

	return token.Jwt, nil
}
