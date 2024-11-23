package main

import "fmt"

func pointerToPointer() {
	fmt.Println("\nPointer to Pointer")

	info :=
		`- A pointer can point to another pointer. This is useful in scenarios where you need to dynamically manage references. 
	`
	fmt.Println(info)

	x := 42

	p := &x  // p points to x
	pp := &p // pp points to p

	fmt.Println("Get value of x via p pointer:", *p)
	fmt.Println("Get memory address of p via pp pointer:", *pp) // this gives you the address of p
	fmt.Println("Value of x via pp pointer:", **pp)             // dereference pp twice to get the value of x
}
