package main

import (
	"fmt"
	"sync"
)

/*
A Mutex is a structure in Go that can be used to lock and unlock critical sections of code.
Only one Goroutine can hold the lock at a time, which ensures that only one Goroutine is executing that section of code.

Mutex: A sync.Mutex is used to prevent multiple Goroutines from modifying the counter variable at the same time.
Lock and Unlock: The mutex.Lock() ensures that only one Goroutine can modify the counter variable. After the update, the lock is released with mutex.Unlock().
sync.WaitGroup: This ensures that the main function waits until all Goroutines have finished executing.
Without the Mutex, the counter value might be incorrect because of race conditions.
*/

var (
	counter int
	mutex   sync.Mutex
)

func increment(wg *sync.WaitGroup) {
	defer wg.Done()

	// locking the critical section
	mutex.Lock()
	counter++
	// unlocking after modification
	mutex.Unlock()
}

func main() {
	fmt.Println("Hello world, hello Yaw!")

	fmt.Println("Mutext Basics")

	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment(&wg)
	}

	wg.Wait() // wait for all Goroutines to finish
	fmt.Println("Final Counter:", counter)

	bankBalanceMutexExample()

	// deadlock() // deadlock error

	solvingDeadlocks()

	readWriteMutex()

	atomicOperations()

	usingMutex()

	usingChannel()

}

/*
Understand how to use Mutexes to prevent race conditions in shared resources.
Be aware of potential pitfalls like deadlocks and how to avoid them with careful locking and unlocking practices.
Know the difference between Mutexes and RWMutexes, and when to use each.
Explore the use of atomic operations for simpler concurrency scenarios.
Recognize the trade-offs between using Mutexes and Channels for concurrency control.
*/
