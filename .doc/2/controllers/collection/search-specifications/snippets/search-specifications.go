options := types.NewQueryOptions()
options.SetFrom(0)
options.SetSize(2)

response, err := kuzzle.Collection.SearchSpecifications(json.RawMessage(`{
  "query": {
    "match_all": {} 
  }
}`), options)

if err != nil {
  log.Fatal(err)
} else {
  fmt.Printf("Successfully retrieved %d specifications", response.Total)
}
