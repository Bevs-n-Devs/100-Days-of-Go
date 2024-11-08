package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

/*
We use bufio.NewWriter to buffer writes to the file, reducing the number of actual I/O operations.
The writer.Flush() call ensures that all data in the buffer is written to the file before closing it.
*/

// writing to a file with Buffered Writer
func writingFilesWithBufferedWriter() {
	fmt.Println("\nWriting Files With Buffered Writer")

	file, err := os.Create("output.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	for i := 1; i <= 5; i++ {
		writer.WriteString(fmt.Sprintf("This is line %d\n", i))
	}

	// make sure data is written to the file
	writer.Flush()
	fmt.Println("Data written to output.txt")
}
