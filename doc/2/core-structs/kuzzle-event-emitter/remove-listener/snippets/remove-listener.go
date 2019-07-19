ch := make(chan json.RawMessage)

kuzzle.On(event.Connected, ch)

go func() {
  for _ = range ch {
    fmt.Println("Connected to Kuzzle")
  }

  fmt.Println("Stopped listening")
}()

kuzzle.RemoveListener(event.Connected, ch)
close(ch)
