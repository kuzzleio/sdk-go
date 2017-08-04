package kuzzle

/*
  Returns every headers.
*/
func (k Kuzzle) GetHeaders() map[string]interface{} {
	return k.headers
}

/*
  Returns a specific header using its key.
*/
func (k Kuzzle) GetHeader(key string) interface{} {
	return k.headers[key]
}
