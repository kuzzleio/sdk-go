indexes, err := kuzzle.Index.List(nil)

if err != nil {
  log.Fatal(err)
} else {
  fmt.Printf("Kuzzle contains %d indexes", len(indexes))
}
