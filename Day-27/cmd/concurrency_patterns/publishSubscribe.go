package concurrencypatterns

import (
	"fmt"
	"sync"
)

/*
In the Pub-Sub pattern, publishers send messages to multiple subscribers via channels.
*/

type Broker struct {
	subscribers []chan string
	mutext      sync.Mutex
}

func (b *Broker) Subscribe() <-chan string {
	ch := make(chan string)
	b.mutext.Lock()
	b.subscribers = append(b.subscribers, ch)
	b.mutext.Unlock()
	return ch
}

func (b *Broker) Publish(msg string) {
	b.mutext.Lock()
	for _, suscriber := range b.subscribers {
		suscriber <- msg
	}
	b.mutext.Unlock()
}

func (b *Broker) Close() {
	b.mutext.Lock()
	for _, suscriber := range b.subscribers {
		close(suscriber)
	}
	b.mutext.Unlock()
}

func PubSub() {
	fmt.Println("\nConcurrency Pattern: Publish-Subscribe (Pub-Sub) Pattern")

	broker := &Broker{}

	sub1 := broker.Subscribe()
	sub2 := broker.Subscribe()

	go broker.Publish("Hello Suscribers!")

	fmt.Println("Suscriber 1 received:", <-sub1)
	fmt.Println("Suscriber 2 received:", <-sub2)

	broker.Close()
}
