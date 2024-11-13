package concurrencypatterns

import (
	"fmt"
	"time"
)

/*
Rate limiting helps control the frequency of operations to avoid overloading the system.
*/

func RateLimiting() {
	fmt.Println("\nConcurrency Patterns: Rate Limiting / Throttling")

	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	for i := 1; i <= 5; i++ {
		<-ticker.C // wait for the ticker
		fmt.Println("Operation", i)
	}
}
