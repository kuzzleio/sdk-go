exists, err := kuzzle.Collection.Exists("nyc-open-data", "green-taxi", nil)

if err != nil {
  log.Fatal(err)
} else if exists {
  fmt.Println("Success")
}
