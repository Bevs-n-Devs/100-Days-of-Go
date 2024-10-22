package main

// A pointer is a variable that holds the memory address of another variable.
// Instead of holding a value directly, it holds a reference to where the value is stored in memory.

import "fmt"

func declaringPointers() {
	fmt.Println("\nDeclaring and Using Pointers")
	num := 10 // create a normal interger value
	p := &num // get the memory address of num and store it in pointer p

	fmt.Println("value num =", num)
	fmt.Println("memory address of p =", p)
	fmt.Println("*p =", *p) // 'dereference' p (get the value of the memory address that p points to - num)

	// we can change the value at the memory address p points to
	*p = 20
	fmt.Println("num afetr *p = 20:", num) // output of the new value of num
	fmt.Println("*p and num are the same value:", *p, num, "- this is because they point to the same memory address!")
	fmt.Println("p and &num have the same memory address:", p, &num)
}

func declaringPointers2() {
	fmt.Println("\nDeclairing Pointers 2")
	i, j := 42, 2701

	p := &i
	fmt.Println(*p)
	*p = 21 // change the value of i through the pointer
	fmt.Println(i)

	p = &j       // now p points to int variable j
	*p = *p / 37 // divide j through the pointer
	fmt.Println(j)
}

// Go uses pass-by-value, meaning that when you pass a variable to a function, a copy is made.
// However, with pointers, you can modify the original value because you're passing the memory address.
func changeValue(p *int) {
	*p = 42
}

// Pointers can have two major benefits:
// 1. Passing large data structures (e.g., arrays, structs) by value can be expensive in terms of memory and time.
// By passing a pointer, you avoid making a copy of the data.
// 2. If you want a function to modify its arguments, you need to pass pointers to those arguments.
func swap(a, b *int) {
	*a, *b = *b, *a // dereference pointers to swap values
}

// Pointers are ver useful when working with structs.
// Instead of passing a large struct by value (which copies the struct), you can pass a pointer
type Person struct {
	name string
	age  int
}

func updatePerson(p *Person) {
	p.name = "Updated Name" // modify the struct through the pointer
}

// Slices & Maps behave like reference types, meaning they automatically point to the underlying data.
// You don't need to use pointers explicitly when passing slices or maps to functions.
func modifySlice(s []int) {
	s[0] = 100 // modifies the orginal slice
}

// task 2
func doubleValue(n *int) {
	*n = *n * 2
}

// task 3
type Book struct {
	title  string
	author string
}

func updateBook(b *Book) {
	b.title = "New Title"
}

func main() {
	fmt.Println("Hello world, hello Yaw!")
	fmt.Println("Pointers and Memory Management")

	declaringPointers()

	declaringPointers2()

	// pointers in functions
	num := 10
	fmt.Println("\nPointers in Functions")
	fmt.Println("Before changed value:", num)

	changeValue(&num) // pass the memory address of num variable
	fmt.Println("After changed value:", num)

	// when we declare a pointer without initialization its value is nil
	// this means the pointer is not pointing to any memory location
	fmt.Println("\nZero Value of Pointers")
	var pt *int
	fmt.Println(pt)

	if pt == nil {
		fmt.Println("pt is nil")
	}

	fmt.Println("\nPointers vs Values")
	num1, num2 := 10, 20
	fmt.Println("Bfore swap: num1 =", num1, ", num2 =", num2)

	swap(&num1, &num2) // pass the memory address of num1 & num2
	fmt.Println("After swap: num1 = ", num1, ", num2 = ", num2)

	fmt.Println("\nPointers and Structs")
	person := Person{name: "Yaw", age: 38}
	fmt.Println("Before update:", person)

	// we pass the memory address of person using &person,
	// allowing the function to directly modify the original struct

	updatePerson(&person)                // pass the pointer to the person struct
	fmt.Println("After update:", person) // the original struct has been modified

	fmt.Println("\nPointrers with Slices and Maps")
	nums := []int{1, 2, 3}
	fmt.Println("Before modifying slice:", nums)

	modifySlice(nums)                        // passing the slice by value, but it acts like a reference
	fmt.Println("Afetr modify slice:", nums) // the original slice modified

	fmt.Println("\nDay 5 Tasks")

	fmt.Println("\nTask 1 - modify int variable")
	myAge := 37
	fmt.Println("Current age:", myAge)
	age := &myAge

	*age = 38
	println("New age:", myAge)

	fmt.Println("\nTask 2 - modify int variable via function")
	originalNumber := 10
	fmt.Println("Original number:", originalNumber)
	doubleValue(&originalNumber)
	fmt.Println("Doubled number:", originalNumber)

	fmt.Println("\nTask 3 - modify struct")
	myBook := Book{title: "100 Days of Go", author: "Yaw Akoto"}
	fmt.Println("My book before update:", myBook)
	updateBook(&myBook)
	fmt.Println("My book after update:", myBook)

	fmt.Println("\nTask 4 - pointer safety")
	var n *string

	if n == nil {
		fmt.Println("This pointer n is empty:", n)
		y := "yAW"
		fmt.Println("new var y:", y)
		n = &y
		*n = "Yaw Akoto"
		fmt.Println("This pointer n is no longer empty:", *n)
		fmt.Println("Updated var y:", y)
	}

	fmt.Println("")
	todoListExercise()
}

// When to Use Pointers?
// Efficiency: If you are dealing with large structs or arrays, passing by value can be inefficient. Use pointers to avoid copying large amounts of data.
// Modifying Values: If you need to modify a variable inside a function, use pointers to pass a reference to the variable.
// Avoiding Nil Pointer Dereferencing: Always check that a pointer is non-nil before dereferencing it, or you will encounter a runtime panic.
