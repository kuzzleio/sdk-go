ch := make(chan json.RawMessage)

kuzzle.AddListener(event.Connected, ch)

go func() {
  for range ch {
    fmt.Println("Connected to Kuzzle")
  }
}()
