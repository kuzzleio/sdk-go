message := json.RawMessage(`{ "realtime": "rule the web" }`)

err := kuzzle.Realtime.Publish("i-dont-exist", "i-database", message, nil)

if err != nil {
  log.Fatal(err)
} else {
  fmt.Println("Success")
}
