id, err := kuzzle.Document.Delete("nyc-open-data", "yellow-taxi", "some-id", nil)

if err != nil {
  log.Fatal(err)
} else if id == "some-id" {
  fmt.Println("Success")
}
