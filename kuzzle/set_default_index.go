package kuzzle

import "errors"

// SetDefaultIndex set the default data index. Has the same effect than the defaultIndex constructor option.
func (k Kuzzle) SetDefaultIndex(index string) error {
	if index == "" {
		return errors.New("Kuzzle.SetDefaultIndex: index required")
	}

	k.defaultIndex = index
	return nil
}
