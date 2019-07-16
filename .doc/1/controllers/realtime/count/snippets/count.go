filters := json.RawMessage(`{}`)

listener := make(chan types.NotificationResult)
go func() {
  <-listener
}()

res, _ := kuzzle.Realtime.Subscribe(
	"nyc-open-data",
	"yellow-taxi",
	filters,
	listener,
	nil)

count, err := kuzzle.Realtime.Count(res.Room, nil)

if err != nil {
  log.Fatal(err)
} else {
  fmt.Printf("Currently %d active subscription\n", count)
}
