/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// readCmd represents the read command
var readCmd = &cobra.Command{
	Use:   "read [filename]",
	Short: "Read and display the content of a given file",
	Args:  cobra.ExactArgs(1), // set the number of args the commands takes
	Run: func(cmd *cobra.Command, args []string) {
		// set the filename to the args index
		filename := args[0]

		// assign variable to the content we want to read from the filename
		content, err := os.ReadFile(filename)
		if err != nil {
			fmt.Printf("Error reading file: %v\n", err)
			return
		}
		fmt.Printf("Filename: %s\n%s\n", filename, string(content))
	},
}

func init() {
	// add command to the root
	rootCmd.AddCommand(readCmd)
}
