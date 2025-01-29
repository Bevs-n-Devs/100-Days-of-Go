package main

import (
	"fmt"
	"log"
)

func main() {
	log.Println("Hello world, hello Yaw!")
	fmt.Println("\nError Handling")

	var description = `
Error handling in go is one of its most defining features. Unlike other langages that use
exceptions, Go employs a simple, explicit, and predictable approach to error handling using 
the error type. This design encourages developers to handle errors immediately, making the code
more robust and easier to debug.
	`
	fmt.Println(description)

	var whyExplicitErrorHandling = `
	Why Explicit Error Handling?
		1. Clarity: Errors are explicitly returned and handled where they occur.
		2. Predictability: No hidden exceptions or stack unwinding; the control flow is straightforward.
		3. Safety: Forces developers to think about and handle errors, reducing the likelihood of unexpected crashes.
	`
	fmt.Println(whyExplicitErrorHandling)

	var errorType = `
	The error Type
	In Go, the error type is a built-in interface that represents an error condition.

	Definition:
	type error interface {
		Error() string
	}
	
	An error is simply a value that implements the Error method and returns a descriptive error message.
	`
	fmt.Println(errorType)

	fmt.Println("\nCreating Errors")

	var creatingErrors1 = `
	1. Using error.New
		- The function creates a new error with a string message
	
	Code Example:
	import "errors"
	func Example() error {
		return errors.New("something went wrong")
	}
	`
	fmt.Println(creatingErrors1)

	var creatingErrors2 = `
	2. Using fmt.Errorf
		- Allows formatted error messages
	
	Code Example:
	import "fmt"
	func Example() error {
		return fmt.Errorf("something went wrong: %w", err)
	}
	`
	fmt.Println(creatingErrors2)

	var creatingErrors3 = `
	3. Using errors.Join
		-  Joins multiple errors into one
	
	Code Example:
	import "errors"
	func Example() error {
		err1 := errors.New("file not found")
		err2 := errors.New("permission denied")
		return errors.Join(err1, err2)
	}
	`
	fmt.Println(creatingErrors3)

	var handlingErrors = `
Handling Errors
	Errors in Go are typically handled using if statements. This ensures they are dealth with explicitly.

	Code Example:
	import (
		"errors"
		"fmt"
	)

	func Divide(a, b int) (int, error) {
		if b == 0 {
			return 0, errors.New("division by zero")
		}
		return a / b, nil
	}

	func main() {
		result, err := Divide(10, 0)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Result:", result)
	}
	`
	fmt.Println(handlingErrors)

	var idiomaticErrorHandling = `
Idiomatic Go Error Handling
	1. Handle Errors Where They Occur
		- Errors should be handled as close to their origin as possible. Avoid propagating errors (unless they provide additional context).
	2. Return Early on Errors
		- To prevent deeply nested code, handle errors and return early.

	Code Example:
	func ReadFile(filename string) ([]byte, error) {
		file, err := os.Open(filename)
		if err != nil {
			return nil, fmt.Errorf("failed to open file: %w", err)
		}
		defer file.Close()

		data, err := io.ReadAll(file)
		if err != nil {
			return nil, fmt.Errorf("failed to read file: %w", err)
		}

		return data, nil
	}

	3. Wrap Errors with Context
		- Add context to errors to make debugging easier.
	
	Code Example:
		func ProcessData() error {
		data, err := ReadFile("data.txt")
		if err != nil {
			return fmt.Errorf("process failed: %w", err)
		}
		fmt.Println("Data:", string(data))
		return nil
	}
	`
	fmt.Println(idiomaticErrorHandling)

	var customErrorTypes = `
	Sometimes you may want to define custom error types for specific use cases.

	Code Example:
	type ValidationError struct {
		Field   string
		Message string
	}

	func (e *ValidationError) Error() string {
		return fmt.Sprintf("validation error: field '%s' - %s", e.Field, e.Message)
	}

	func Validate(value int) error {
		if value < 0 {
			return &ValidationError{Field: "value", Message: "must be non-negative"}
		}
		return nil
	}

	func main() {
		err := Validate(-1)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
	type ValidationError struct {
		Field   string
		Message string
	}

	func (e *ValidationError) Error() string {
		return fmt.Sprintf("validation error: field '%s' - %s", e.Field, e.Message)
	}

	func Validate(value int) error {
		if value < 0 {
			return &ValidationError{Field: "value", Message: "must be non-negative"}
		}
		return nil
	}

	func main() {
		err := Validate(-1)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}

	`
	fmt.Println(customErrorTypes)

	var sentinelErrors = `
Sentinel Errors
	Sentinel errors are pre-declared errors that act as constants and can be compared using ==.

	Code Example:
	var ErrNotFound = errors.New("not found")

	func FindItem(id int) error {
		if id != 42 {
			return ErrNotFound
		}
		return nil
	}

	func main() {
		err := FindItem(1)
		if errors.Is(err, ErrNotFound) {
			fmt.Println("Item not found")
		}
	}
	`
	fmt.Println(sentinelErrors)

	fmt.Println("\nChecking and Unwrapping Errors")

	var errorsIs = `
	1. errors.Is
		- Used to check if an error matches a specific error, including wrapped errors.
	
	Cdoe Example:
	if errors.Is(err, ErrNotFound) {
		fmt.Println("Item not found")
	}
	`
	fmt.Println(errorsIs)

	var errorsAs = `
	2. errors.As (type assertion)
		- Used to check if an error is of a specific type.
	
	Code Example:
	var vErr *ValidationError
	if errors.As(err, &vErr) {
		fmt.Printf("Validation failed: %s\n", vErr.Message)
	}
	`
	fmt.Println(errorsAs)

	var errorsWrap = `
	3, errors.Wrap (unwrapping nested errors)
		- Extracts the underlying error from a wrapped error.
	
	Code Example:
	var vErr *ValidationError
	if errors.As(err, &vErr) {
		fmt.Printf("Validation failed: %s\n", vErr.Message)
	}
	`
	fmt.Println(errorsWrap)

	fmt.Println("\nCommon Go Error Handling Practices")

	var dontIgnoreErrors = `
	1. Don't Ignore Errors
		- Never ignore errors by writing _ or discarding them.

	Code Example:
	_, err := os.Open("file.txt") // BAD: Ignoring error
	if err != nil {
		fmt.Println("Error:", err)
	}
	`
	fmt.Println(dontIgnoreErrors)

	var avoidPanics = `
	2. Avoid Panics for Normal Errors
		- Use panic only for truly exceptional cases (e.g., unrecoverable errors).

	Code Example:
	func ReadFile(filename string) ([]byte, error) {
		file, err := os.Open(filename)
		if err != nil {
			return nil, err // Do not panic here
		}
		defer file.Close()
		return io.ReadAll(file)
	}
	`
	fmt.Println(avoidPanics)

	fmt.Println("\nBest Practices")

	var returnError = `
	1. Return error as the last value in a function

	Code Example:
	func Foo() (int, error) { ... }
	`
	fmt.Println(returnError)

	var addContext = `
	2. Add context to errors when propagating them.

	Code Example:
	return fmt.Errorf("additional context: %w", err)
	`
	fmt.Println(addContext)

	var moreBestPractices = `
	3. Use errors.Is and errors.As for checking and handling specific errors.
	4. Handle errors immediately and return early to avoid deep nesting.
	5. Avoid panic unless necessary.
	6. Use custom error types when more information is needed.
	`
	fmt.Println(moreBestPractices)

	var conclusion = `
Go’s error handling may feel verbose compared to languages with exception-based mechanisms, but its 
simplicity and explicit nature lead to robust and maintainable code. By following these principles 
and patterns, you can ensure that your Go applications handle errors predictably and effectively.
	`
	fmt.Println(conclusion)
}
