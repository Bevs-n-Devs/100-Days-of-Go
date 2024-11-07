package main

import (
	"fmt"
)

/*
1. Set up a basic WebSocket server in Go.
2. Manage client connections and broadcast messages to all connected clients.
3. Test our WebSocket server with a simple client.
*/

func main() {
	fmt.Println("Hello world, hello Yaw!")

	fmt.Println("\nBuilding Real-Time Applications with Websockets")

	basicWebsocketServer()

	multipleClientsBroadcastMessaging()
}
