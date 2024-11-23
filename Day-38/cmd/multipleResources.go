package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func downloadResource(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Downloading resource: %d...\n", id)

	time.Sleep(time.Duration(rand.Intn(3)+1) * time.Millisecond) // random delay
	fmt.Printf("Downloaded resource: %d\n", id)
}

func waitGroupsWithMultipleResources() {
	fmt.Println("\nUsing WaitGroups with Multiple Resources")

	var wg sync.WaitGroup
	resources := 5

	for i := 1; i <= resources; i++ {
		wg.Add(1)
		go downloadResource(i, &wg)
	}

	wg.Wait()
	fmt.Println("All resources downloaded")
}
