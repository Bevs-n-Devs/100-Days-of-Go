/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update [filename] [new content]",
	Short: "Append content to an existing file",
	Args:  cobra.MinimumNArgs(2), // set the minimum amount of arguments needed for the command
	Run: func(cmd *cobra.Command, args []string) {
		// set the filename you wish to update and the content you wish to add
		filename := args[0]
		newContent := args[1]

		// specify the mode whcih the file should be opened
		// O_APPEND: This tells the file to open in append mode, meaning new data will be added to the end of the file.
		// os.O_WRONLY: This specifies write-only mode, meaning you can only write to the file, not read from it.
		// 0644: This sets the file’s permissions:
		file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)

		if err != nil {
			fmt.Printf("Error openign file: %v\n", err)
			return
		}
		defer file.Close()

		// update existing file
		file.WriteString("\n" + newContent)
		fmt.Printf("File %s updated successfully.\n", filename)
	},
}

func init() {
	// add command to the root (main.go file)
	rootCmd.AddCommand(updateCmd)
}
