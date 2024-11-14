/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// goodbyeCmd represents the goodbye command
var goodbyeCmd = &cobra.Command{
	Use:   "goodbye",
	Short: "Prints goodbye to the user",
	Long:  `A longer description can be made here`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		if name == "" {
			name = "friend"
		}
		fmt.Printf("Goodbye, %s!", name)
	},
}

func init() {
	helloCmd.AddCommand(goodbyeCmd)

	// Here you will define your flags and configuration settings.
	goodbyeCmd.Flags().StringP("name", "n", "", "The name to goodbye to")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// goodbyeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// goodbyeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
