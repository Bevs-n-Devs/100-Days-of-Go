package main

import "fmt"

func channelDirections() {
	fmt.Println("\nChannel Directions")

	intro := `
Channels can have directional types:

- Send-only: chan<- type
- Receive-only: <-chan type
	`
	fmt.Println(intro)

	keyPoint := `
Key Point:
- Use directional channels to enforce how channels are used, improving code clarity and safety.
	`
	fmt.Println(keyPoint)
}
