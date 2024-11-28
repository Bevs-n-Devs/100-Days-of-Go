package main

import "fmt"

func main() {
	fmt.Println("Hello world, hello Yaw!")

	fmt.Println("\nDeep Dive into WebSockets in Go")

	intro := `
WebSockets are a powerful protocol for enabling real-time, full-duplex communication between a client and server. 
They are especially valuable for financial applications such as live trading dashboards, notification systems, 
or any use case requiring low-latency updates.
	`
	fmt.Println(intro)

	whatAreWebSockets := `
	What are WebSockets?

- Persistent Connection: Unlike HTTP, which is stateless and follows a request-response model, WebSockets establish a persistent connection between the client and server.
- Full Duplex: Communication can flow simultaneously in both directions (client to server and server to client).
- Efficient: Eliminates the overhead of HTTP headers, making it suitable for high-frequency updates.
	`
	fmt.Println(whatAreWebSockets)

	webSocketsInFinance := `
	Why Use WebSockets in Finance?

- Real-time Data Feeds: Stock prices, cryptocurrency tickers, or transaction updates.
- Instant Notifications: Alerts for significant events like price changes or trades.
- Collaborative Applications: Shared dashboards or multi-user trading interfaces.
	`
	fmt.Println(webSocketsInFinance)

	howToUseWebSockets := `
	WebSockets in Go

The github.com/gorilla/websocket package is a popular library for handling WebSockets in Go.

Install the Gorilla WebSocket library:

	go get -u github.com/gorilla/websocket
	`
	fmt.Println(howToUseWebSockets)

	simpleWebSocketServer()
}
