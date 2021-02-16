err := kuzzle.Collection.Refresh("nyc-open-data", "yellow-taxi", nil)

if err != nil {
  log.Fatal(err)
} else {
  fmt.Println("Success")
}
