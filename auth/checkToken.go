package auth

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
func (a *Auth) CheckToken(token string) (*TokenValidity, error) {
	if token == "" {
		return nil, types.NewError("Kuzzle.CheckToken: token required", 400)
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
	go a.kuzzle.Query(query, nil, result)

	res := <-result

	if res.Error != nil {
		return nil, res.Error
	}

	tokenValidity := &TokenValidity{}
	json.Unmarshal(res.Result, tokenValidity)

	return tokenValidity, nil
}
