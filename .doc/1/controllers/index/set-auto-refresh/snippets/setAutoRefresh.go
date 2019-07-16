err := kuzzle.Index.SetAutoRefresh("nyc-open-data", true, nil)

if err != nil {
  log.Fatal(err)
} else {
  fmt.Println("autorefresh flag is set to true")
}
