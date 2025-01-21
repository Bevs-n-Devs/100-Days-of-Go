package main

import "fmt"

func Factor5() {
	var factor5 = `
	5. Build, Release, Run: Strictly Separate Build and Run Stages
		- Principle: Application should have distinct stages for building, releasing and running.
			- Build: Compile the app and dependencies.
			- Release: Combine build artifacts and confihuration
			- Run: Execute the app in the runtime environment
		- Why: Ensure consistency between deployments

		Example in CI/CD:
			- Build and release with Docker
		
		Code Example:
			docker build -t your-app:latest .
			docker run --env-file .env your-app:latest
	`
	fmt.Println(factor5)
}
