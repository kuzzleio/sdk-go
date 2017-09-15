package kuzzle_test

import (
	"fmt"
	"github.com/kuzzleio/sdk-go/connection/websocket"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/kuzzleio/sdk-go/kuzzle"
)

func ExampleKuzzle_Connect() {
	opts := types.NewOptions()
	opts.SetConnect(types.Manual)

	conn := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(conn, opts)

	err := k.Connect()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(k.State)
}

func ExampleKuzzle_GetJwt() {
	conn := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(conn, nil)

	jwt := k.GetJwt()
	fmt.Println(jwt)
}

func ExampleKuzzle_GetOfflineQueue() {
	//todo
}