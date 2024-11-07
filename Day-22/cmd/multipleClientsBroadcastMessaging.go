package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader2 = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// to support multiple clients we'll use a map to store active connections and a broadcast
// channel to send messages to all connected clients
var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan []byte)

// Function to handle incoming WebSocket connections
// This function also registers each new client and removes them upon disconnection
func handleConnections2(w http.ResponseWriter, r *http.Request) {
	// Upgrade the initial HTTP request to a WebSocket connection
	ws, err := upgrader2.Upgrade(w, r, nil)
	if err != nil {
		// Log any upgrade error and stop further execution
		log.Fatal(err)
	}
	// Ensure the WebSocket connection is closed when the function completes
	defer ws.Close()

	// Register the new WebSocket client in the 'clients' map, setting it to active (true)
	clients[ws] = true

	// Continuously listen for messages from this WebSocket client
	for {
		// Read a new message from the WebSocket client
		_, message, err := ws.ReadMessage()
		if err != nil {
			// If an error occurs (e.g., client disconnects), remove the client from 'clients' map
			delete(clients, ws)
			break // Exit the loop, ending communication with this client
		}
		// Send the received message to the 'broadcast' channel to distribute to other clients
		broadcast <- message
	}
}

// Function to handle incoming messages and broadcast them to all connected clients
func handleMessages() {
	for {
		// Retrieve the next message from the 'broadcast' channel
		message := <-broadcast
		// Loop through each registered client in the 'clients' map
		for client := range clients {
			// Send the message to the current client
			err := client.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				// If an error occurs (e.g., client has disconnected), log it and remove client
				log.Printf("Error sending message: %v", err)
				client.Close()          // Close the client connection
				delete(clients, client) // Remove client from the map
			}
		}
	}
}

// Function to start the WebSocket server with message broadcasting capabilities
func multipleClientsBroadcastMessaging() {
	// Print a message to indicate the broadcast server is active
	fmt.Println("\nManaging Multiple Clients and Broadcast Messages")

	// Map the "/ws2" endpoint to handleConnections2 to manage WebSocket connections
	http.HandleFunc("/ws2", handleConnections2)

	// Start the 'handleMessages' goroutine to continuously broadcast messages to all clients
	go handleMessages()

	// Start the HTTP server on port 8080, logging any fatal errors if they occur
	fmt.Println("WebSocket server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
