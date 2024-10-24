package main

import (
	"fmt"
	"time"
)

/*
A Goroutine is a function or method that runs concurrently with other Goroutines.
To create a Goroutine, you simply prepend the go keyword before a function call.
*/

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	fmt.Println("Hello world, hello Yaw!")

	fmt.Println("\nGoroutines and How To Use Them")
	// start go routine
	go say("Hello world")
	say("Hello Yaw!")

	goroutineAnonymousFunctions()

	channels()

	goroutinesAndChannels()

	bufferdChannels()

	closingChannels()

	selectStatements()

	practicalExercises()
}
