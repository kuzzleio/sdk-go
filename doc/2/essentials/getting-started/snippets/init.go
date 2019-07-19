package main

import (
	"fmt"
	"os"
	"log"

	"github.com/kuzzleio/sdk-go/protocol/websocket"
	"github.com/kuzzleio/sdk-go/kuzzle"
)

func main()	{
	// Creates a WebSocket connection.
	// Replace "kuzzle" with
	// your Kuzzle hostname like "localhost"
	c := websocket.NewWebSocket("kuzzle", nil)
	// Instantiates a Kuzzle client
	kuzzle, _ := kuzzle.NewKuzzle(c, nil)

	// Connects to the server.
	err := kuzzle.Connect()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Println("Connected!")

	// Freshly installed Kuzzle servers are empty: we need to create
	// a new index.
	if err := kuzzle.Index.Create("nyc-open-data", nil); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Println("Index nyc-open-data created!")

	// Creates a collection 
	if err := kuzzle.Collection.Create(
		"nyc-open-data",
		"yellow-taxi",
		nil,
		nil,
	); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Println("Collection yellow-taxi created!")

	// Disconnects the SDK
	kuzzle.Disconnect()
}
