package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

/*
  Creates a key holding the provided value, or overwrites it if it already exists.
*/
func (ms Ms) Set(key string, value interface{}, options types.QueryOptions) (string, error) {
	if key == "" {
		return "", errors.New("Ms.Set: key required")
	}

	result := make(chan types.KuzzleResponse)

	type body struct {
		Value interface{} `json:"value"`
		Ex    int         `json:"ex,omitempty"`
		Px    int         `json:"px,omitempty"`
		Nx    bool        `json:"nx"`
		Xx    bool        `json:"xx"`
	}

	bodyContent := body{Value: value}

	if options != nil {
		if options.GetEx() != 0 {
			bodyContent.Ex = options.GetEx()
		}

		if options.GetPx() != 0 {
			bodyContent.Px = options.GetPx()
		}

		bodyContent.Nx = options.GetNx()
		bodyContent.Xx = options.GetXx()
	}

	query := types.KuzzleRequest{
		Controller: "ms",
		Action:     "set",
		Id:         key,
		Body:       &bodyContent,
	}

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error.Message != "" {
		return "", errors.New(res.Error.Message)
	}

	var returnedResult string
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
