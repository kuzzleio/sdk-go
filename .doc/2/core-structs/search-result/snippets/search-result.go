for i := 0; i < 5; i++ {
  kuzzle.Document.Create("nyc-open-data", "yellow-taxi", "", json.RawMessage(`{
    "category": "suv"
  }`), nil)
}
for i := 5; i < 15; i++ {
  kuzzle.Document.Create("nyc-open-data", "yellow-taxi", "", json.RawMessage(`{
    "category": "limousine"
  }`), nil)
}
kuzzle.Index.Refresh("nyc-open-data", nil)

options := types.NewQueryOptions()
options.SetScroll("1m")
options.SetSize(2)

response, err := kuzzle.Document.Search("nyc-open-data", "yellow-taxi", json.RawMessage(`{
  "query": {
    "match": {
      "category": "suv"
    }
  }
}`), options)

nextPage, err := response.Next()

if err != nil {
  log.Fatal(err)
} else {
  fmt.Printf("Successfully retrieved %d documents", nextPage.Fetched)
}