err := kuzzle.Index.Create("nyc-open-data", nil)

if err != nil {
  log.Fatal(err)
} else {
  fmt.Println("index created")
}
