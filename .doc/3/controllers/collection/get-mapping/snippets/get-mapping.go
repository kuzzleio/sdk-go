mapping, err := kuzzle.Collection.GetMapping("nyc-open-data", "yellow-taxi", nil)

if err != nil {
  log.Fatal(err)
} else if mapping != nil {
  fmt.Println("Success")
}
