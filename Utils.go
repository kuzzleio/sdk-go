package core

import (
	"github.com/kuzzleio/sdk-core/types"
)

func MakeQuery(controller, action, index, collection string, body interface{}) types.KuzzleRequest {
	return types.KuzzleRequest{
		Controller: controller,
		Action: action,
		Index: index,
		Collection: collection,
		Body: body,
	}
}
