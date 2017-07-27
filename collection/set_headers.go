package collection

/*
  Helper function allowing to set headers.
  If the replace argument is set to true, replace the current headers with the provided content.
  Otherwise, it appends the content to the current headers, only replacing already existing values
 */
func (dc Collection) SetHeaders(content map[string]interface{}, replace bool) Collection {
	dc.kuzzle.SetHeaders(content, replace)

	return dc
}
