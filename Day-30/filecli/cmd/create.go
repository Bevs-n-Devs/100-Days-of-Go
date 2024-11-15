/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create [filename] [content]",
	Short: "Create a new file with the specified content",
	Args:  cobra.MinimumNArgs(2), // set the amount of args the comamnds needs
	Run: func(cmd *cobra.Command, args []string) {
		// create arguments for create command
		filename := args[0]
		content := args[1]

		file, err := os.Create(filename) // create the filename
		if err != nil {
			fmt.Printf("Error creating file: %v\n", err)
			return
		}
		defer file.Close()

		// create the content for the specified file
		file.WriteString(content)
		fmt.Printf("File %s created successfully.\n", filename)
	},
}

func init() {
	// add command to the root (main.go file)
	rootCmd.AddCommand(createCmd)
}
