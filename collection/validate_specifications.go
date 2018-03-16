package collection

import (
	"encoding/json"
	"fmt"

	"github.com/kuzzleio/sdk-go/types"
)

// ValidateSpecifications validates the provided specifications.
func (dc *Collection) ValidateSpecifications(body json.RawMessage) (bool, error) {
	if body == nil {
		return false, types.NewError("Collection.ValidateSpecifications: body required", 400)
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
		return false, res.Error
	}

	var validationRes struct {
		Valid       bool
		Details     []string
		Descritpion string
	}

	err := json.Unmarshal(res.Result, &validationRes)

	if err != nil {
		return false, types.NewError(fmt.Sprintf("Unable to parse response: %s\n%s", err.Error(), res.Result), 500)
	}

	return validationRes.Valid, nil
}
