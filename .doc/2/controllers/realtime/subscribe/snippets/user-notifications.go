// Subscribe to notifications for documents containing a 'name' property
filters := json.RawMessage(`{ "exists": "name" }`)

// Start an async listener
listener := make(chan types.NotificationResult)
go func() {
  notification := <-listener

  if notification.Type == "user" {
    fmt.Printf("Volatile data: %s\n", notification.Volatile)
    // Volatile data: {"sdkVersion":"1.0.0","username":"nina vkote"}
    fmt.Printf("Currently %d users in the room\n", notification.Result.Count)
  }
}()

options := types.NewRoomOptions()
options.SetUsers(types.USERS_ALL)

_, err := kuzzle.Realtime.Subscribe(
	"nyc-open-data",
	"yellow-taxi",
	filters,
	listener,
	options)

if err != nil {
  log.Fatal(err)
}

// Instantiates a second kuzzle client: multiple subscriptions
// made by the same user will not trigger "new user" notifications
ws2 := websocket.NewWebSocket("kuzzle", nil)
kuzzle2, _ := kuzzlepkg.NewKuzzle(ws2, nil)

connectErr = kuzzle2.Connect()
if connectErr != nil {
  log.Fatal(connectErr)
  os.Exit(1)
}

// Set some volatile data
options2 := types.NewRoomOptions()
options2.SetVolatile(json.RawMessage(`{ "username": "nina vkote" }`))

// Subscribe to the same room with the second client
kuzzle2.Realtime.Subscribe(
	"nyc-open-data",
	"yellow-taxi",
	filters,
	make(chan types.NotificationResult),
	options2)
