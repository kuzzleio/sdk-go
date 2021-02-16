ls, err := kuzzle.Server.GetLastStats(nil)

if err != nil {
  log.Fatal(err)
} else {
  fmt.Println("Last Kuzzle Stats as JSON string:", string(ls))
}
