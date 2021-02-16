specifications, err := kuzzle.Collection.GetSpecifications("nyc-open-data", "yellow-taxi", nil)

if err != nil {
  log.Fatal(err)
} else if specifications != nil {
  fmt.Println("Success")
}
