package main

import "fmt"

func main() {
	fmt.Println("Hello world, hello Yaw!")

	fmt.Println("\nCore Development Principles in Go")

	intro := `
As a mid-level developer, you’re transitioning from understanding basic concepts to tackling more complex software engineering challenges. 
This lesson will cover key principles, tools, and practices that elevate your skills and prepare you for senior-level responsibilities.	
	`
	fmt.Println(intro)

	keyPrinciples := `
	Core Development Principles:

1. SOLID Principles:
	- Single Responsibility: Each module/class should have one responsibility.
	- Open/Closed: Open for extension, closed for modification.
	- Liskov Substitution: Subtypes should be substitutable for their base types.
	- Interface Segregation: Prefer smaller, specific interfaces over large, general ones.
	- Dependency Inversion: Depend on abstractions, not concrete implementations.
2. DRY (Don’t Repeat Yourself): Avoid duplicate code by abstracting reusable logic into functions, modules, or services.
3. YAGNI (You Aren’t Gonna Need It): Don’t implement features until they’re absolutely necessary.
4. KISS (Keep It Simple, Stupid): Simplicity in design and code increases maintainability and reduces bugs.
	`
	fmt.Println(keyPrinciples)

	solidAntipatterns := `
	Anti-Pattern:

	type User struct {
		Name  string
		Email string
	}

	func (u User) Notify(emailService EmailService) {
		emailService.SendEmail(u.Email, "Notification")
	}
Here, the User struct has too many responsibilities (violating Single Responsibility Principle).
	`
	fmt.Println(solidAntipatterns)

	refactored := `
	Refactored:

	type Notifier interface {
		Notify(email string)
	}

	type EmailNotifier struct {
		EmailService EmailService
	}

	func (n EmailNotifier) Notify(email string) {
		n.EmailService.SendEmail(email, "Notification")
	}
	`
	fmt.Println(refactored)

	debuggingTools := `
	Debugging Tools:

1. dlv (Delve): A debugger for Go.
	- Set breakpoints and inspect variables.

	dlv debug main.go

2. pprof: A tool for profiling CPU and memory usage.
	- Add net/http/pprof to expose profiling endpoints.

	import _ "net/http/pprof"

Profiling Example
	go test -bench . -cpuprofile=cpu.prof
	go tool pprof cpu.prof
Use commands like top and list to analyze bottlenecks.
	`
	fmt.Println(debuggingTools)

	interfaceModularDesign := `
	Use Interfaces for Modular Design

Interfaces decouple implementation from usage, making your application more flexible and testable.

	type Storage interface {
		Save(data string) error
	}

	type FileStorage struct{}

	func (fs FileStorage) Save(data string) error {
		// Save to a file
		return nil
	}
	`
	fmt.Println(interfaceModularDesign)

	productionReadyPractices := `
	Continuous Integration/Continuous Deployment (CI/CD)

- Use CI tools like GitHub Actions, GitLab CI, or CircleCI for automated testing and deployment.

	Monitoring and Logging

- Integrate Prometheus for metrics collection and Grafana for visualization.
- Use Zap or Logrus for structured logging.
	`
	fmt.Println(productionReadyPractices)

	collaboration := `
	Collaboration and Teamwork

1. Code Reviews:
	- Provide actionable feedback.
	- Focus on logic, readability, and test coverage.
2. Pair Programming:
	- Collaborate with peers to solve complex issues.
3. Effective Communication:
	- Use concise, clear language in PRs, documentation, and discussions.
	`
	fmt.Println(collaboration)

	advancedTopics := `
	Advanced Topics

To stay competitive, explore advanced concepts like:

	- Event-Driven Architecture: Learn about message brokers like Kafka or RabbitMQ.
	- Design Patterns: Study patterns like Singleton, Factory, and Observer.
	- Cloud Services: Familiarize yourself with AWS, GCP, or Azure.
	`
	fmt.Println(advancedTopics)

	summary := `
	CONCLUSION 
	
- Adhere to core principles like SOLID, DRY, and YAGNI.
- Master debugging and profiling tools.
- Organize code for scalability and maintainability.
- Implement CI/CD and logging for production readiness.
- Collaborate effectively and mentor juniors.
- Continuously learn advanced topics to stay competitive.
	`
	fmt.Println(summary)
}
