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
}
