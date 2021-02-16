exists, err := kuzzle.Index.Exists("nyc-open-data", nil)

if err != nil {
  log.Fatal(err)
} else if exists == true {
  fmt.Println("index exists")
} else {
  fmt.Println("index does not exist")
}
