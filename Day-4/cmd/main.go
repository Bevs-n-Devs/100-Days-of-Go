package main

import (
	"fmt"
	"os"
)

// simple function that prinst a message
func helloWorld() {
	fmt.Println("Hello world, hello Yaw!")
}

// function accepting arguements
func add(x int, y int) {
	fmt.Println("\nFunctions accepting arguments:")
	total := 0
	total = x + y
	fmt.Println(total)
}

// functions with return values
func add2(x int, y int) int {
	fmt.Println("\nFunctions return values:")
	total := 0
	total = x + y
	return total
}

// return values of a function can be named
func rectangle(height int, width int) (area int) {
	fmt.Println("\nFunctions with named return values:")
	parameter := 2 * (height + width)
	fmt.Println("Parameter: ", parameter)

	area = height * width

	// return statement without specifying the variable name
	return
}

// variadic functions (functions that accept a variable amount of arguments - the number of function inputs can vary)
func sum(nums ...int) int {
	fmt.Println("Variadic functions")
	total := 0

	for _, num := range nums {
		total += num
	}
	return total
}

// passing a function as an arguement
func applyDouble(f func(int) int, x int) int {
	fmt.Println("Passing a Function as an argument:")
	return f(x)
}

func double(n int) int {
	return n * 2
}

// closures
func counter() func() int {
	fmt.Println("Function closures:")
	count := 0
	return func() int {
		count++
		return count
	}
}

// multiple deters - this will execute the LAST deter 1st, then the FIRST deter last!
func multipleDeters() {
	fmt.Println("Using multiple deters")
	fmt.Println("Start of the function")

	// Deferred calls, which will be executed in reverse order
	defer fmt.Println("Deferred call 1")
	defer fmt.Println("Deferred call 2")
	defer fmt.Println("Deferred call 3")

	fmt.Println("End of the function")
}

func main() {
	fmt.Println("Functions with Go")

	helloWorld()

	add(50, 30)

	myAge := add2(21, 17)
	fmt.Printf("My age: %v\n", myAge)

	fmt.Println("Calculate Area: ", rectangle(20, 30))

	result := sum(1, 2)
	fmt.Println(result)
	result = sum(1, 2, 3)
	fmt.Println(result)

	nums := []int{1, 2, 3, 4}
	result = sum(nums...)
	print(result)

	// anonymous functions - defining a function wihtout having to name it
	fmt.Println("\nAnonymous functions:")
	greet := func(name string) {
		fmt.Printf("Hello %s\n", name)
	}

	greet("Yaw")

	increment := counter()
	fmt.Println(increment())
	fmt.Println(increment())

	double := applyDouble(double, 5)
	fmt.Println(double)

	// defer - keyword used to schedule a function call to be run AFTER the function completes.
	// This is often used for resource cleanup, such as closing files or rel;easing blocks
	file, err := os.Open("C:/Users/yawak/Desktop/The BOX/The VAULT/goVault/100 Days of Go/100-Days-of-Go/Day-4/cmd/employees.txt") // full file path is needed!
	if err != nil {
		fmt.Println("Error opening file", err)
		return
	}
	defer file.Close() // this will run when the main function exits

	fmt.Println("File opened successfully!")

	multipleDeters()

}
