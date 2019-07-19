// Subscribe to notifications for documents containing a 'name' property
filters := json.RawMessage(`{ "exists": "name" }`)

// Start an async listener
listener := make(chan types.NotificationResult)
go func() {
  notification := <-listener

  if notification.Scope == "in" {
    fmt.Printf("Document %s enter the scope\n", notification.Result.Id)
  } else {
    fmt.Printf("Document %s leave the scope\n", notification.Result.Id)
  }
}()

_, err := kuzzle.Realtime.Subscribe(
	"nyc-open-data",
	"yellow-taxi",
	filters,
	listener,
	nil)

if err != nil {
  log.Fatal(err)
}

document := json.RawMessage(`{ "name": "nina vkote", "age": 19 }`)
kuzzle.Document.Create(
	"nyc-open-data",
	"yellow-taxi",
	"nina-vkote",
	document,
	nil)
