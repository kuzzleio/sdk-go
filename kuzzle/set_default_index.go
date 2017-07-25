package kuzzle

import "errors"

/*
 * Default index setter
 */
func (k *Kuzzle) SetDefaultIndex(index string) error {
  if index == "" {
    return errors.New("Kuzzle.SetDefaultIndex: index required")
  }

  k.defaultIndex = index
  return nil
}