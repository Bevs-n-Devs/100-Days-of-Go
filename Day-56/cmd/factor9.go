package main

import "fmt"

func Factor9() {
	var factor9 = `
	9. Disposability: Maximize Robustness with Fast Startup and Graceful Shutdown
		Principle: Processes should start quickly, handle termination gracefully, and release resources.
		Why: Supports rapid scaling and minimize downtime.

		Code Example:
			func main() {
				srv := &http.Server{Addr: ":8080"}
				
				go func() {
					if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
						log.Fatalf("ListenAndServe(): %v", err)
					}
				}()
				
				// Graceful shutdown
				quit := make(chan os.Signal, 1)
				signal.Notify(quit, os.Interrupt)
				<-quit
				log.Println("Shutting down server...")
				srv.Shutdown(context.Background())
			}
	`
	fmt.Println(factor9)
}
