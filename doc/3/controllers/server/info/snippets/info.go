info, err := kuzzle.Server.Info(nil)

if err != nil {
  log.Fatal(err)
} else {
  fmt.Println("Kuzzle Server information as JSON string:", string(info))
}
