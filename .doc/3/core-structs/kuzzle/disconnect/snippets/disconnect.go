err := kuzzle.Disconnect()

if err != nil {
  log.Fatal(err)
} else {
  fmt.Println("Success")
}
