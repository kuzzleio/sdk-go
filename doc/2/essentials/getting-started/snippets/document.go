package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/protocol/websocket"
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

	// New document content
	content := json.RawMessage(`
		{
			"name": "Sirkis",
			"birthday": "1959-06-22",
			"license": "B"
		}
	`)

	// Stores the document in the "yellow-taxi" collection.
	if _, err := kuzzle.Document.Create(
		"nyc-open-data",
		"yellow-taxi",
		"some-id",
		content,
		nil,
	); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Println("New document added to the yellow-taxi collection!")

	// Disconnects the SDK.
	kuzzle.Disconnect()
}
