filters := json.RawMessage(`{}`)

listener := make(chan types.NotificationResult)
go func() {
  <-listener
}()

res, err := kuzzle.Realtime.Subscribe(
	"nyc-open-data",
	"yellow-taxi",
	filters,
	listener,
	nil)

if err != nil {
  log.Fatal(err)
}

err = kuzzle.Realtime.Unsubscribe(res.Room, nil)

if err != nil {
  log.Fatal(err)
} else {
  fmt.Println("Success")
}
