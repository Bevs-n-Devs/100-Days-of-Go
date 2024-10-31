package main

import (
	"fmt"
	"io/ioutil"
)

/*
The os and io/ioutil packages in Go provide the primary methods for working with files.

Use ioutil.ReadFile to read the entire content of a file.
os.Open and defer allow for more controlled file handling and efficient memory use.
Writing to a File:

Use ioutil.WriteFile to quickly write data to a file.
The os package offers more fine-grained control if you need it.

ioutil.WriteFile writes data to a file with the specified permissions.
ioutil.ReadFile reads the entire file content as bytes, which can be converted to a string if needed.
*/

func fileHandling() {
	fmt.Println("\nFile Handling With Go")

	// write a file
	data := []byte("Go file handling\nHello world, hello Yaw!")
	err := ioutil.WriteFile("hello.txt", data, 0644)

	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Println("Data written to file successfully.")

	// reading from a file
	fileData, err := ioutil.ReadFile("hello.txt")
	if err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}
	fmt.Println("\nData read from file:", string(fileData))
}
