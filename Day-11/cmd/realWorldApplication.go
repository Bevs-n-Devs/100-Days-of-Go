package main

import (
	"fmt"
	"net/http"
	"time"
)

/*
The context package is essential in web servers where requests may need to time out or be canceled.
For example, you can pass a context from the HTTP request handler to various functions to ensure
everything cancels if the request is terminated by the client.

The handler uses r.Context() to get the context associated with the HTTP request.
If the client disconnects or the request times out, ctx.Done() cancels the handler, and the server responds with a 408 status (Request Timeout).
*/
func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Println("Request started")

	select {
	case <-time.After(2 * time.Second): // simulate work
		fmt.Fprint(w, "Request completed successfully")
	case <-ctx.Done():
		fmt.Println("Request canceled by client:", ctx.Err())
		http.Error(w, "Request canceled", http.StatusRequestTimeout)
	}
}

func realWorldApplication() {
	fmt.Println("\nReal-World Application: HTTP Server with Context")

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
