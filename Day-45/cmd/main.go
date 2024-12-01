package main

import "fmt"

func main() {
	fmt.Println("Hello world, hello Yaw!")

	fmt.Println("\nUnderstanding and Implementing Middleware in Go")

	intro := `
Middleware is a foundational concept for building scalable, modular, and reusable web applications. 
It allows you to perform operations on HTTP requests and responses before and after reaching the main 
application logic. As a mid-level developer, mastering middleware patterns and their use cases will 
elevate your ability to build robust web servers and APIs.
	`
	fmt.Println(intro)

	whatIsMiddleware := `
	What is middleware in Go?

Middleware is a function or a chain of functions that intercept and process HTTP requests/responses. Middleware can:

	- Modify requests (e.g., adding headers).
	- Perform actions such as logging, authentication, or input validation.
	- Short-circuit requests (e.g., deny access based on certain criteria).
	`
	fmt.Println(whatIsMiddleware)

	middlewareInGo := `
	Middleware in Go

Middleware can be implemented in Go by wrapping HTTP handlers. You can use the standard net/http package or frameworks like Gin, Echo, or Chi.
	`
	fmt.Println(middlewareInGo)

	// uncomment to see basic logging middleware
	// basicLoggingMiddleware()

	// uncomment to see authentication middleware
	// customMiddlewareForAuthentication()

	keyConcepts := `
	Key Concepts in Middleware

1. Chainable Middleware:	
Middleware can be chained together, allowing multiple pieces of logic to process a request.

	func ChainMiddleware(middlewares ...func(http.Handler) http.Handler) func(http.Handler) http.Handler {
		return func(final http.Handler) http.Handler {
			for i := len(middlewares) - 1; i >= 0; i-- {
				final = middlewares[i](final)
			}
			return final
		}
	}

	mux := http.NewServeMux()
	mux.Handle("/", ChainMiddleware(LoggingMiddleware, AuthMiddleware)(http.HandlerFunc(mainHandler)))

2. Short-Circuiting Requests
Middleware can terminate a request early without calling the next handler.

3. Context Passing
Middleware can attach data to the request context to share information with downstream handlers.

	func ContextMiddleware(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), "user", "John Doe")
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
	`
	fmt.Println(keyConcepts)

	bestPractices := `
	Best Practices

1. Keep Middleware Small and Focused:
	- Each middleware should perform a single responsibility, such as logging, authentication, or rate limiting.
2. Avoid Blocking Operations:
	- Middleware should not block for extended periods (e.g., long database queries).
3. Use Context for Shared State:
	- Pass data between middleware and handlers using the context.Context.
4. Log Errors Explicitly:
	- Middleware should log errors when they occur, as it often runs before the main handler.
5. Chain Middleware Carefully:
- Order matters. For example, an authentication middleware should run before a logging middleware to avoid logging sensitive information.
	`
	fmt.Println(bestPractices)

	advancedMiddleware := `
	Advanced Middleware Use Caes 

1. Rate Limiting:
	- Throttle the number of requests from a client using a token bucket or leaky bucket algorithm.

	func RateLimitingMiddleware(next http.Handler) http.Handler {
		rateLimiter := time.Tick(100 * time.Millisecond) // Allow 10 requests/second
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			select {
			case <-rateLimiter:
				next.ServeHTTP(w, r)
			default:
				http.Error(w, "Too many requests", http.StatusTooManyRequests)
			}
		})
	}

2. Panic Recovery:
	- Handle panics in downstream handlers and prevent server crashes.

	func RecoveryMiddleware(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				}
			}()
			next.ServeHTTP(w, r)
		})
	}

3. Request Tracing:
	- Add unique request IDs to logs and headers for debugging and monitoring.
	
	func RequestIDMiddleware(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := uuid.New().String()
			w.Header().Set("X-Request-ID", requestID)
			log.Printf("Request ID: %s", requestID)
			next.ServeHTTP(w, r)
		})
	}
	`
	fmt.Println(advancedMiddleware)

	conclusion := `
	CONCLUSION

Middleware is an indispensable tool for building robust web applications in Go. 
Whether you're adding logging, authentication, rate limiting, or request tracing, 
understanding middleware patterns will make your applications more maintainable, 
scalable, and testable.

For your next project, implement middleware for logging, authentication, and error 
handling to gain hands-on experience.
	`
	fmt.Println(conclusion)
}
