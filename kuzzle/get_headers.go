package kuzzle

// GetHeaders returns every headers.
func (k Kuzzle) GetHeaders() map[string]interface{} {
	return k.headers
}

// GetHeader returns a specific header using its key.
func (k Kuzzle) GetHeader(key string) interface{} {
	return k.GetHeaders()[key]
}
