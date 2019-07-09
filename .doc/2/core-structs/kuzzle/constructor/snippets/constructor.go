copts := types.NewOptions()
copts.SetPort(7512)
copts.SetAutoResubscribe(false)
conn := websocket.NewWebSocket("kuzzle", copts)

k, _ := kuzzle.NewKuzzle(conn, nil)
