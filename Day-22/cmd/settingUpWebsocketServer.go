package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

/*
Upgrading HTTP to WebSocket: upgrader.Upgrade converts a standard HTTP connection to a WebSocket.
Read/Write Loop: We use ReadMessage and WriteMessage to echo messages back to the client.
*/

// config the upgrader that will convert HTTP requests into WebSocket connections
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024, // size for the buffer to read data from the WebSocket
	WriteBufferSize: 1024, // size for the buffer to write data to the WebSocket
	CheckOrigin: func(r *http.Request) bool {
		// Allow all cross-origin WebSocket requests by returning true.
		// This can be restricted based on origin headers if needed.
		return true
	},
}

// function to handle incoming WebSocket connections
func handleConnections(w http.ResponseWriter, r *http.Request) {
	// upgrade initial GET request to a WebSocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		// log error and exit if the upgrader fails
		log.Fatal(err)
	}
	// make sure WebSocket closes when function ends
	defer ws.Close()

	// continuously read messages from the WebSocket connection
	for {
		// read message from the WebSocket connection
		messageType, message, err := ws.ReadMessage()
		if err != nil {
			// log any read errors and close the connection & infinate loop
			log.Println("Error reading message:", err)
			break
		}
		// echo the message back to the client
		if err := ws.WriteMessage(messageType, message); err != nil {
			// log any write errors and close the connection & infinate loop
			log.Println("Error writing message:", err)
			break
		}
	}
}

// sets up an endpoint for the WebSocket connection, runs the server and logs any fatal errors
func basicWebsocketServer() {
	fmt.Println("\nCreating Basic Websoket Server")

	// map the /ws1 endpoint to the handleConnections function
	http.HandleFunc("/ws1", handleConnections)

	// print a message to indicate WebSocket is running
	fmt.Println("WebSocket server started on :8080")

	// start the HTTP server on port 8080 and log any fatal errors
	log.Fatal(http.ListenAndServe(":8080", nil))
}
