package collection

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

// ValidateSpecifications validates the provided specifications.
func (dc *Collection) ValidateSpecifications(body string) (string, error) {
	if body == "" {
		return "", types.NewError("Collection.ValidateSpecifications: body required", 400)
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "collection",
		Action:     "validateSpecifications",
		Body:       body,
	}
	go dc.Kuzzle.Query(query, nil, ch)

	res := <-ch

	if res.Error != nil {
		return "", res.Error
	}

	var valid string
	json.Unmarshal(res.Result, &valid)

	return valid, nil
}
