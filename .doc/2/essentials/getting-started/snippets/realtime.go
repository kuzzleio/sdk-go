package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/protocol/websocket"
	"github.com/kuzzleio/sdk-go/types"
)

func main() {
	// Creates a WebSocket connection.
	// Replace "kuzzle" with
	// your Kuzzle hostname like "localhost"
	c := websocket.NewWebSocket("kuzzle", nil)
	// Instantiates a Kuzzle client
	kuzzle, _ := kuzzle.NewKuzzle(c, nil)

	// Connects to the server.
	if err := kuzzle.Connect(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Println("Connected!")

	// Prevents the program from exiting before receiving a notification
	exit := make(chan bool)

	// Starts an async listener
	listener := make(chan types.NotificationResult)
	go func() {
		notification := <-listener

		// Parses the document content embedded in the notification.
		var doc struct {
			Name     string `json:"name"`
			Birthday string `json:"birthday"`
			License  string `json:"license"`
		}

		json.Unmarshal(notification.Result.Content, &doc)
		fmt.Printf("Driver %s born on %s got a %s license.\n",
			doc.Name,
			doc.Birthday,
			doc.License,
		)

		// Allows the program to exit
		exit <- true
	}()

	// Subscribes to notifications for drivers having a "B" driver license.
	filters := json.RawMessage(`
		{
			"equals": {
				"license":"B"
			}
		}
	`)

	// Sends the subscription
	if _, err := kuzzle.Realtime.Subscribe(
		"nyc-open-data",
		"yellow-taxi",
		filters,
		listener,
		nil,
	); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Println("Successfully subscribed!")

	// Writes a new document. This triggers a notification sent to our subscription.
	content := json.RawMessage(`
		{
			"name": "John",
			"birthday": "1995-11-27",
			"license": "B"
		}
	`)

	if _, err := kuzzle.Document.Create(
		"nyc-open-data",
		"yellow-taxi",
		"",
		content,
		nil,
	); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Println("New document added to the yellow-taxi collection!")

	// Waits for a notification to be received
	<-exit

	// Disconnects the SDK.
	kuzzle.Disconnect()
}
