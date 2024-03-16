package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "krouly",
	Short: "CLI tool for managing data extraction workflows",
}

var createCmd = &cobra.Command{
	Use:   "create [name]",
	Short: "Create a new data extraction workflow",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		fmt.Printf("Creating data extraction workflow: %s\n", name)
		// Add logic to create a new data extraction workflow
	},
}

func main() {
	rootCmd.AddCommand(createCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
