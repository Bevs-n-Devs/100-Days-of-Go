package main

import (
	"fmt"
	"log"
)

func main() {
	log.Println("Hello world, hello Yaw!")
	fmt.Println("\nLaw of Demeter")

	var definition = `
The Law of Demeter (LoD), also known as the "Principle of Least Knowledge", is a design guideline
for software development that encourages modules or objects to only interact with their close collaborators (direct dependencies).

In simple terms:
	- "Only talk to your friends, not strangers."
	- Each unit (class, function, or module) should only call methods or access properties of:
		1. Itself
		2. Its own fields
		3. Its method parameters
		4. Objects it creates directly
	`
	fmt.Println(definition)

	var whyLoD = `
	Why the Law of Demeter Matters
		1. Reduces Coupling:
			- Prevents one module from depending on the internal details of another.
			- Makes the codebase more modular and easier to change.
		2. Improves Maintainability:
			- Changes in one part of the system won't ripple across multiple modules.
		3. Enhances Security:
			- Code becomes easier to understand, as interactions are limited to direct collaborators.
		4. Promotes Encapsulation:
			- Internal details of objects are hidden, and behavior is exposed through clear interfaces.
	`
	fmt.Println(whyLoD)

	var breakingLoD = `
	Breaking the Law of Demeter
		- When a module reaches into nested objects to access their fields or methods, it breaks LoD.
	
	Code Example:
	type Address struct {
		City string
	}

	type User struct {
		Name    string
		Address *Address
	}

	type Order struct {
		User *User
	}

	func GetOrderCity(order *Order) string {
		// Breaks the Law of Demeter: Accessing deeply nested fields
		return order.User.Address.City
	}
		- Here GetOrderCity depends on the entire chain of Order -> User -> Address, creating tight coupling. 
		  Any change in the structure of User or Address will affect GetOrderCity.
	`
	fmt.Println(breakingLoD)

	var followingLoD = `
	Following the Law of Demeter
		- The code above can be fixed by exposing only what’s necessary through proper methods

	Code Example:
	type Address struct {
		City string
	}

	type User struct {
		Name    string
		Address *Address
	}

	func (u *User) GetCity() string {
		return u.Address.City
	}

	type Order struct {
		User *User
	}

	func (o *Order) GetCity() string {
		return o.User.GetCity()
	}
		- GetOrderCity only interacts with the Order and User objects directly. It doesn't "reach through" to the Address struct.
	`
	fmt.Println(followingLoD)

	var coreRules = `
	Core Rules of the Law of Demeter
		1. A method M of a class C should only call:
			- Methods of C (itself)
			- Methods of objects passed to it as arguments
			- Methods of objects it creates
			- Method of its direct instance variables
		2. Avoid "method chaining" or "train wrecks" (object.a.b.c.d)
	`
	fmt.Println(coreRules)

	Example1()
	Example2()
	Example3()

	var benefits = `
	Benefits of Following the Law of Demeter
		1. Reduced Ripple Effect:
			- Changes in one part of the system have minimal impact on others.
		2. Easier Testing:
			- With fewer dependencies, units are simpler to test in isolation.
		3. Improved Readability:
			- The code is less cluttered with long chains of calls and is easier to understand.
		4. Encapsulation:
			- Prevents leaking implementation details of one module to another.
	`
	fmt.Println(benefits)

	var tipsToFollow = `
	Practical Tips to Follow
		1. Use Wrapper Methods:
			- Add helper methods to encapsulate nested interactions.
		2. Limit Object Exposure:
			- Avoid returning raw objects that expose internal state.
		3. Dependency Injection:
			- Pass required dependencies explicitly instead of creating them internally.
		4. Avoid Method Chaining:
			- Chainable methods may tempt you to violate LoD. Use them sparingly.
	`
	fmt.Println(tipsToFollow)

	var whenToBreak = `
	When to Break the Law of Demeter
		- Accessing well-defined getters/setters for a value.
		- Fluent interfaces, where chaining is intentional and improves readability (e.g., SQL builders or Go structs like http.Request).
	`
	fmt.Println(whenToBreak)

	var summary = `
The Law of Demeter encourages limiting interactions to direct collaborators to ireduce coupling
and imporve maintainability. While it requires some upfront design effort, it pays off by creating a
clean, modular and scalable codebase.

By adopting this principle, you'll build systems that are easier to understand, modify and test.
	`
	fmt.Println(summary)
}
