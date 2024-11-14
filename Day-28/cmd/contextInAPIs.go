package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

/*
Contexts are often used in HTTP servers to set timeouts for client requests.
Here’s an example of how context can be used in an HTTP server.

- Setting Timeout for a Request: The dataHandler function sets a timeout of 3 seconds on the request context.
- Simulating Data Retrieval: The fetchData function uses a select to either wait for data
  retrieval (time.After(2 * time.Second)) or listen for the context to cancel (ctx.Done()).
- Handling Timeout Errors: If fetchData doesn’t complete within the context’s deadline, it returns a timeout error.
  The handler then returns an HTTP 408 (Request Timeout) to the client.
*/

func fetchData(ctx context.Context) (string, error) {
	select {
	// simulate data retrieval time
	case <-time.After(2 * time.Second):
		return "Data recieved successfully", nil
	// context cancellation or timeout
	case <-ctx.Done():
		return "", ctx.Err()
	}
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	// create a context with a 3-second timeout for this request
	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	// simulate data processing
	result, err := fetchData(ctx)
	if err != nil {
		http.Error(w, "Request timed out", http.StatusRequestTimeout)
		return
	}
	fmt.Fprintf(w, "Request: %s", result)
}

func contextInAPIs() {
	fmt.Println("\nContext With APIs")

	http.HandleFunc("/data", dataHandler)
	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
