package kuzzle

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

type TokenValidity struct {
	Valid     bool   `json:"valid"`
	State     string `json:"state"`
	ExpiresAt int    `json:"expiresAt"`
}

// CheckToken checks the validity of a JSON Web Token.
func (k Kuzzle) CheckToken(token string) (TokenValidity, error) {
	if token == "" {
		return TokenValidity{}, errors.New("Kuzzle.CheckToken: token required")
	}

	result := make(chan types.KuzzleResponse)

	type body struct {
		Token string `json:"token"`
	}

	query := types.KuzzleRequest{
		Controller: "auth",
		Action:     "checkToken",
		Body:       &body{token},
	}
	go k.Query(query, nil, result)

	res := <-result

	if res.Error.Message != "" {
		return TokenValidity{}, errors.New(res.Error.Message)
	}
	tokenValidity := TokenValidity{}
	json.Unmarshal(res.Result, &tokenValidity)

	return tokenValidity, nil
}
