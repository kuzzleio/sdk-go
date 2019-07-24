conf, err := kuzzle.Server.GetConfig(nil)

if err != nil {
  log.Fatal(err)
} else {
  fmt.Println("Kuzzle Server configuration as JSON string:", string(conf))
}
