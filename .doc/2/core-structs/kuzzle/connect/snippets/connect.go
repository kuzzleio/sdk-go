err := kuzzle.Connect()

if err != nil {
  log.Fatal(err)
} else {
  fmt.Println("Successfully connected")
}
