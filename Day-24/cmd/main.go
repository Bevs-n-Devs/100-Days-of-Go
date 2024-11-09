package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

/*
Go’s scheduler is responsible for managing and scheduling goroutines, lightweight threads in Go,
allowing for efficient concurrent execution of tasks. The scheduler is part of the Go runtime and abstracts away most complexity involved in task scheduling.

Go’s concurrency model uses a few core concepts:
1. Goroutines (G): Functions or methods that run concurrently, initiated with go keyword.
2. OS Threads (M): System threads utilized by the Go runtime.
3. Processor (P): Logical processors that the scheduler uses to manage and assign tasks. By default, GOMAXPROCS is set to the number of CPU cores, but you can change it.

The scheduler works by mapping goroutines (G) onto processors (P), which are in turn mapped onto OS threads (M).

Key Components of the Scheduler
1. Goroutine (G): Each goroutine has a stack, a set of registers, and some metadata, like its status (running, waiting, etc.).
2. Processor (P): Each processor has a local run queue where goroutines are kept. The scheduler’s task is to match goroutines with processors.
3. Threads (M): Threads execute goroutines on available processors. If goroutines block (e.g., for I/O), new threads can be created to ensure other goroutines continue running.

How the Scheduler Works
1. Run Queue: Each processor has a queue of ready-to-run goroutines. When a goroutine is created, it’s placed in a local or global queue.
2. Preemption: To ensure fairness, the scheduler preempts goroutines, giving each a turn to execute. It monitors system calls and long-running goroutines to avoid starvation.
3. Work-Stealing: If a processor’s queue is empty, it can "steal" a goroutine from another processor, balancing the load across cores.

In this example, we create five goroutines, each running a small loop to simulate work.
Setting runtime.GOMAXPROCS(2) limits Go to using only two processors, illustrating how the scheduler manages multiple tasks with limited resources.
*/

func main() {
	fmt.Println("Hello world, hello Yaw!")

	fmt.Println("\nUnderstanding Go's Scheduler")

	// set the minimum number of CPUs to use
	runtime.GOMAXPROCS(2) // adjust this based on availabalke CPUs

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 1; j <= 3; j++ {
				fmt.Printf("Gorotine %d - Execution %d\n", id, j)
				time.Sleep(100 * time.Millisecond) // ssimulate some work
			}
		}(i)
	}

	wg.Wait()
}
