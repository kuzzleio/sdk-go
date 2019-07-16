specifications := json.RawMessage(`{ "strict": false, "fields": { "license": { "mandatory": true, "type": "string" } } }`)
vr, err := kuzzle.Collection.ValidateSpecifications("nyc-open-data", "yellow-taxi", specifications, nil)

if err != nil {
  log.Fatal(err)
} else if vr.Valid == true {
  fmt.Println("Success")
}
