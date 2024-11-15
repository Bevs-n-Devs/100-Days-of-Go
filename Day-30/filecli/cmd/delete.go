/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [filename]",
	Short: "A brief description of your command",
	Args:  cobra.ExactArgs(1), // set the number of args the command takes
	Run: func(cmd *cobra.Command, args []string) {
		// set filename to the args index
		filename := args[0]

		// remove the file from operating system
		err := os.Remove(filename)

		if err != nil {
			fmt.Printf("Error deleting file: %v\n", err)
			return
		}
		fmt.Printf("File %s deleted successfully.\n", filename)
	},
}

func init() {
	// add command to the root (main.go file)
	rootCmd.AddCommand(deleteCmd)
}
