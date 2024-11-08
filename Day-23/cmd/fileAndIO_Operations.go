package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

/*
The bufio package wraps os.File and other I/O sources, enabling buffered reading and writing.
This can be particularly beneficial for files and network streams where reading or writing one byte at a time is inefficient.

Using bufio.NewScanner, we read the file line-by-line, which is efficient for large files and uses minimal memory.
*/

// reading a file line-by-line
func efficientFileAndIOoperations() {
	fmt.Println("\nEfficient File and I/O Operations")

	file, err := os.Open("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text()) // print each line
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
