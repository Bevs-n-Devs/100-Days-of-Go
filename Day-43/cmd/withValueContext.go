package main

import (
	"context"
	"fmt"
)

func processingRequest(ctx context.Context) {
	userID := ctx.Value("userID")
	if userID == nil {
		fmt.Println("No userID provided!")
	}
	fmt.Println("Processing request for userID:", userID)
}

func withValueContext() {
	fmt.Print("\nPractical Examples: Using context.WithValue")

	ctx := context.WithValue(context.Background(), "userID", 42)

	processingRequest(ctx)
}
