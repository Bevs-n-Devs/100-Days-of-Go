package main

import "fmt"

func main() {
	fmt.Println("Hello world, hello Yaw!")

	fmt.Println("\nDeep Dive Into Caching With Go")

	intro := `
	Caching is a critical optimization strategy for improving application performance, reducing latency, and minimizing expensive 
	operations like database queries or API calls. In this lesson, we’ll explore caching concepts, tools, and best practices in Go, 
	something mid-level developers must understand to build efficient, scalable systems.
	`
	fmt.Print(intro)

	objectives := `
	Objectives 

1. Understand caching and its benefits.
2. Learn about common caching strategies.
3. Implement in-memory caching using Go’s sync.Map and third-party libraries like GoCache and Redis.
4. Explore caching best practices and pitfalls.
	`
	fmt.Println(objectives)

	whatIsCaching := `
	What Is Caching

Caching temporarily stores frequently accessed data in memory or disk for faster access. Instead of fetching the same data repeatedly 
from a slow source like a database, cache stores the result of the first fetch and reuses it.
	`
	fmt.Println(whatIsCaching)

	benefitsOfCaching := `
	Benefits of Caching

- Performance: Reduce load times for expensive operations.
- Scalability: Minimize strain on databases and external services.
- Cost Savings: Fewer expensive calls to external APIs or database queries.
	`
	fmt.Println(benefitsOfCaching)

	commonUseCases := `
	Common Use Cases

- Storing the results of API calls.
- Caching database query results.
- Caching static assets like configuration files or templates.
	`
	fmt.Println(commonUseCases)

	cachingStrategies := `
	Caching Strategies 

1. Write-Through:
	- Data is written to the cache and the underlying datastore simultaneously.
	- Ensures cache and database are always in sync but has higher latency on writes.

2. Write-Back:
	- Data is written only to the cache initially and flushed to the database periodically.
	- Improves write performance but risks data loss during cache eviction or server crashes.

3. Read-Through:
	- Cache is checked before accessing the datastore. If data is not in cache, it is fetched from the database and added to the cache.

4. Time-Based Expiry:
	- Cached data is valid for a set duration and expires afterward.
	`
	fmt.Println(cachingStrategies)

	inMemoryCaching()

	goCache()

	redisCaching()

	bestPractices := `
	Best Practices for Caching

1. Use TTL (Time-To-Live):
	- Set expiration times to prevent stale data and reduce memory usage.

2. Invalidate Stale Data:
	- Have strategies to invalidate or refresh outdated cached values.

3. Cache Only When Necessary:
	- Avoid caching data that changes frequently or is already fast to retrieve.

4. Use Appropriate Cache Size:
	- Prevent excessive memory usage by limiting cache size.

5. Monitor Cache Performance:
	- Use tools like Prometheus to track cache hit/miss rates.
	`
	fmt.Println(bestPractices)

	commonPitfalls := `
	Avoiding Common Pitfalls 

1. Over-Caching:
	- Too much caching can lead to excessive memory usage or stale data.

2. Cache Inconsistency:
	- Ensure synchronization between cache and datastore.

3. Cache Stampede:
	- Prevent multiple processes from fetching the same data when the cache expires. Use techniques like request coalescing.

4. Security Risks:
	- Don’t cache sensitive data in unsecured locations.	
	`
	fmt.Println(commonPitfalls)

	advancedTopics := `
	Advanced Topics

Distributed Caching:
	- Use distributed systems like Redis or Memcached for large-scale caching.

Lazy Loading:
	- Only cache data when it is first requested.

Write-Around Caching:
	- Write directly to the database and bypass the cache.
	`
	fmt.Println(advancedTopics)

	summary := `
	CONCLUSION 

Caching is a vital technique for building high-performance applications. This lesson covered:

	- Caching strategies like read-through, write-through, and time-based expiry.
	- Practical implementations using sync.Map, GoCache, and Redis.
	- Best practices for efficient and secure caching.
	- How to avoid common pitfalls like over-caching and cache stampede.
	`
	fmt.Println(summary)
}
