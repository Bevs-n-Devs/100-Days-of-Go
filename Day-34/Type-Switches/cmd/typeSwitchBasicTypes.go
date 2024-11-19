package main

import "fmt"

func process(value interface{}) {
	switch v := value.(type) {
	case string:
		fmt.Printf("String: %s\n", v)
	case int:
		fmt.Printf("Integer: %d\n", v)
	case bool:
		fmt.Printf("Boolean: %t\n", v)
	default:
		fmt.Println("Unknown type")
	}
}

func typeSwitchWithBasicTypes() {
	fmt.Println("\nType Swicthes - With Baics Types")

	process("hello")
	process(42)
	process(true)
	process(3.14)
}
