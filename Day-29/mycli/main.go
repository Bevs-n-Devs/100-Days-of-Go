/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"

	"github.com/Bevs-n-Devs/100-Days-of-Go/Day-29/mycli/cmd"
)

func main() {
	fmt.Println("Hello world, hello Yaw!")
	fmt.Println("\nCreating A Basic CLI with Golang")

	intro := `
Creating command-line applications (CLI) in Go is straightforward and powerful. 
The Cobra library is one of the most popular tools for building CLIs in Go, used by tools like Kubernetes' 'kubectl' and 'Hugo'. 
Today, we’ll dive into the essentials of creating a CLI with Cobra, covering how to set up commands, flags, and subcommands`

	objectives := `
Objectives:
1. Install and Initialize Cobra: Setting up your Go environment to work with Cobra.
2. Creating Commands and Flags: Building primary commands and adding flags.
3. Creating Subcommands: Structuring your CLI with nested commands.
4. Best Practices for CLI Development: Design considerations for user-friendly CLI tools.`

	runHello := `
Run the main file:
- go run main.go

Run the hello command:
- go run main.go hello --name=Yaw
- go run main.go hello -n=Yaw`

	runGoodbye := `
Run the goodbyge command:
- go run main.go hello goodbye --name=Yaw
- go run main.go hello goodybe -n=Yaw`

	fmt.Println(intro)
	fmt.Println(objectives)
	fmt.Println(runHello)
	fmt.Println(runGoodbye)

	cmd.Execute()
}
