err := kuzzle.Index.Delete("nyc-open-data", nil)

if err != nil {
  log.Fatal(err)
} else {
  fmt.Println("index deleted")
}
