as, err := kuzzle.Server.GetAllStats(nil)

if err != nil {
  log.Fatal(err)
} else {
  fmt.Println("All Kuzzle Stats as JSON string:", string(as))
}
