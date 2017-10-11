package kuzzle

import "github.com/kuzzleio/sdk-go/types"

// SetDefaultIndex set the default data index. Has the same effect than the defaultIndex constructor option.
func (k Kuzzle) SetDefaultIndex(index string) error {
	if index == "" {
		return types.NewError("Kuzzle.SetDefaultIndex: index required")
	}

	k.defaultIndex = index
	return nil
}
