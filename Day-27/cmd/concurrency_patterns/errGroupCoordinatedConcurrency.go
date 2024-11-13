package concurrencypatterns

import (
	"fmt"
	"net/http"

	"golang.org/x/sync/errgroup"
)

/*
The errgroup package simplifies concurrent error handling and coordination.
*/

func ErrgroupCoordinatedConcurrency() {
	fmt.Println("\nConcurrency Patterns: Errgroup For Coordinated Concurrency")

	g := new(errgroup.Group)
	urls := []string{"https://golang.org", "https://google.com", "https://example.com"}

	for _, url := range urls {
		url := url
		g.Go(func() error {
			resp, err := http.Get(url)
			if err != nil {
				return err
			}
			fmt.Println(resp.Status)
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		fmt.Println("Encountered an error:", err)
	} else {
		fmt.Println("All requests succeeded")
	}

}
