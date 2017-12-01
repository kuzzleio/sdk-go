package ms

import (
	"github.com/kuzzleio/sdk-go/types"
)

// Set creates a key holding the provided value, or overwrites it if it already exists.
func (ms *Ms) Set(key string, value interface{}, options types.QueryOptions) error {
	result := make(chan *types.KuzzleResponse)

	type body struct {
		Value interface{} `json:"value"`
		Ex    int         `json:"ex,omitempty"`
		Px    int         `json:"px,omitempty"`
		Nx    bool        `json:"nx"`
		Xx    bool        `json:"xx"`
	}

	bodyContent := body{Value: value}

	if options != nil {
		if options.Ex() != 0 {
			bodyContent.Ex = options.Ex()
		}

		if options.Px() != 0 {
			bodyContent.Px = options.Px()
		}

		bodyContent.Nx = options.Nx()
		bodyContent.Xx = options.Xx()
	}

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "set",
		Id:         key,
		Body:       &bodyContent,
	}

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	return res.Error
}
