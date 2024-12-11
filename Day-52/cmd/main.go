package main

import "fmt"

func main() {
	fmt.Println("Hello world, hello Yaw!")

	fmt.Println("\n	DRY (Don't Repeat Yourself) Principles in SWD & Go")

	intro := `
The DRY (Don’t Repeat Yourself) principle is one of the most fundamental guidelines in software development. 
It emphasizes reducing code duplication to improve maintainability, consistency, and clarity.
	`
	fmt.Println(intro)

	whatIsDRY := `
	What is DRY?

The DRY principle states:

	"Every piece of knowledge must have a single, unambiguous, and authoritative representation within a system."

This means that you should avoid duplicating logic, structures, or configurations. If something needs to change, it should be updated in only one place.
	`
	fmt.Println(whatIsDRY)

	whyDRY := `
	Whis is DRY Important?

1. Consistency: Reducing redundancy ensures a single source of truth for any logic or functionality.
2. Ease of Maintenance: Changes can be made in one place without worrying about missing updates in duplicates.
3. Reduced Bugs: Avoiding duplication reduces the likelihood of introducing errors when changes are made.
4. Cleaner Codebase: Helps keep the codebase organized and easier to understand.
	`
	fmt.Println(whyDRY)

	duplicationCauses := `
	Common Causes of Code Duplication

1. Copy-Paste Programming: Reusing code without abstraction.
2. Lack of Modularization: Failing to split logic into reusable components.
3. Overlooked Shared Functionality: Different parts of the application perform similar tasks but are implemented separately.
	`
	fmt.Println(duplicationCauses)

	fmt.Println("\n	Strategies to Avoid Code Duplication")

	functionsAndMethods := `
	1. Use Functions and Methods

Encapsulate repeated logic in functions or methods.

Example (Without DRY):

	func CalculateAreaRectangle(length, width float64) float64 {
		return length * width
	}

	func CalculateAreaSquare(side float64) float64 {
		return side * side
	}

Example (With DRY):

	func CalculateArea(shape string, dimensions ...float64) float64 {
		switch shape {
		case "rectangle":
			return dimensions[0] * dimensions[1]
		case "square":
			return dimensions[0] * dimensions[0]
		default:
			return 0
		}
	}
	`
	fmt.Println(functionsAndMethods)

	costantsAndConfigurations := `
	2. Use Constants and Configurations

Define constants or configuration files for values used in multiple places.

Example (Without DRY):

	const taxRate = 0.15

	func CalculateTax(price float64) float64 {
		return price * 0.15 // Duplicate tax rate
	}

Example (With DRY):

	const taxRate = 0.15

	func CalculateTax(price float64) float64 {
		return price * taxRate
	}

	`
	fmt.Println(costantsAndConfigurations)

	leverageStructsAndInterfaces := `
	3. Leverage Structs and Interfaces

Use structs and interfaces to handle similar behaviors or entities.

Example (Without DRY):

	type Employee struct {
		Name string
		Age  int
	}

	type Manager struct {
		Name string
		Age  int
		TeamSize int
	}

Refactored (With DRY):

	type Person struct {
		Name string
		Age  int
	}

	type Manager struct {
		Person
		TeamSize int
	}
	`
	fmt.Println(leverageStructsAndInterfaces)

	usePackagesForReusableCode := `
	4. Use Packages for Reusable Code

Abstract shared code into packages to reuse across the project.

Here, each file has its own logging logic, leading to duplication.

Example (Without DRY): File 1: User Service

	package user

	import "log"

	func GetUser(id int) {
		log.Printf("[INFO] Fetching user with ID: %d\n", id)
		// Fetch user logic...
		log.Printf("[INFO] Successfully fetched user with ID: %d\n", id)
	}

Example (Without DRY): File 2: Product Service

	package product

	import "log"

	func GetProduct(id int) {
		log.Printf("[INFO] Fetching product with ID: %d\n", id)
		// Fetch product logic...
		log.Printf("[INFO] Successfully fetched product with ID: %d\n", id)
	}

Example (Without DRY): File 3: Order Service

	package order

	import "log"

	func GetOrder(id int) {
		log.Printf("[INFO] Fetching order with ID: %d\n", id)
		// Fetch order logic...
		log.Printf("[INFO] Successfully fetched order with ID: %d\n", id)
	}

In this example:
	- Each service implements its own logging logic.
	- If the logging format changes, you must update every occurrence manually, increasing the maintenance burden.
	- Duplication increases the risk of inconsistencies

Refactored (With DRY): Logger Package

We abstract the logging logic into a reusable logger package, following the DRY principle.

	package logger

	import "log"

	func LogMessage(level, message string) {
		log.Printf("[%s] %s\n", level, message)
	}

Refactored (With DRY): File 1: User Service

	package user

	import "myapp/logger"

	func GetUser(id int) {
		logger.LogMessage("INFO", "Fetching user with ID: "+fmt.Sprint(id))
		// Fetch user logic...
		logger.LogMessage("INFO", "Successfully fetched user with ID: "+fmt.Sprint(id))
	}

Refactored (With DRY): File 2: Product Service

	package product

	import "myapp/logger"

	func GetProduct(id int) {
		logger.LogMessage("INFO", "Fetching product with ID: "+fmt.Sprint(id))
		// Fetch product logic...
		logger.LogMessage("INFO", "Successfully fetched product with ID: "+fmt.Sprint(id))
	}

Refactored (With DRY): File 3: Order Service

	package order

	import "myapp/logger"

	func GetOrder(id int) {
		logger.LogMessage("INFO", "Fetching order with ID: "+fmt.Sprint(id))
		// Fetch order logic...
		logger.LogMessage("INFO", "Successfully fetched order with ID: "+fmt.Sprint(id))
	}
	`
	fmt.Println(usePackagesForReusableCode)

	templateReuseForWebApplications := `
	5. Template Reuse for Web Applications

For web apps, use templates to avoid repeating HTML/CSS structures.

Refactored (With DRY): Create base.html:

	<!DOCTYPE html>
	<html>
	<head>
		<title>{{ .Title }}</title>
	</head>
	<body>
		{{ block "content" . }}{{ end }}
	</body>
	</html>

Extend base.html for specific pages:

	{{ define "content" }}
	<h1>Welcome to the Homepage</h1>
	<p>This is the content of the homepage.</p>
	{{ end }}
	`
	fmt.Println(templateReuseForWebApplications)

	whenNotToUseDRY := `
	When Not to Overapply DRY

1. Over-Abstraction: Abstracting prematurely or excessively can make code harder to understand.
2. Different Contexts: Code may look similar but have different purposes; forcing DRY might obscure intent.
3. Simple Scripts: In small programs or one-off scripts, the overhead of DRY may not be worth it.	
	`
	fmt.Println(whenNotToUseDRY)

	benefits := `
	Benefits of DRY:

1. Reduced Maintenance Cost: Update logic in one place.
2. Consistency: Ensure uniform behavior across the application.
3. Improved Readability: Less clutter and repeated logic.
	`
	fmt.Println(benefits)

	conclusion := `
	CONCLUSION 

The DRY principle is a cornerstone of clean and maintainable code. However, it must be applied judiciously 
to avoid unnecessary complexity. By recognizing common patterns and abstracting them effectively, you can 
significantly enhance the quality of your codebase.
	`
	fmt.Println(conclusion)
}
