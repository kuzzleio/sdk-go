// Start an async listener
listener := make(chan types.NotificationResult)
go func() {
  <-listener

  fmt.Printf("Message notification received")
}()

// Subscribe to a room
_, err := kuzzle.Realtime.Subscribe(
	"i-dont-exist",
	"i-database",
	json.RawMessage(`{}`),
	listener,
	nil)

if err != nil {
  log.Fatal(err)
}

message := json.RawMessage(`{ "metAt": "Insane", "hello": "world" }`)
// Publish a message to this room
kuzzle.Realtime.Publish("i-dont-exist", "i-database", message, nil)
