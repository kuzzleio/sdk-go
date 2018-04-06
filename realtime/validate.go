package realtime

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

// Validate validates data against existing validation rules
func (r *Realtime) Validate(index string, collection string, body string) (bool, error) {
	if (index == "" || collection == "") || body == "" {
		return false, types.NewError("Realtime.Validate: index, collection and body required", 400)
	}

	query := &types.KuzzleRequest{
		Controller: "realtime",
		Action:     "validate",
		Index:      index,
		Collection: collection,
		Body:       body,
	}

	result := make(chan *types.KuzzleResponse)

	go r.k.Query(query, nil, result)

	res := <-result

	if res.Error != nil {
		return false, res.Error
	}

	var isValid struct {
		Value bool `json:"valid"`
	}

	json.Unmarshal(res.Result, &isValid)

	return isValid.Value, nil
}
