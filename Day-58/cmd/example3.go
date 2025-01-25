package main

import "fmt"

func Example3() {
	fmt.Println("\nExample 3: Dependency Injection")
	var scenario = `
	Following PoLA:
		- Passing dependencies explicitly through a constructor or method arguments adheres to LoD, as collaborators are defined directly.

	Code Example:
	type Logger struct{}

	func (l *Logger) Log(msg string) {
		fmt.Println(msg)
	}

	type Service struct {
		Logger *Logger
	}

	func (s *Service) DoSomething() {
		s.Logger.Log("Doing something")
	}
		- The Service only interacts with its collaborator, the Logger, and doesn't try to create it internally.
	`
	fmt.Println(scenario)
}
