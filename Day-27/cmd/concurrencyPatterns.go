package main

import (
	"fmt"

	concurrencypatterns "github.com/Bevs-n-Devs/100-Days-of-Go/Day-27/cmd/concurrency_patterns"
)

/*
Go offers a variety of concurrency patterns, each designed to address specific problems in concurrent applications.
Here’s a list of essential concurrency patterns in Go, along with examples:

1. Worker Pools
2. Fan-Out / Fan-In
3. Pipeline Pattern
4. Publish-Subscribe (Pub-Sub)
5. Select Statement for Multiplexing
6. Rate Limiting / Throttling
7. Or-Done Channel Pattern
8. Errgroup for Coordinated Concurrency
*/

func concurrencyPatterns() {
	fmt.Println("\nConcurrency Patterns in Go")

	concurrencypatterns.WorkerPools()

	concurrencypatterns.FanInFanOut()

	concurrencypatterns.PipelinePattern()

	concurrencypatterns.PubSub()

	concurrencypatterns.SelectStatements()

	concurrencypatterns.RateLimiting()

	concurrencypatterns.OrDone()

	concurrencypatterns.ErrgroupCoordinatedConcurrency()
}
