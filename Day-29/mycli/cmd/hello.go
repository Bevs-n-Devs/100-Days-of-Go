/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// helloCmd represents the hello command
var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Prints a greeting",
	Long:  `This prints a longer message`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		if name == "" {
			name = "there"
		}
		fmt.Printf("Hello, %s!\n", name)
	},
}

func init() {
	rootCmd.AddCommand(helloCmd)

	// Here you will define your flags and configuration settings.
	helloCmd.Flags().StringP("name", "n", "", "The name to greet")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// helloCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// helloCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
