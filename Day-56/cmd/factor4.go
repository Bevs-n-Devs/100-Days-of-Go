package main

import "fmt"

func Factor4() {
	var factor4 = `
	4. Backing Services: Treat Backing Services as Attached Resources.
		- Principle: Treat database, message queues, or other external services as replaceable services.
		- Why: Enables swapping services without changing the app code.
		
		Example:
			- Replace an in-memory cache with a Redis instance by changing the 'CACHE_URL' environment variable.
	`
	fmt.Println(factor4)
}
