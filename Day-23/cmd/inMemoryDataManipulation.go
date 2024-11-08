package main

import (
	"bytes"
	"fmt"
)

/*
The bytes.Buffer package allows you to store data in memory and perform efficient read and write operations.
This is helpful when you need to construct or modify strings, and for intermediate data storage.

We use buffer.WriteString to add multiple strings to the buffer without creating new strings at each step.
This approach is memory-efficient compared to repeatedly concatenating strings.
*/

func inMemoryDataManipulation() {
	fmt.Println("\nIn-Memory Data Manipulation")

	var buffer bytes.Buffer

	// write stgrings into the buffer
	buffer.WriteString("Hello, ")
	buffer.WriteString("World!")
	buffer.WriteString("\n")

	// convert buffer contents to string
	fmt.Println(buffer.String())
}
