specifications := json.RawMessage(`{ "strict": false, "fields": { "license": { "mandatory": true, "type": "string" } } }`)
response, err := kuzzle.Collection.UpdateSpecifications("nyc-open-data", "yellow-taxi", specifications, nil)

if err != nil {
  log.Fatal(err)
} else if response != nil {
  fmt.Println("Success")
}
