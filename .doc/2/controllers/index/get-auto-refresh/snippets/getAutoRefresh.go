status, err := kuzzle.Index.GetAutoRefresh("nyc-open-data", nil)

if err != nil {
  log.Fatal(err)
} else if status == true {
  fmt.Println("autorefresh is true")
} else {
  fmt.Println("autorefresh is false")
}
