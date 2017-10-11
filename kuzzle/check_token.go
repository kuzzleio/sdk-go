package kuzzle

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

type TokenValidity struct {
	Valid     bool   `json:"valid"`
	State     string `json:"state"`
	ExpiresAt int    `json:"expiresAt"`
}

// CheckToken checks the validity of a JSON Web Token.
func (k Kuzzle) CheckToken(token string) (*TokenValidity, error) {
	tokenValidity := &TokenValidity{}

	if token == "" {
		return tokenValidity, types.NewError("Kuzzle.CheckToken: token required")
	}

	result := make(chan *types.KuzzleResponse)

	type body struct {
		Token string `json:"token"`
	}

	query := &types.KuzzleRequest{
		Controller: "auth",
		Action:     "checkToken",
		Body:       &body{token},
	}
	go k.Query(query, nil, result)

	res := <-result

	if res.Error == nil {
		json.Unmarshal(res.Result, tokenValidity)
	}

	return tokenValidity, res.Error
}
