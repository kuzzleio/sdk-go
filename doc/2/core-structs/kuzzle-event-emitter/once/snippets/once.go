ch := make(chan json.RawMessage)

kuzzle.Once(event.Connected, ch)

go func() {
  for _ = range ch {
    fmt.Println("Connected to Kuzzle")
  }
}()
