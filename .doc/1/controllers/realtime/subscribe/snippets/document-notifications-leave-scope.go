// Subscribe to notifications for documents containing a 'name' property
filters := json.RawMessage(`{ "range": { "age": { "lte": 20 } } }`)

// Start an async listener
listener := make(chan types.NotificationResult)
go func() {
  notification := <-listener

  fmt.Printf("Document moved %s from the scope\n", notification.Scope)
}()

options := types.NewRoomOptions()
options.SetScope(types.SCOPE_OUT)

_, err := kuzzle.Realtime.Subscribe(
	"nyc-open-data",
	"yellow-taxi",
	filters,
	listener,
	options)

if err != nil {
  log.Fatal(err)
}

document := json.RawMessage(`{ "name": "nina vkote", "age": 19 }`)

// The document is in the scope
kuzzle.Document.Create(
	"nyc-open-data",
	"yellow-taxi",
	"nina-vkote",
	document,
	nil)

// The document isn't in the scope anymore
kuzzle.Document.Update(
	"nyc-open-data",
	"yellow-taxi",
	"nina-vkote",
	json.RawMessage(`{ "age": 42 }`),
	nil)
