package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world, hello Yaw!")

	fmt.Println("\n	Mastering the SOLID Principles in Go")

	intro := `
The SOLID principles are foundational design guidelines for creating maintainable, scalable, and robust software.	
	`
	fmt.Println(intro)

	whatIsSOLID := `
	What is SOLID?

The acronym SOLID stands for:

	- S: Single Responsibility Principle
	- O: Open/Closed Principle
	- L: Liskov Substitution Principle
	- I: Interface Segregation Principle
	- D: Dependency Inversion Principle
	`
	fmt.Println(whatIsSOLID)

	singleResponsibilityPrinciple := `
	Single Responsibility Principle

A class (or struct in Go) should have only one reason to change, meaning it should only handle one responsibility.
	`
	fmt.Println(singleResponsibilityPrinciple)

	violationOfSRP := `
Let's say we are building an application that processes orders:

	Violation of SRP

	type OrderProcessor struct{}

	func (op *OrderProcessor) Process(order string) {
		fmt.Println("Processing order:", order)
		// Sends email notification (additional responsibility)
		op.sendEmail(order)
	}

	func (op *OrderProcessor) sendEmail(order string) {
		fmt.Println("Sending email for order:", order)
	}

Here, OrderProcessor handles two responsibilities: processing orders and sending emails. This violates SRP.
	`
	fmt.Println(violationOfSRP)

	adheringToSRP := `
	Addhering to SRP

	type OrderProcessor struct{}

	func (op *OrderProcessor) Process(order string) {
		fmt.Println("Processing order:", order)
	}

	type EmailNotifier struct{}

	func (en *EmailNotifier) SendEmail(order string) {
		fmt.Println("Sending email for order:", order)
	}

Now, OrderProcessor only processes orders, and EmailNotifier only sends emails.
	`
	fmt.Println(adheringToSRP)

	openClosedPrinciple := `
	Open/Closed Principle (OCP)

A class (or struct) should be open for extension but closed for modification.	
	`
	fmt.Println(openClosedPrinciple)

	violationOfOCP := `
Suppose we need to calculate discounts for different customers.

	Violation of OCP

	type DiscountCalculator struct{}

	func (dc *DiscountCalculator) Calculate(customerType string, price float64) float64 {
		if customerType == "regular" {
			return price * 0.9
		} else if customerType == "premium" {
			return price * 0.8
		}
		return price
	}

Adding a new customer type requires modifying the method, violating OCP.
	`
	fmt.Println(violationOfOCP)

	adheringToOCP := `
	Addhering to OCP

	type DiscountStrategy interface {
	Calculate(price float64) float64
	}

	type RegularCustomer struct{}

	func (rc RegularCustomer) Calculate(price float64) float64 {
		return price * 0.9
	}

	type PremiumCustomer struct{}

	func (pc PremiumCustomer) Calculate(price float64) float64 {
		return price * 0.8
	}

	func main() {
		var customer DiscountStrategy = PremiumCustomer{}
		fmt.Println("Discounted price:", customer.Calculate(100))
	}
	`
	fmt.Println(adheringToOCP)

	liskovSubstitutionPrinciple := `
	Liskov Substitution Principle (LSP)

Subtypes must be substitutable for their base types without altering the correctness of the program.
This means that if a base type is used where a subtype is expected, the program should behave as if 
it was using the subtype - without breaking the program or causing it to behave incorrectly.
	`
	fmt.Println(liskovSubstitutionPrinciple)

	violationOfLSP := `

Imagine a Bird interface with a Fly method.

	Violation of LSP

	type Bird interface {
		Fly()
	}

	type Sparrow struct{}

	func (s Sparrow) Fly() {
		fmt.Println("Sparrow is flying")
	}

	type Ostrich struct{}

	func (o Ostrich) Fly() {
		panic("Ostrich can't fly!")
	}

Calling Fly on Ostrich breaks the program. This violates LSP because Ostrich is not a proper subtype of Bird.
	`
	fmt.Println(violationOfLSP)

	adheringToLSP := `
	Addhering to LSP

	type Bird interface {
		Move()
	}

	type FlyingBird struct{}

	func (fb FlyingBird) Move() {
		fmt.Println("Flying bird is flying")
	}

	type WalkingBird struct{}

	func (wb WalkingBird) Move() {
		fmt.Println("Walking bird is walking")
	}

Now, all Bird types adhere to the contract without causing runtime errors.
	`
	fmt.Println(adheringToLSP)

	interfaceSegregationPrinciple := `
	Interface Segregation Principle (ISP)

Clients should not be forced to implement interfaces they don’t use.
(This means that clients should not be forced to depend on methods they don’t use).
	`
	fmt.Println(interfaceSegregationPrinciple)

	violationOfISP := `
Suppose we have an interface for different types of printers.

	Violation of ISP

	type Printer interface {
		Print()
		Scan()
		Fax()
	}

	type BasicPrinter struct{}

	func (bp BasicPrinter) Print() {
		fmt.Println("Printing...")
	}

	func (bp BasicPrinter) Scan() {
		panic("BasicPrinter doesn't support scanning!")
	}

	func (bp BasicPrinter) Fax() {
		panic("BasicPrinter doesn't support faxing!")
	}

Forcing BasicPrinter to implement methods it doesn’t support violates ISP.
	`
	fmt.Println(violationOfISP)

	adheringToISP := `
	Addhering to ISP

	type Printer interface {
		Print()
	}

	type Scanner interface {
		Scan()
	}

	type BasicPrinter struct{}

	func (bp BasicPrinter) Print() {
		fmt.Println("Printing...")
	}

	type AdvancedPrinter struct{}

	func (ap AdvancedPrinter) Print() {
		fmt.Println("Printing...")
	}

	func (ap AdvancedPrinter) Scan() {
		fmt.Println("Scanning...")
	}

Now, BasicPrinter and AdvancedPrinter only implement the interfaces they need.	
	`
	fmt.Println(adheringToISP)

	dependencyInversionPrinciple := `
	Dependency Inversion Principle (DIP)

High-level modules should not depend on low-level modules. Both should depend on abstractions.
	`
	fmt.Println(dependencyInversionPrinciple)

	violationOfDIP := `
Suppose we have a Database struct.

	Violation of DIP

	type Database struct{}

	func (db *Database) Save(data string) {
		fmt.Println("Saving data:", data)
	}

	type DataProcessor struct {
		db Database
	}

	func (dp *DataProcessor) Process(data string) {
		dp.db.Save(data)
	}

DataProcessor is tightly coupled to Database. Changing Database requires modifying DataProcessor.
	`
	fmt.Println(violationOfDIP)

	adheringToDIP := `
	Adhering to DIP

	type Database interface {
		Save(data string)
	}

	type MySQLDatabase struct{}

	func (db MySQLDatabase) Save(data string) {
		fmt.Println("Saving data in MySQL:", data)
	}

	type DataProcessor struct {
		db Database
	}

	func (dp *DataProcessor) Process(data string) {
		dp.db.Save(data)
	}

	func main() {
		db := MySQLDatabase{}
		dp := DataProcessor{db: db}
		dp.Process("Sample Data")
	}

Here, DataProcessor depends on the Database abstraction, making it flexible to use different database implementations.
	`
	fmt.Println(adheringToDIP)

	benefitsOfSOLID := `
	Benefits of SOLID:

1. Improved Maintainability: Easier to update and modify without introducing bugs.
2. Enhanced Testability: Code becomes more modular, making it easier to test individual components.
3. Scalability: Easy to extend functionality without disrupting existing code.
4. Readability: Clean, well-structured code is easier for developers to understand.
	`
	fmt.Println(benefitsOfSOLID)

	commonPitfalls := `
	Common Pitfalls:

- Over-engineering: Don’t apply SOLID principles to simple use cases unnecessarily.
- Misinterpreting principles: Ensure that your implementation solves a genuine problem, not just adheres to the principle superficially.	
	`
	fmt.Println(commonPitfalls)

	conclusion := `
	CONCLUSION

SOLID principles are essential for creating high-quality, robust software. When applied correctly, they enable developers to write flexible, maintainable, and scalable code.
	`
	fmt.Println(conclusion)
}
