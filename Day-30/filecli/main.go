/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"

	"github.com/Bevs-n-Devs/100-Days-of-Go/Day-30/filecli/cmd"
)

func main() {
	fmt.Println("Hello world, hello Yaw!")

	fmt.Println("\nBuiding Real-World CLI with Cobra - File Management System")

	intro := `
This file management CLI that allows users to create, read, update, and delete (CRUD) text files. 
This example will demonstrate Cobra’s capabilities while giving you a tool to manage files from the command line.`
	objectives := `
Objectives:
1. Set Up Cobra for a File Management CLI: Initialize the CLI with commands and subcommands.
2. Implementing CRUD Commands: Write functions for creating, reading, updating, and deleting files.
3. Best Practices: Cover error handling, argument validation, and user-friendly output.`

	operations := `
Create a new file:
- go run main.go create <filename> <some random file content goes here>

Read existing file:
- go run main.go read <filename>

Update existing file:
- go run main.go update <filename> <add to content to existing file>

Delete existing file:
- go run main.go delete <filename>
`

	fmt.Println(intro)
	fmt.Println(objectives)
	fmt.Println(operations)
	cmd.Execute()
}
