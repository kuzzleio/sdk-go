package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Object inspects the low-level properties of a key.
func (ms Ms) Object(key string, subcommand string, options types.QueryOptions) (string, error) {
	if key == "" {
		return "", types.NewError("Ms.Object: key required")
	}
	if subcommand != "refcount" && subcommand != "encoding" && subcommand != "idletime" {
		return "", types.NewError("Ms.Object: subcommand required, possible values: refcount|encoding|idletime")
	}

	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "object",
		Id:         key,
		Subcommand: subcommand,
	}
	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return "", res.Error
	}
	var returnedResult string
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
