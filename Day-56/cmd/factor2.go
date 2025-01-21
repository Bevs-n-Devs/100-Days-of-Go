package main

import "fmt"

func Factor2() {
	var factor2 = `
	2. Dependencies: Explicitly Declare and Isolate Dependencies
		- Principle: Dependencies must be declared explicitly (e.g., in go.mod) and isolated.
		- Why: Prevents "it works on my machine" issues.
		
		Example:
			- Use go.mod to manage dependencies.
		
		Code Example:
			go mod init your-project-name
			go get guthub.com/gin-gonic/gin
	`
	fmt.Println(factor2)
}
