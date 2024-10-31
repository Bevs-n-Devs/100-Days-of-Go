package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

/*
The encoding/json package provides methods to work with JSON data.

Defining a Struct for JSON Data:

Define a Go struct that matches the JSON structure. This allows for easy parsing when you know the structure in advance.
Decoding JSON:

Use json.Unmarshal to decode JSON into a struct if the data is in a file or byte slice.
Encoding JSON:

Use json.Marshal to encode data from a struct into JSON format.

We define a User struct matching the JSON format.
json.MarshalIndent formats JSON with indentation, making it easier to read.
json.Unmarshal decodes the JSON data back into Go structs.
*/
type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func parsingJSON() {
	fmt.Println("\nParsing and Encoding JSON files")

	//JSON data to be written to a file
	users := []User{
		{Name: "John Doe", Age: 25, Email: "john_doe@email.com"},
		{Name: "Jane Doe", Age: 30, Email: "jane_doe@email.com"},
	}

	// encoding users slice to JSON and writing to a file
	file, err := json.MarshalIndent(users, "", " ")
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	err = ioutil.WriteFile("users.json", file, 0644)
	if err != nil {
		fmt.Println("Error writing to JSON file:", err)
		return
	}
	fmt.Println("JSON data written to file successfully.")

	// reading JSON data from the file and decoding
	jsonData, err := ioutil.ReadFile("users.json")
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}

	var decodeUsers []User
	err = json.Unmarshal(jsonData, &decodeUsers)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	fmt.Println("\nDecoded JSON data:", decodeUsers)
}
