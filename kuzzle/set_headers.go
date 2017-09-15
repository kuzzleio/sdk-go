package kuzzle


// SetHeaders is a helper function allowing to set headers.
// If the replace argument is set to true, replace the current headers with the provided content.
// Otherwise, it appends the content to the current headers, only replacing already existing values
func (k *Kuzzle) SetHeaders(content map[string]interface{}, replace bool) {
	if replace {
		k.headers = content
		return
	}

	for i, v := range content {
		k.headers[i] = v
	}
}
